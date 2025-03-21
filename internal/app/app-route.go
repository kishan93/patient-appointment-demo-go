package app

import "patient-appointment-demo-go/internal/routes"

func (a *App) initRouter() {
    routes.NewAuthRouter(a.Mux, a.UserRepo()).Register()
    routes.NewPatientRouter(a.Mux, a.PatientRepo(), a.UserRepo()).Register()
    routes.NewAppointmentRouter(a.Mux, a.AppointmentRepo(), a.UserRepo()).Register()
}
