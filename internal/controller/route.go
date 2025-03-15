package controller

import (
	"patient-appointment-demo-go/internal/app"
	"patient-appointment-demo-go/internal/database"
	"patient-appointment-demo-go/internal/repositories"
)

type Router struct {
	app             *app.App
	patientRepo     repositories.PatientRepositoryInterface
	appointmentRepo repositories.AppointmentRepositoryInterface
}

func NewRouter(app *app.App) {
	r := Router{
		app:             app,
		patientRepo:     repositories.NewPatientRepository(*database.New(app.DbConn)),
		appointmentRepo: repositories.NewAppointmentRepository(*database.New(app.DbConn)),
	}

	r.registerAuthRoute()
	r.registerPatientRoute()
	r.registerAppointmentRoute()
}

func (r *Router) registerAuthRoute() {
	c := NewAuthController(r.app.DbConn)
	r.app.Mux.HandleFunc("POST /api/auth/login", c.Login)
	r.app.Mux.HandleFunc("POST /api/auth/logout", c.Logout)
}

func (r *Router) registerPatientRoute() {
	authMiddleware := NewAuthMiddleware(*database.New(r.app.DbConn))
	patientController := NewPatientController(r.patientRepo)

	r.app.Mux.Handle("GET /api/patients", authMiddleware.ValidateLogin(patientController.GetAll))
	r.app.Mux.Handle("GET /api/patients/{id}", authMiddleware.ValidateLogin(patientController.Get))
	r.app.Mux.Handle("POST /api/patients", authMiddleware.ValidateLogin(patientController.Create))
	r.app.Mux.Handle("PUT /api/patients/{id}", authMiddleware.ValidateLogin(patientController.Update))
	r.app.Mux.Handle("DELETE /api/patients/{id}", authMiddleware.ValidateLogin(patientController.Delete))

}

func (r *Router) registerAppointmentRoute() {
	authMiddleware := NewAuthMiddleware(*database.New(r.app.DbConn))
	appointmentController := NewAppointmentController(r.appointmentRepo)

	r.app.Mux.Handle(
		"GET /api/appointments",
		authMiddleware.ValidateLogin(appointmentController.GetAll),
	)

	r.app.Mux.Handle(
		"GET /api/appointments/date/{date}",
		authMiddleware.ValidateLogin(appointmentController.GetByDate),
	)

	r.app.Mux.Handle(
		"GET /api/appointments/{id}",
		authMiddleware.ValidateLogin(appointmentController.Get),
	)

	r.app.Mux.Handle(
		"GET /api/patients/{patientId}/appointments",
		authMiddleware.ValidateLogin(appointmentController.GetByPatient),
	)

	r.app.Mux.Handle(
		"POST /api/patients/{patientId}/appointments",
		authMiddleware.ValidateLogin(appointmentController.Create),
	)

	r.app.Mux.Handle(
		"PUT /api/appointments/{id}",
		authMiddleware.ValidateLogin(
			authMiddleware.ValidateRole(appointmentController.Update, "doctor"),
		),
	)

	r.app.Mux.Handle(
        "DELETE /api/appointments/{id}",
        authMiddleware.ValidateLogin(appointmentController.Delete),
    )

}
