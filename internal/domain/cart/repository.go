package cart

import (
	"context"

	"backend/internal/domain/param"
)

type Repository interface {
	GetCartByUser(ctx context.Context, userID string, pagination *param.Params) (*Model, error)
}
