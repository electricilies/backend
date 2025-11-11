package cart

import (
	"context"

	"backend/internal/domain/pagination"
)

type Repository interface {
	GetCartByUser(ctx context.Context, userID string, pagination *pagination.Params) (*Model, error)
}
