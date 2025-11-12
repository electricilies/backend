package cart

import (
	"context"

	"backend/internal/domain/params"
)

type Repository interface {
	GetCartByUser(ctx context.Context, userID string, pagination *params.Params) (*Model, error)
}
