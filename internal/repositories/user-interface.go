package repositories

import (
	"context"
	"patient-appointment-demo-go/internal/database"
)

type UserRepositoryInterface interface {
	GetAll(ctx context.Context) ([]database.User, error)
	Get(ctx context.Context, id int32) (database.User, error)
    GetByEmail(ctx context.Context, email string) (database.User, error)
	Create(ctx context.Context, data CreateUserParams) (database.User, error)
	Update(ctx context.Context, id int32, data UpdateUserParams) (database.User, error)
	Delete(ctx context.Context, id int32) error
}

