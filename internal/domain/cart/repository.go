package cart

import (
	"context"
)

type Repository interface {
	Get(ctx context.Context, id int, queryParams *QueryParams) (*Model, error)
	AddItem(ctx context.Context, cartItem *ItemModel) (*ItemModel, error)
	UpdateItem(ctx context.Context, cartItem *ItemModel, id string) (*ItemModel, error)
	RemoveItem(ctx context.Context, id string) error
}
