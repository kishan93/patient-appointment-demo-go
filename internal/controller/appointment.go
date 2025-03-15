package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"patient-appointment-demo-go/internal/database"
	"patient-appointment-demo-go/internal/repositories"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
)

type AppointmentController struct {
	repo repositories.AppointmentRepositoryInterface
}

func NewAppointmentController(repo repositories.AppointmentRepositoryInterface) AppointmentController {
	return AppointmentController{
		repo: repo,
	}
}

func (ac *AppointmentController) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	appointments, err := ac.repo.GetAll(ctx)

	if err != nil {
		http.Error(w, "Failed to fetch appointments", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(appointments)
}

func (ac *AppointmentController) GetByDate(w http.ResponseWriter, r *http.Request) {
	date := r.PathValue("date")

	parsedDate, err := time.Parse("2006-01-02", date)

	if err != nil {
        fmt.Println(err)
		http.Error(w, "Invalid date", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	appointments, err := ac.repo.GetByDate(ctx, parsedDate)

	if err != nil {
		http.Error(w, "Failed to fetch appointments", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(appointments)
}

func (ac *AppointmentController) GetByPatient(w http.ResponseWriter, r *http.Request) {
	patientIdStr := r.PathValue("patientId")
	patientId, err := strconv.ParseInt(patientIdStr, 10, 64)

	if err != nil {
		http.Error(w, "invalid patient id", http.StatusBadRequest)
        return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	appointments, err := ac.repo.GetByPatient(ctx, int32(patientId))

	if err != nil {
		http.Error(w, "Failed to fetch appointments for the patient", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(appointments)
}

func (ac *AppointmentController) Get(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		http.Error(w, "invalid appointment id", http.StatusBadRequest)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	appointment, err := ac.repo.Get(ctx, int32(id))

	if err != nil {
		http.Error(w, "Failed to fetch appointments for the patient", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(appointment)

}

func (ac *AppointmentController) Create(w http.ResponseWriter, r *http.Request) {

	patientIdStr := r.PathValue("patientId")
	patientId, err := strconv.ParseInt(patientIdStr, 10, 64)

	if err != nil {
        fmt.Println(err)
		http.Error(w, "invalid patient id", http.StatusBadRequest)
        return
	}

	var req AppointmentCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        fmt.Println(err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
	}

    validate := validator.New()

    if err := validate.Struct(req); err != nil {
		ve := err.(validator.ValidationErrors)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]map[string]string{
			"errors": getValidationErrors(ve),
		})
		return
    }

	user, ok := r.Context().Value("user").(database.User)
	if !ok {
		http.Error(w, "Failed to get user data", http.StatusInternalServerError)
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	appointment, err := ac.repo.Create(ctx, user.ID, int32(patientId), repositories.CreateAppointmentParams{
        VisitTimestamp: req.VisitTime,
        PatientNotes: req.PatientNotes,
    })

	if err != nil {
		http.Error(w, "Failed to create appointment", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(appointment)
}

func (ac *AppointmentController) Update(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		http.Error(w, "Invalid appointment id", http.StatusBadRequest)
        return
	}

	var req AppointmentUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	appointment, err := ac.repo.Update(ctx, int32(id), repositories.UpdateAppointmentParams{
        PatientNotes: req.PatientNotes,
        DoctorNotes: req.DoctorNotes,
    })

	if err != nil {
        fmt.Println(err)
		http.Error(w, "Failed to update appointment", http.StatusInternalServerError)
        return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(appointment)

}

func (ac *AppointmentController) Delete(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		http.Error(w, "Invalid appointment id", http.StatusBadRequest)
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	err = ac.repo.Delete(ctx, int32(id))

	w.Write([]byte(""))

}
