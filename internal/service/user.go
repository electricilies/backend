package service

import (
	"context"

	"backend/internal/domain"
)

type CreateUserParam struct {
	ID string `json:"id" binding:"required"`
}

type UpdateUserParam struct {
	FirstName   *string `json:"firstName" binding:"required"`
	LastName    *string `json:"lastName" binding:"required"`
	Email       *string `json:"email" binding:"required,email"`
	DateOfBirth *string `json:"dateOfBirth,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	Address     *string `json:"address,omitempty"`
}

type ListUsersParam struct {
	Limit  int
	Offset int
}
type GetUserCartParam struct {
	CartID string `json:"cartId" binding:"required"`
}

type User interface {
	Create(context.Context, CreateUserParam) (*domain.User, error)
	Update(context.Context, UpdateUserParam) (*domain.User, error)
	List(context.Context, ListUsersParam) (*domain.Pagination[domain.User], error)
	GetCart(context.Context, GetUserCartParam) (*domain.Cart, error)
}

type UserImpl struct{}

func ProvideUser() *UserImpl {
	return &UserImpl{}
}

var _ User = &UserImpl{}

func (s *UserImpl) Create(ctx context.Context, param CreateUserParam) (*domain.User, error) {
	return nil, nil
}

func (s *UserImpl) Update(ctx context.Context, param UpdateUserParam) (*domain.User, error) {
	return nil, nil
}

func (s *UserImpl) List(ctx context.Context, param ListUsersParam) (*domain.Pagination[domain.User], error) {
	return nil, nil
}

func (s *UserImpl) GetCart(ctx context.Context, param GetUserCartParam) (*domain.Cart, error) {
	return nil, nil
}
