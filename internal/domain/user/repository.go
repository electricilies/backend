package user

import (
	"backend/internal/domain/cart"
	"context"
)

type Repository interface {
	Get(ctx context.Context, id string) (*Model, error)
	List(ctx context.Context) ([]*Model, error)
	Create(ctx context.Context, user *Model) (*Model, error)
	Update(ctx context.Context, user *Model, queryParams *QueryParams) error
	Delete(ctx context.Context, id string) error
	GetCart(ctx context.Context, id string) (*cart.Model, error)
}
