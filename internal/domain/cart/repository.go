package cart

import (
	"context"
)

type Repository interface {
	GetCartByUser(ctx context.Context, userID string, queryParams *QueryParams) (*Model, error)
}
