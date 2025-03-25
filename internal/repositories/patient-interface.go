package repositories

import (
	"context"
	"patient-appointment-demo-go/internal/database"
)

type PatientRepositoryInterface interface {
	GetAll(ctx context.Context, option GetPatientsOption) ([]database.Patient, error)
	Get(ctx context.Context, id int32) (database.Patient, error)
	Create(ctx context.Context, data CreatePatientParams) (database.Patient, error)
	Update(ctx context.Context, id int32, data UpdatePatientParams) (database.Patient, error)
	Delete(ctx context.Context, id int32) error
}

type PatientQueriesContract interface {
    GetAllPatients(context.Context, database.GetAllPatientsParams) ([]database.Patient, error)
    GetPatientByID(context.Context, int32) (database.Patient, error)
    CreatePatient(context.Context, database.CreatePatientParams) (database.Patient, error)
    UpdatePatient(context.Context, database.UpdatePatientParams) (database.Patient, error)
    DeletePatient(context.Context, int32) error
}
