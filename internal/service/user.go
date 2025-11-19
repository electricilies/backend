package service

import (
	"context"

	"backend/internal/domain"
)

type User interface {
	Create(context.Context, CreateUserParam) (*domain.User, error)
	Update(context.Context, UpdateUserParam) (*domain.User, error)
	List(context.Context, ListUsersParam) (*Pagination[domain.User], error)
	Get(context.Context, GetUserParam) (*domain.User, error)
	GetCart(context.Context, GetUserCartParam) (*domain.Cart, error)
}
