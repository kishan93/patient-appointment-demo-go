package controller

import (
	"encoding/json"
	"net/http"

	"patient-appointment-demo-go/internal/database"
	"patient-appointment-demo-go/internal/utils"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	queries database.Queries
}

func NewAuthController(dbConn *pgx.Conn) *AuthController {
	return &AuthController{
		queries: *database.New(dbConn),
	}
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string `json:"token"`
}

func (u *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Fetch user from DB
	user, err := u.queries.GetUserByEmail(r.Context(), req.Email) // Ensure GetUserByEmail exists in database.Queries
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(loginResponse{Token: tokenString})
}

func (u *AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	// TODO: implement token blacklisting

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logout successful"))
}
