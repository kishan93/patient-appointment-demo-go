package routes

import (
	"net/http"
)

func ApiMiddleware(next http.Handler) http.HandlerFunc{
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.Header().Set("Access-Control-Allow-Origin", "*")

		next.ServeHTTP(w, r)
    })
}

