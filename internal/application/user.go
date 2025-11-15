package application

import (
	"context"

	"backend/internal/domain/cart"
	"backend/internal/domain/user"
)

type User interface {
	Get(ctx context.Context, id string) (*user.Model, error)
	List(ctx context.Context) ([]*user.Model, error)
	Create(ctx context.Context, u *user.Model) (*user.Model, error)
	Update(ctx context.Context, u *user.Model, queryParams *user.QueryParams) error
	Delete(ctx context.Context, id string) error
	GetCart(ctx context.Context, id string) (*cart.Model, error)
}

type UserApp struct {
	userRepo    user.Repository
	userService user.Service
}

func NewUser(userRepo user.Repository, userService user.Service) User {
	return &UserApp{
		userRepo:    userRepo,
		userService: userService,
	}
}

func ProvideUser(
	userRepo user.Repository,
	userService user.Service,
) *UserApp {
	return &UserApp{
		userRepo:    userRepo,
		userService: userService,
	}
}

func (a *UserApp) Get(
	ctx context.Context,
	id string,
) (*user.Model, error) {
	return a.userRepo.Get(ctx, id)
}

func (a *UserApp) List(ctx context.Context) ([]*user.Model, error) {
	return a.userRepo.List(ctx)
}

func (a *UserApp) Create(ctx context.Context, u *user.Model) (*user.Model, error) {
	return a.userRepo.Create(ctx, u)
}

func (a *UserApp) Update(
	ctx context.Context,
	u *user.Model,
	queryParams *user.QueryParams,
) error {
	return a.userRepo.Update(ctx, u, queryParams)
}

func (a *UserApp) Delete(ctx context.Context, id string) error {
	return a.userRepo.Delete(ctx, id)
}

func (a *UserApp) GetCart(ctx context.Context, id string) (*cart.Model, error) {
	return a.userRepo.GetCart(ctx, id)
}
