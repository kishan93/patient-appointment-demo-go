package repositories

import (
	"context"
	"patient-appointment-demo-go/internal/database"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type AppointmentRepository struct {
	queries database.Queries
}

type CreateAppointmentParams struct {
	VisitTimestamp time.Time
	PatientNotes   *string
}

type UpdateAppointmentParams struct {
	PatientNotes *string
	DoctorNotes  *string
}

func NewAppointmentRepository(queries database.Queries) AppointmentRepositoryInterface {
	return &AppointmentRepository{
		queries: queries,
	}
}

func (a *AppointmentRepository) GetAll(ctx context.Context) ([]database.Appointment, error) {
	res, err := a.queries.GetAllAppointments(ctx)
	return res, err
}

func (a *AppointmentRepository) GetByDate(ctx context.Context, date time.Time) ([]database.Appointment, error) {
	pgDate := pgtype.Date{
		Time:  date,
		Valid: true,
	}

	res, err := a.queries.GetAppointmentsByDate(ctx, pgDate)
	return res, err
}

func (a *AppointmentRepository) GetByPatient(ctx context.Context, patientId int32) ([]database.Appointment, error) {
	res, err := a.queries.GetAppointmentsByPatient(ctx, patientId)
	return res, err
}

func (a *AppointmentRepository) Get(ctx context.Context, id int32) (database.Appointment, error) {
	res, err := a.queries.GetAppointmentByID(ctx, id)

	return res, err
}

func (a *AppointmentRepository) Create(ctx context.Context, userId int32, patientId int32, data CreateAppointmentParams) (database.Appointment, error) {

	var pgTimestamp pgtype.Timestamptz
	pgTimestamp.Time = data.VisitTimestamp
	pgTimestamp.Valid = true

	var pgDate pgtype.Date

	visitTime := data.VisitTimestamp
	startOfDay := time.Date(
		visitTime.Year(),
		visitTime.Month(),
		visitTime.Day(),
		0,
		0,
		0,
		0,
		visitTime.Location(),
	)
	pgDate.Time = startOfDay
	pgDate.Valid = true

    var patientNotes string
    if data.PatientNotes !=nil {
        patientNotes = *data.PatientNotes
    }

	res, err := a.queries.CreateAppointment(ctx, database.CreateAppointmentParams{
		UserID:         pgtype.Int4{Int32: userId},
		PatientID:      patientId,
		VisitDate:      pgDate,
		VisitTimestamp: pgTimestamp,
		PatientNotes:   pgtype.Text{String: patientNotes, Valid: data.PatientNotes != nil},
	})

	return res, err
}

func (a *AppointmentRepository) Update(ctx context.Context, appointmentId int32, data UpdateAppointmentParams) (database.Appointment, error) {

    var patientNotes string
    if data.PatientNotes !=nil {
        patientNotes = *data.PatientNotes
    }

    var doctorNotes string
    if data.DoctorNotes !=nil {
        doctorNotes = *data.DoctorNotes
    }

	updatedAppointment, err := a.queries.UpdateAppointment(ctx, database.UpdateAppointmentParams{
		ID:           appointmentId,
		PatientNotes: pgtype.Text{String: patientNotes, Valid: data.PatientNotes != nil},
		DoctorNotes:  pgtype.Text{String: doctorNotes, Valid: data.DoctorNotes != nil},
	})

	return updatedAppointment, err
}

func (a *AppointmentRepository) Delete(ctx context.Context, id int32) error {

	err := a.queries.DeleteAppointment(ctx, id)

	return err
}
