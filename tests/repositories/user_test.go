package repositories_test

import (
	"context"
	"patient-appointment-demo-go/internal/database"
	"patient-appointment-demo-go/internal/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserQueries struct {
	mock.Mock
}

func (m *MockUserQueries) GetAllUsers(ctx context.Context) ([]database.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]database.User), args.Error(1)
}

func (m *MockUserQueries) GetUserByEmail(ctx context.Context, email string) (database.User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(database.User), args.Error(1)
}

func (m *MockUserQueries) GetUser(ctx context.Context, id int32) (database.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(database.User), args.Error(1)
}

func (m *MockUserQueries) CreateUser(ctx context.Context, params database.CreateUserParams) (database.User, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(database.User), args.Error(1)
}

func (m *MockUserQueries) UpdateUser(ctx context.Context, params database.UpdateUserParams) (database.User, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(database.User), args.Error(1)
}

func (m *MockUserQueries) DeleteUser(ctx context.Context, id int32) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestUserRepository_GetAll(t *testing.T) {
	mockQueries := new(MockUserQueries)
	repo := repositories.NewUserRepository(mockQueries)
	ctx := context.Background()
	users := []database.User{{ID: 1, Email: "test@example.com"}}

	mockQueries.On("GetAllUsers", ctx).Return(users, nil)

	result, err := repo.GetAll(ctx)

	assert.NoError(t, err)
	assert.Equal(t, users, result)
	mockQueries.AssertExpectations(t)
}

func TestUserRepository_GetByEmail(t *testing.T) {
	mockQueries := new(MockUserQueries)
	repo := repositories.NewUserRepository(mockQueries)
	ctx := context.Background()
	user := database.User{ID: 1, Email: "test@example.com"}

	mockQueries.On("GetUserByEmail", ctx, "test@example.com").Return(user, nil)

	result, err := repo.GetByEmail(ctx, "test@example.com")

	assert.NoError(t, err)
	assert.Equal(t, user, result)
	mockQueries.AssertExpectations(t)
}

func TestUserRepository_Get(t *testing.T) {
	mockQueries := new(MockUserQueries)
	repo := repositories.NewUserRepository(mockQueries)
	ctx := context.Background()
	user := database.User{ID: 1, Email: "test@example.com"}

	mockQueries.On("GetUser", ctx, int32(1)).Return(user, nil)

	result, err := repo.Get(ctx, 1)

	assert.NoError(t, err)
	assert.Equal(t, user, result)
	mockQueries.AssertExpectations(t)
}

func TestUserRepository_Create(t *testing.T) {
	mockQueries := new(MockUserQueries)
	repo := repositories.NewUserRepository(mockQueries)
	ctx := context.Background()
	user := database.User{ID: 1, Email: "test@example.com"}
	params := repositories.CreateUserParams{
		Email:    "test@example.com",
		Password: "password123",
		Type:     "admin",
	}

	mockQueries.On("CreateUser", ctx, mock.Anything).Return(user, nil)

	result, err := repo.Create(ctx, params)

	assert.NoError(t, err)
	assert.Equal(t, user, result)
	mockQueries.AssertExpectations(t)
}

func TestUserRepository_Update(t *testing.T) {
	mockQueries := new(MockUserQueries)
	repo := repositories.NewUserRepository(mockQueries)
	ctx := context.Background()
	user := database.User{ID: 1, Email: "updated@example.com"}
	params := repositories.UpdateUserParams{
		Email:    "updated@example.com",
		Password: "newpassword",
	}

	mockQueries.On("UpdateUser", ctx, mock.Anything).Return(user, nil)

	result, err := repo.Update(ctx, 1, params)

	assert.NoError(t, err)
	assert.Equal(t, user, result)
	mockQueries.AssertExpectations(t)
}

func TestUserRepository_Delete(t *testing.T) {
	mockQueries := new(MockUserQueries)
	repo := repositories.NewUserRepository(mockQueries)
	ctx := context.Background()

	mockQueries.On("DeleteUser", ctx, int32(1)).Return(nil)

	err := repo.Delete(ctx, 1)

	assert.NoError(t, err)
	mockQueries.AssertExpectations(t)
}

