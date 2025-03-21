package app

import (
	"patient-appointment-demo-go/internal/database"
	"patient-appointment-demo-go/internal/repositories"
)


func (a *App) UserRepo() repositories.UserRepositoryInterface {
    return repositories.NewUserRepository(*database.New(a.DbConn))
}

func (a *App) PatientRepo() repositories.PatientRepositoryInterface {
    return repositories.NewPatientRepository(*database.New(a.DbConn))
}

func (a *App) AppointmentRepo() repositories.AppointmentRepositoryInterface {
    return repositories.NewAppointmentRepository(*database.New(a.DbConn))
}

