package repositories_test

import (
	"context"
	"patient-appointment-demo-go/internal/database"
	"patient-appointment-demo-go/internal/repositories"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAppointmentQueries struct {
	mock.Mock
}

func (m *MockAppointmentQueries) GetAllAppointments(ctx context.Context) ([]database.Appointment, error) {
	args := m.Called(ctx)
	return args.Get(0).([]database.Appointment), args.Error(1)
}

func (m *MockAppointmentQueries) GetAppointmentsByDate(ctx context.Context, date pgtype.Date) ([]database.Appointment, error) {
	args := m.Called(ctx, date)
	return args.Get(0).([]database.Appointment), args.Error(1)
}

func (m *MockAppointmentQueries) GetAppointmentsByPatient(ctx context.Context, patientId int32) ([]database.Appointment, error) {
	args := m.Called(ctx, patientId)
	return args.Get(0).([]database.Appointment), args.Error(1)
}

func (m *MockAppointmentQueries) GetAppointmentByID(ctx context.Context, id int32) (database.Appointment, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(database.Appointment), args.Error(1)
}

func (m *MockAppointmentQueries) CreateAppointment(ctx context.Context, params database.CreateAppointmentParams) (database.Appointment, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(database.Appointment), args.Error(1)
}

func (m *MockAppointmentQueries) UpdateAppointment(ctx context.Context, params database.UpdateAppointmentParams) (database.Appointment, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(database.Appointment), args.Error(1)
}

func (m *MockAppointmentQueries) DeleteAppointment(ctx context.Context, id int32) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestAppointmentRepository_GetAll(t *testing.T) {
	mockQueries := new(MockAppointmentQueries)
	repo := repositories.NewAppointmentRepository(mockQueries)
	ctx := context.Background()
	appointments := []database.Appointment{{ID: 1}}

	mockQueries.On("GetAllAppointments", ctx).Return(appointments, nil)

	result, err := repo.GetAll(ctx)

	assert.NoError(t, err)
	assert.Equal(t, appointments, result)
	mockQueries.AssertExpectations(t)
}

func TestAppointmentRepository_GetByDate(t *testing.T) {
	mockQueries := new(MockAppointmentQueries)
	repo := repositories.NewAppointmentRepository(mockQueries)
	ctx := context.Background()
	date := time.Now()
	appointments := []database.Appointment{{ID: 1}}

	mockQueries.On("GetAppointmentsByDate", ctx, mock.Anything).Return(appointments, nil)

	result, err := repo.GetByDate(ctx, date)

	assert.NoError(t, err)
	assert.Equal(t, appointments, result)
	mockQueries.AssertExpectations(t)
}

func TestAppointmentRepository_Get(t *testing.T) {
	mockQueries := new(MockAppointmentQueries)
	repo := repositories.NewAppointmentRepository(mockQueries)
	ctx := context.Background()
	appointment := database.Appointment{ID: 1}

	mockQueries.On("GetAppointmentByID", ctx, int32(1)).Return(appointment, nil)

	result, err := repo.Get(ctx, 1)

	assert.NoError(t, err)
	assert.Equal(t, appointment, result)
	mockQueries.AssertExpectations(t)
}

func TestAppointmentRepository_Create(t *testing.T) {
	mockQueries := new(MockAppointmentQueries)
	repo := repositories.NewAppointmentRepository(mockQueries)
	ctx := context.Background()
	appointment := database.Appointment{ID: 1}
	params := repositories.CreateAppointmentParams{
		VisitTimestamp: time.Now(),
		PatientNotes:   nil,
	}

	mockQueries.On("CreateAppointment", ctx, mock.Anything).Return(appointment, nil)

	result, err := repo.Create(ctx, 1, 2, params)

	assert.NoError(t, err)
	assert.Equal(t, appointment, result)
	mockQueries.AssertExpectations(t)
}

func TestAppointmentRepository_Update(t *testing.T) {
	mockQueries := new(MockAppointmentQueries)
	repo := repositories.NewAppointmentRepository(mockQueries)
	ctx := context.Background()
	appointment := database.Appointment{ID: 1}
	params := repositories.UpdateAppointmentParams{
		PatientNotes: nil,
		DoctorNotes:  nil,
	}

	mockQueries.On("UpdateAppointment", ctx, mock.Anything).Return(appointment, nil)

	result, err := repo.Update(ctx, 1, params)

	assert.NoError(t, err)
	assert.Equal(t, appointment, result)
	mockQueries.AssertExpectations(t)
}

func TestAppointmentRepository_Delete(t *testing.T) {
	mockQueries := new(MockAppointmentQueries)
	repo := repositories.NewAppointmentRepository(mockQueries)
	ctx := context.Background()

	mockQueries.On("DeleteAppointment", ctx, int32(1)).Return(nil)

	err := repo.Delete(ctx, 1)

	assert.NoError(t, err)
	mockQueries.AssertExpectations(t)
}

