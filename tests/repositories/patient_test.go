package repositories_test

import (
	"context"
	"patient-appointment-demo-go/internal/database"
	"patient-appointment-demo-go/internal/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockQueries struct {
	mock.Mock
}

func (m *MockQueries) GetAllPatients(ctx context.Context, params database.GetAllPatientsParams) ([]database.Patient, error) {
	args := m.Called(ctx, params)
	return args.Get(0).([]database.Patient), args.Error(1)
}

func (m *MockQueries) GetPatientByID(ctx context.Context, id int32) (database.Patient, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(database.Patient), args.Error(1)
}

func (m *MockQueries) CreatePatient(ctx context.Context, params database.CreatePatientParams) (database.Patient, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(database.Patient), args.Error(1)
}

func (m *MockQueries) UpdatePatient(ctx context.Context, params database.UpdatePatientParams) (database.Patient, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(database.Patient), args.Error(1)
}

func (m *MockQueries) DeletePatient(ctx context.Context, id int32) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestPatientRepository_GetAll(t *testing.T) {
	mockQueries := new(MockQueries)
	repo := repositories.NewPatientRepository(mockQueries)
	ctx := context.Background()
	patients := []database.Patient{{ID: 1, Name: "John Doe"}}
	params := database.GetAllPatientsParams{Name: "John"}

	mockQueries.On("GetAllPatients", ctx, params).Return(patients, nil)

	result, err := repo.GetAll(ctx, repositories.GetPatientsOption{Name: "John"})

	assert.NoError(t, err)
	assert.Equal(t, patients, result)
	mockQueries.AssertExpectations(t)
}

func TestPatientRepository_Get(t *testing.T) {
	mockQueries := new(MockQueries)
	repo := repositories.NewPatientRepository(mockQueries)
	ctx := context.Background()
	patient := database.Patient{ID: 1, Name: "John Doe"}

	mockQueries.On("GetPatientByID", ctx, int32(1)).Return(patient, nil)

	result, err := repo.Get(ctx, 1)

	assert.NoError(t, err)
	assert.Equal(t, patient, result)
	mockQueries.AssertExpectations(t)
}

func TestPatientRepository_Create(t *testing.T) {
	mockQueries := new(MockQueries)
	repo := repositories.NewPatientRepository(mockQueries)
	ctx := context.Background()
	patient := database.Patient{ID: 1, Name: "John Doe"}
	params := repositories.CreatePatientParams{
		Name: "John Doe", Phone: "1234567890", Email: "john@example.com", Age: 30,
		Weight: 70.5, Height: 175.2, Gender: "Male", Address: "123 Street"}

	mockQueries.On("CreatePatient", ctx, mock.Anything).Return(patient, nil)

	result, err := repo.Create(ctx, params)

	assert.NoError(t, err)
	assert.Equal(t, patient, result)
	mockQueries.AssertExpectations(t)
}

func TestPatientRepository_Update(t *testing.T) {
	mockQueries := new(MockQueries)
	repo := repositories.NewPatientRepository(mockQueries)
	ctx := context.Background()
	patient := database.Patient{ID: 1, Name: "John Updated"}
	params := repositories.UpdatePatientParams{
		Name: "John Updated", Phone: "1234567890", Email: "john@example.com", Age: 31,
		Weight: 71.5, Height: 176.2, Gender: "Male", Address: "123 Updated Street"}

	mockQueries.On("UpdatePatient", ctx, mock.Anything).Return(patient, nil)

	result, err := repo.Update(ctx, 1, params)

	assert.NoError(t, err)
	assert.Equal(t, patient, result)
	mockQueries.AssertExpectations(t)
}

func TestPatientRepository_Delete(t *testing.T) {
	mockQueries := new(MockQueries)
	repo := repositories.NewPatientRepository(mockQueries)
	ctx := context.Background()

	mockQueries.On("DeletePatient", ctx, int32(1)).Return(nil)

	err := repo.Delete(ctx, 1)

	assert.NoError(t, err)
	mockQueries.AssertExpectations(t)
}

