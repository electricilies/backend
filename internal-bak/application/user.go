package application

import (
	"context"

	"backend/internal/domain/cart"
	"backend/internal/domain/user"
)

type User interface {
	Get(context.Context, string) (*user.Model, error)
	List(context.Context) ([]*user.Model, error)
	Create(context.Context, *user.Model) (*user.Model, error)
	Update(context.Context, *user.Model, *user.QueryParams) error
	Delete(context.Context, string) error
	GetCart(context.Context, string) (*cart.Model, error)
}

type UserImpl struct {
	userRepo    user.Repository
	userService user.Service
}

func NewUser(userRepo user.Repository, userService user.Service) User {
	return &UserImpl{
		userRepo:    userRepo,
		userService: userService,
	}
}

func ProvideUser(
	userRepo user.Repository,
	userService user.Service,
) *UserImpl {
	return &UserImpl{
		userRepo:    userRepo,
		userService: userService,
	}
}

func (a *UserImpl) Get(
	ctx context.Context,
	id string,
) (*user.Model, error) {
	return a.userRepo.Get(ctx, id)
}

func (a *UserImpl) List(ctx context.Context) ([]*user.Model, error) {
	return a.userRepo.List(ctx)
}

func (a *UserImpl) Create(ctx context.Context, model *user.Model) (*user.Model, error) {
	return a.userRepo.Create(ctx, model)
}

func (a *UserImpl) Update(
	ctx context.Context,
	model *user.Model,
	queryParams *user.QueryParams,
) error {
	return a.userRepo.Update(ctx, model, queryParams)
}

func (a *UserImpl) Delete(ctx context.Context, id string) error {
	return a.userRepo.Delete(ctx, id)
}

func (a *UserImpl) GetCart(ctx context.Context, id string) (*cart.Model, error) {
	return a.userRepo.GetCart(ctx, id)
}
