package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"patient-appointment-demo-go/internal/repositories"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
)

type PatientRouter struct {
	mux      *http.ServeMux
	userRepo repositories.UserRepositoryInterface
	repo     repositories.PatientRepositoryInterface
}

func NewPatientRouter(mux *http.ServeMux, patientRepo repositories.PatientRepositoryInterface, userRepo repositories.UserRepositoryInterface) *PatientRouter {
    return &PatientRouter{
        mux: mux,
        repo: patientRepo,
        userRepo: userRepo,
    }
}

func (r *PatientRouter) Register() *PatientRouter {
    authMiddleware := NewAuthMiddleware(r.userRepo)

	NewRoute("GET", "/api/patients").
        SetHandler(r.GetAll).
        AddMiddlewares(authMiddleware.ValidateLogin).
        Register(r.mux)

	NewRoute("GET", "/api/patients/{id}").
        SetHandler(r.Get).
        AddMiddlewares(authMiddleware.ValidateLogin).
        Register(r.mux)


	NewRoute("POST", "/api/patients").
        SetHandler(r.Create).
        AddMiddlewares(authMiddleware.ValidateLogin).
        Register(r.mux)

	NewRoute("PUT", "/api/patients/{id}").
        SetHandler(r.Update).
        AddMiddlewares(authMiddleware.ValidateLogin).
        Register(r.mux)

	NewRoute("DELETE", "/api/patients/{id}").
        SetHandler(r.Delete).
        AddMiddlewares(authMiddleware.ValidateLogin).
        Register(r.mux)

	return r
}


func (p *PatientRouter) GetAll(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")
	sortBy := r.URL.Query().Get("sort_by")
	sortDirection := r.URL.Query().Get("sort_direction")

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	patients, err := p.repo.GetAll(ctx, repositories.GetPatientsOption{
		Name:          name,
		SortBy:        sortBy,
		SortDirection: sortDirection,
	})

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to fetch patients", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(PatientDbArrayToResponse(patients))
}

func (p *PatientRouter) Get(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid patient ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	patient, err := p.repo.Get(ctx, int32(id))
	if err != nil {
		http.Error(w, "Patient not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(PatientDbToResponse(patient))
}

func (p *PatientRouter) Create(w http.ResponseWriter, r *http.Request) {
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

	patient, err := p.repo.Create(ctx, repositories.CreatePatientParams{
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

	json.NewEncoder(w).Encode(PatientDbToResponse(patient))
}

func (p *PatientRouter) Update(w http.ResponseWriter, r *http.Request) {
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

	updatedPatient, err := p.repo.Update(ctx, int32(id), repositories.UpdatePatientParams{
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

	json.NewEncoder(w).Encode(PatientDbToResponse(updatedPatient))
}

func (p *PatientRouter) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		http.Error(w, "Invalid patient ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	err = p.repo.Delete(ctx, int32(id))
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to delete patient", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
