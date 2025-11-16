package user

import (
	"context"

	"backend/internal/domain/cart"
)

type Repository interface {
	Get(context.Context, string) (*Model, error)
	List(context.Context) ([]*Model, error)
	Create(context.Context, *Model) (*Model, error)
	Update(context.Context, *Model, *QueryParams) error
	Delete(context.Context, string) error
	GetCart(context.Context, string) (*cart.Model, error)
}
