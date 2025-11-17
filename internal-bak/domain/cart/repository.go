package cart

import (
	"context"
)

type Repository interface {
	Get(context.Context, int, *QueryParams) (*Model, error)
	AddItem(context.Context, *ItemModel) (*ItemModel, error)
	UpdateItem(context.Context, *ItemModel, string) (*ItemModel, error)
	RemoveItem(context.Context, string) error
}
