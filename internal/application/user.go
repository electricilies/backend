package application

import (
	userdomain "backend/internal/domain/user"
	"context"
)

type User interface {
	Get(ctx context.Context, id string) (*userdomain.User, error)
	List(ctx context.Context) ([]*userdomain.User, error)
	Create(ctx context.Context, u *userdomain.User) (*userdomain.User, error)
	Update(ctx context.Context, u *userdomain.User) error
	Delete(ctx context.Context, id string) error
}

type userApp struct {
	userRepo    userdomain.Repository
	userService userdomain.Service
}

func NewUser(userRepo userdomain.Repository, userService userdomain.Service) User {
	return &userApp{
		userRepo:    userRepo,
		userService: userService,
	}
}

func (a *userApp) Get(ctx context.Context, id string) (*userdomain.User, error) {
	u, err := a.userRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (a *userApp) List(ctx context.Context) ([]*userdomain.User, error) {
	return a.userRepo.List(ctx)
}

func (a *userApp) Create(ctx context.Context, u *userdomain.User) (*userdomain.User, error) {
	return a.userRepo.Create(ctx, u)
}

func (a *userApp) Update(ctx context.Context, u *userdomain.User) error {
	_, err := a.userRepo.Get(ctx, u.ID)
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
