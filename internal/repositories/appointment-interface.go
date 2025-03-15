package repositories

import (
	"context"
	"patient-appointment-demo-go/internal/database"
	"time"
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
