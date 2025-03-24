package routes

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

type AppointmentRouter struct {
	mux      *http.ServeMux
	userRepo repositories.UserRepositoryInterface
	repo     repositories.AppointmentRepositoryInterface
}

func NewAppointmentRouter(mux *http.ServeMux, appointmentRepo repositories.AppointmentRepositoryInterface, userRepo repositories.UserRepositoryInterface) *AppointmentRouter {
    return &AppointmentRouter{
        mux: mux,
        repo: appointmentRepo,
        userRepo: userRepo,
    }
}

func (r *AppointmentRouter) Register() *AppointmentRouter {
    authMiddleware := NewAuthMiddleware(r.userRepo)

	NewRoute("GET", "/api/appointments").
        SetHandler(r.GetAll).
        AddMiddlewares(authMiddleware.ValidateLogin).
        Register(r.mux)

	NewRoute("GET", "/api/appointments/date/{date}").
        SetHandler(r.GetByDate).
        AddMiddlewares(authMiddleware.ValidateLogin).
        Register(r.mux)

	NewRoute("GET", "/api/patients/{patientId}/appointments").
        SetHandler(r.GetByPatient).
        AddMiddlewares(authMiddleware.ValidateLogin).
        Register(r.mux)

	NewRoute("GET", "/api/appointments/{id}").
        SetHandler(r.Get).
        AddMiddlewares(authMiddleware.ValidateLogin).
        Register(r.mux)

	NewRoute("POST", "/api/patients/{patientId}/appointments").
        SetHandler(r.Create).
        AddMiddlewares(authMiddleware.ValidateLogin).
        Register(r.mux)

	NewRoute("PUT", "/api/appointments/{id}").
        SetHandler(r.Update).
        AddMiddlewares(
            authMiddleware.ValidateLogin,
            authMiddleware.ValidateRole("doctor"),
        ).
        Register(r.mux)

	NewRoute("DELETE", "/api/appointments/{id}").
        SetHandler(r.Delete).
        AddMiddlewares(authMiddleware.ValidateLogin).
        Register(r.mux)

	return r
}

func (ac *AppointmentRouter) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	appointments, err := ac.repo.GetAll(ctx)

	if err != nil {
		http.Error(w, "Failed to fetch appointments", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AppointmentDbArrayToResponse(appointments))
}

func (ac *AppointmentRouter) GetByDate(w http.ResponseWriter, r *http.Request) {
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
	json.NewEncoder(w).Encode(AppointmentDbArrayToResponse(appointments))
}

func (ac *AppointmentRouter) GetByPatient(w http.ResponseWriter, r *http.Request) {
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
	json.NewEncoder(w).Encode(AppointmentDbArrayToResponse(appointments))
}

func (ac *AppointmentRouter) Get(w http.ResponseWriter, r *http.Request) {
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
	json.NewEncoder(w).Encode(AppointmentDbToResponse(appointment))

}

func (ac *AppointmentRouter) Create(w http.ResponseWriter, r *http.Request) {

	patientIdStr := r.PathValue("patientId")
	patientId, err := strconv.ParseInt(patientIdStr, 10, 64)

	if err != nil {
        fmt.Println(err)
		http.Error(w, "Invalid patient id", http.StatusBadRequest)
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
	json.NewEncoder(w).Encode(AppointmentDbToResponse(appointment))
}

func (ac *AppointmentRouter) Update(w http.ResponseWriter, r *http.Request) {

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
	json.NewEncoder(w).Encode(AppointmentDbToResponse(appointment))

}

func (ac *AppointmentRouter) Delete(w http.ResponseWriter, r *http.Request) {

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
