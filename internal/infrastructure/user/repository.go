package user

import (
	"backend/internal/domain/user"
	errormapper "backend/internal/infrastructure/error"
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
		return nil, errormapper.ToDomainError(err)
	}

	return ToDomain(u), nil
}

func (r *repositoryImpl) List(ctx context.Context) ([]*user.User, error) {
	users, err := r.db.ListUser(ctx)
	if err != nil {
		return nil, errormapper.ToDomainError(err)
	}

	result := make([]*user.User, len(users))
	for i, u := range users {
		result[i] = ToDomain(u)
	}

	return result, nil
}

func (r *repositoryImpl) Create(ctx context.Context, u *user.User) (*user.User, error) {
	createdUser, err := r.db.CreateUser(ctx, postgres.CreateUserParams{
		Name: u.Name,
	})

	if err != nil {
		return nil, errormapper.ToDomainError(err)
	}

	return ToDomain(createdUser), nil
}

func (r *repositoryImpl) Update(ctx context.Context, u *user.User) error {
	return errormapper.ToDomainError(r.db.UpdateUser(ctx, postgres.UpdateUserParams{
		ID:   uuid.MustParse(u.ID),
		Name: u.Name,
	}))
}

func (r *repositoryImpl) Delete(ctx context.Context, id string) error {
	return errormapper.ToDomainError(r.db.DeleteUser(ctx, postgres.DeleteUserParams{
		ID: uuid.MustParse(id),
	}))
}
