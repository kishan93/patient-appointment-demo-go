package app

import (
	"fmt"
	"net/http"
	"patient-appointment-demo-go/internal/routes"
)

func (a *App) initRouter() http.Handler{
	routes.NewRoute("OPTION", "/api/").
		SetHandler(func(w http.ResponseWriter, r *http.Request) {

			fmt.Printf("OPTIONS req: %s\n", r.URL)
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

            w.WriteHeader(http.StatusNoContent)

		}).
		Register(a.Mux)

	routes.NewAuthRouter(a.Mux, a.UserRepo()).Register()
	routes.NewPatientRouter(a.Mux, a.PatientRepo(), a.UserRepo()).Register()
	routes.NewAppointmentRouter(a.Mux, a.AppointmentRepo(), a.UserRepo()).Register()

    return routes.CorsMiddleware(a.Mux)
}
