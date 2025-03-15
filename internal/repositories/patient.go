package repositories

import (
	"context"
	"math/big"
	"patient-appointment-demo-go/internal/database"

	"github.com/jackc/pgx/v5/pgtype"
)

type PatientRepository struct {
	queries database.Queries
}

type CreatePatientParams struct {
	Name    string
	Phone   string
	Email   string
	Age     int32
	Weight  float32
	Height  float32
	Gender  string
	Address string
}

type UpdatePatientParams struct {
	Name    string
	Phone   string
	Email   string
	Age     int32
	Weight  float32
	Height  float32
	Gender  string
	Address string
}

type GetPatientsOption struct {
	Name          string
	SortBy        string
	SortDirection string
}

func NewPatientRepository(queries database.Queries) PatientRepositoryInterface {
	return &PatientRepository{
        queries: queries,
    }
}

func (p *PatientRepository) GetAll(ctx context.Context, option GetPatientsOption) ([]database.Patient, error) {

	patients, err := p.queries.GetAllPatients(ctx, database.GetAllPatientsParams{
		Name:          option.Name,
		SortBy:        option.SortBy,
		SortDirection: option.SortDirection,
	})

	return patients, err
}

func (p *PatientRepository) Get(ctx context.Context, id int32) (database.Patient, error) {

	patient, err := p.queries.GetPatientByID(ctx, id)

	return patient, err
}

func (pc *PatientRepository) Create(ctx context.Context, data CreatePatientParams) (database.Patient, error) {

    //TODO: validate email or phone not already taken

    weightNumeric := pgtype.Numeric{
		Int:   big.NewInt(int64(data.Weight * 100)),
		Exp:   -2,
		Valid: data.Weight > 0,
	}
	heightNumeric := pgtype.Numeric{
		Int:   big.NewInt(int64(data.Height * 100)),
		Exp:   -2,
		Valid: data.Height > 0,
	}
	patient, err := pc.queries.CreatePatient(ctx, database.CreatePatientParams{
		Name:    data.Name,
		Phone:   pgtype.Text{String: data.Phone, Valid: data.Phone != ""},
		Email:   data.Email,
		Age:     pgtype.Int2{Int16: int16(data.Age), Valid: data.Age > 0},
		Weight:  weightNumeric,
		Height:  heightNumeric,
		Gender:  pgtype.Text{String: data.Gender, Valid: data.Gender != ""},
		Address: pgtype.Text{String: data.Address, Valid: data.Address != ""},
	})

	return patient, err
}

func (pc *PatientRepository) Update(ctx context.Context, id int32, data UpdatePatientParams) (database.Patient, error) {

    //TODO: validate email or phone not already taken

    weightNumeric := pgtype.Numeric{
		Int:   big.NewInt(int64(data.Weight * 100)),
		Exp:   -2,
		Valid: data.Weight > 0,
	}
	heightNumeric := pgtype.Numeric{
		Int:   big.NewInt(int64(data.Height * 100)),
		Exp:   -2,
		Valid: data.Height > 0,
	}
	updatedPatient, err := pc.queries.UpdatePatient(ctx, database.UpdatePatientParams{
        ID: id,
		Name:    data.Name,
		Phone:   pgtype.Text{String: data.Phone, Valid: data.Phone != ""},
		Email:   data.Email,
		Age:     pgtype.Int2{Int16: int16(data.Age), Valid: data.Age > 0},
		Weight:  weightNumeric,
		Height:  heightNumeric,
		Gender:  pgtype.Text{String: data.Gender, Valid: data.Gender != ""},
		Address: pgtype.Text{String: data.Address, Valid: data.Address != ""},
	})

	return updatedPatient, err
}

func (pc *PatientRepository) Delete(ctx context.Context, id int32) error {

	err := pc.queries.DeletePatient(ctx, id)

	return err
}
