package user

import (
	"backend/internal/domain/user"
	"backend/internal/infrastructure/errors"
	"backend/internal/infrastructure/presistence/postgres"
	"context"

	"github.com/google/uuid"
)

type repositoryImpl struct {
	db *postgres.Queries
}

func NewRepository(query *postgres.Queries) user.Repository {
	return &repositoryImpl{
		db: query,
	}
}

func (r *repositoryImpl) Get(ctx context.Context, id string) (*user.User, error) {
	u, err := r.db.GetUser(ctx, postgres.GetUserParams{ID: uuid.MustParse(id)})
	if err != nil {
		return nil, errors.ToDomainError(err)
	}

	return ToDomain(u), nil
}

func (r *repositoryImpl) List(ctx context.Context) ([]*user.User, error) {
	users, err := r.db.ListUsers(ctx)
	if err != nil {
		return nil, errors.ToDomainError(err)
	}

	result := make([]*user.User, len(users))
	for i, u := range users {
		result[i] = ToDomain(u)
	}

	return result, nil
}

func (r *repositoryImpl) Create(ctx context.Context, u *user.User) (*user.User, error) {
	createdUser, err := r.db.CreateUser(ctx, ToCreateUserParams(u))

	if err != nil {
		return nil, errors.ToDomainError(err)
	}

	return ToDomain(createdUser), nil
}

func (r *repositoryImpl) Update(ctx context.Context, u *user.User) error {
	return errors.ToDomainError(r.db.UpdateUser(ctx, ToUpdateUserParams(u)))
}

func (r *repositoryImpl) Delete(ctx context.Context, id string) error {
	return errors.ToDomainError(r.db.DeleteUser(ctx, postgres.DeleteUserParams{
		ID: uuid.MustParse(id),
	}))
}
