package service

import (
	"context"

	"backend/internal/domain"
)

type UserImpl struct{}

func ProvideUser() *UserImpl {
	return &UserImpl{}
}

var _ User = &UserImpl{}

func (s *UserImpl) Create(ctx context.Context, param CreateUserParam) (*domain.User, error) {
	panic("implement me")
}

func (s *UserImpl) Update(ctx context.Context, param UpdateUserParam) (*domain.User, error) {
	panic("implement me")
}

func (s *UserImpl) List(ctx context.Context, param ListUsersParam) (*Pagination[domain.User], error) {
	panic("implement me")
}

func (s *UserImpl) Get(ctx context.Context, param GetUserParam) (*domain.User, error) {
	panic("implement me")
}

func (s *UserImpl) GetCart(ctx context.Context, param GetUserCartParam) (*domain.Cart, error) {
	panic("implement me")
}
