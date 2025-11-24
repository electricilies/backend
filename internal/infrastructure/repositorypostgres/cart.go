package repositorypostgres

import (
	"context"

	"backend/internal/domain"
	"backend/internal/infrastructure/repositorypostgres/sqlc"

	"github.com/google/uuid"
)

type Cart struct {
	queries *sqlc.Queries
}

var _ domain.CartRepository = (*Cart)(nil)

func ProvideCart(q *sqlc.Queries) *Cart {
	return &Cart{queries: q}
}

func (r *Cart) Get(ctx context.Context, id uuid.UUID) (*domain.Cart, error) {
	panic("implement me")
}

func (r *Cart) Save(ctx context.Context, cart domain.Cart) error {
	panic("implement me")
}
