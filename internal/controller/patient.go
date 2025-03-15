package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"patient-appointment-demo-go/internal/repositories"

	"github.com/go-playground/validator/v10"
)

type PatientController struct {
	repo repositories.PatientRepositoryInterface
}

func NewPatientController(repo repositories.PatientRepositoryInterface) PatientController {
	return PatientController{
		repo: repo,
	}
}

func (pc *PatientController) GetAll(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")
	sortBy := r.URL.Query().Get("sort_by")
	sortDirection := r.URL.Query().Get("sort_direction")

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	patients, err := pc.repo.GetAll(ctx, repositories.GetPatientsOption{
		Name:          name,
		SortBy:        sortBy,
		SortDirection: sortDirection,
	})

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to fetch patients", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patients)
}

func (pc *PatientController) Get(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid patient ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	patient, err := pc.repo.Get(ctx, int32(id))
	if err != nil {
		http.Error(w, "Patient not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patient)
}

func (pc *PatientController) Create(w http.ResponseWriter, r *http.Request) {
	var req PatientCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
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

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	patient, err := pc.repo.Create(ctx, repositories.CreatePatientParams{
		Name:    req.Name,
		Phone:   req.Phone,
		Email:   req.Email,
		Age:     int32(req.Age),
		Weight:  float32(req.Weight),
		Height:  float32(req.Height),
		Gender:  req.Gender,
		Address: req.Address,
	})

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to create patient", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patient)
}

func (pc *PatientController) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		http.Error(w, "Invalid patient ID", http.StatusBadRequest)
		return
	}

	var req PatientUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		validationErr := err.(validator.ValidationErrors)
		json.NewEncoder(w).Encode(validationErr)
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

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	updatedPatient, err := pc.repo.Update(ctx, int32(id), repositories.UpdatePatientParams{
		Name:    req.Name,
		Phone:   req.Phone,
		Email:   req.Email,
		Age:     int32(req.Age),
		Weight:  float32(req.Weight),
		Height:  float32(req.Height),
		Gender:  req.Gender,
		Address: req.Address,
	})
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to update patient", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedPatient)
}

func (pc *PatientController) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		http.Error(w, "Invalid patient ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	err = pc.repo.Delete(ctx, int32(id))
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to delete patient", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
