package repositories

import (
	"context"
	"patient-appointment-demo-go/internal/database"
)

type UserRepository struct {
	queries database.Queries
}

type CreateUserParams struct {
	Email    string
	Type     string
	Password string
}

type UpdateUserParams struct {
	Email    string
	Password string
}

func NewUserRepository(queries database.Queries) UserRepositoryInterface {
	return &UserRepository{
		queries: queries,
	}
}

func (r *UserRepository) GetAll(ctx context.Context) ([]database.User, error) {

	res, err := r.queries.GetAllUsers(ctx)

	return res, err
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (database.User, error) {

	res, err := r.queries.GetUserByEmail(ctx, email)

	return res, err
}

func (r *UserRepository) Get(ctx context.Context, id int32) (database.User, error) {

	res, err := r.queries.GetUser(ctx, id)

	return res, err
}

func (r *UserRepository) Create(ctx context.Context, data CreateUserParams) (database.User, error) {

	res, err := r.queries.CreateUser(ctx, database.CreateUserParams{
		Email:    data.Email,
		Password: data.Password,
		Type:     data.Type,
	})

	return res, err
}

func (r *UserRepository) Update(ctx context.Context, id int32, data UpdateUserParams) (database.User, error) {

	res, err := r.queries.UpdateUser(ctx, database.UpdateUserParams{
		ID:       id,
		Email:    data.Email,
		Password: data.Password,
	})

	return res, err
}

func (r *UserRepository) Delete(ctx context.Context, id int32) error {

	err := r.queries.DeleteUser(ctx, id)

	return err
}
