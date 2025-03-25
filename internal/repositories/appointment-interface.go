package repositories

import (
	"context"
	"patient-appointment-demo-go/internal/database"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type AppointmentRepositoryInterface interface {
	GetAll(ctx context.Context) ([]database.Appointment, error)
	GetByDate(ctx context.Context, date time.Time) ([]database.Appointment, error)
	GetByPatient(ctx context.Context, patientId int32) ([]database.Appointment, error)
	Get(ctx context.Context, id int32) (database.Appointment, error)
	Create(ctx context.Context, userId int32, patientId int32, data CreateAppointmentParams) (database.Appointment, error)
	Update(ctx context.Context, appointmentid int32, data UpdateAppointmentParams) (database.Appointment, error)
	Delete(ctx context.Context, id int32) error
}

type AppointmentQueriesContract interface {
    GetAllAppointments(context.Context) ([]database.Appointment, error)
    GetAppointmentsByDate(context.Context, pgtype.Date) ([]database.Appointment, error)
    GetAppointmentsByPatient(context.Context, int32) ([]database.Appointment, error)
    GetAppointmentByID(context.Context, int32) (database.Appointment, error)
    CreateAppointment(context.Context, database.CreateAppointmentParams) (database.Appointment, error)
    UpdateAppointment(context.Context, database.UpdateAppointmentParams) (database.Appointment, error)
    DeleteAppointment(context.Context, int32) error
}
