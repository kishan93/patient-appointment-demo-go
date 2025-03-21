package routes

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"patient-appointment-demo-go/internal/database"
	"patient-appointment-demo-go/internal/repositories"
	"patient-appointment-demo-go/internal/utils"
)

type AuthMiddleware struct {
	userRepo repositories.UserRepositoryInterface
}

func NewAuthMiddleware(userRepo repositories.UserRepositoryInterface) AuthMiddleware {
	return AuthMiddleware{
		userRepo: userRepo,
	}
}

func (m AuthMiddleware) ValidateLogin(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := extractAuthToken(r)
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		userID, err := utils.ParseJWT(token)
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		user, err := m.userRepo.Get(r.Context(), userID)
		if err != nil {
			http.Error(w, "User not found", http.StatusUnauthorized)
			return
		}

		// Add user info to request context
		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m AuthMiddleware) ValidateRole(allowedRoles ...string) func(http.HandlerFunc) http.HandlerFunc{
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, err := getUserFromContext(r)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			if !hasRole(user.Type, allowedRoles) {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func extractAuthToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if strings.HasPrefix(authHeader, "Bearer ") {
		return strings.TrimPrefix(authHeader, "Bearer ")
	}
	return ""
}

func hasRole(userRole string, allowedRoles []string) bool {
	for _, role := range allowedRoles {
		if userRole == role {
			return true
		}
	}
	return false
}

func getUserFromContext(r *http.Request) (database.User, error) {
	user, ok := r.Context().Value("user").(database.User)
	if !ok {
		return database.User{}, errors.New("no user in context")
	}
	return user, nil
}
