package routes

import (
	"encoding/json"
	"net/http"
	"patient-appointment-demo-go/internal/repositories"
	"patient-appointment-demo-go/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthRouter struct {
    mux *http.ServeMux
	repo repositories.UserRepositoryInterface
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string `json:"token"`
}

func NewAuthRouter(mux *http.ServeMux, repo repositories.UserRepositoryInterface) *AuthRouter {
	return &AuthRouter{
		repo: repo,
        mux: mux,
	}
}

func (r *AuthRouter) Register() *AuthRouter {
    authMiddleware := NewAuthMiddleware(r.repo)
	NewRoute("POST", "/api/auth/login").
        SetHandler(r.Login).
        Register(r.mux)

	NewRoute("POST", "/api/auth/logout").
        SetHandler(r.Logout).
        AddMiddlewares(authMiddleware.ValidateLogin).
        Register(r.mux)

	return r
}

func (a *AuthRouter) Login(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Fetch user from DB
	user, err := a.repo.GetByEmail(r.Context(), req.Email) // Ensure GetUserByEmail exists in database.Queries

	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Validate password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	tokenString, err := utils.GenerateJWT(user.ID)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(loginResponse{Token: tokenString})
}

func (a *AuthRouter) Logout(w http.ResponseWriter, r *http.Request) {
	// TODO: implement token blacklisting

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logout successful"))
}
