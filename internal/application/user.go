package application

import (
	"context"

	"backend/internal/domain/user"
)

type User interface {
	Get(ctx context.Context, id string) (*user.Model, error)
	List(ctx context.Context) ([]*user.Model, error)
	Create(ctx context.Context, u *user.Model) (*user.Model, error)
	Update(ctx context.Context, u *user.Model) error
	Delete(ctx context.Context, id string) error
}

type userApp struct {
	userRepo    user.Repository
	userService user.Service
}

func NewUser(userRepo user.Repository, userService user.Service) User {
	return &userApp{
		userRepo:    userRepo,
		userService: userService,
	}
}

func (a *userApp) Get(ctx context.Context, id string) (*user.Model, error) {
	u, err := a.userRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (a *userApp) List(ctx context.Context) ([]*user.Model, error) {
	return a.userRepo.List(ctx)
}

func (a *userApp) Create(ctx context.Context, u *user.Model) (*user.Model, error) {
	return a.userRepo.Create(ctx, u)
}

func (a *userApp) Update(ctx context.Context, u *user.Model) error {
	_, err := a.userRepo.Get(ctx, u.ID.String())
	if err != nil {
		return err
	}
	return a.userRepo.Update(ctx, u)
}

func (a *userApp) Delete(ctx context.Context, id string) error {
	_, err := a.userRepo.Get(ctx, id)
	if err != nil {
		return err
	}

	return a.userRepo.Delete(ctx, id)
}
