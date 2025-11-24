package repositorypostgres

import (
	"context"

	"backend/internal/domain"
	"backend/internal/infrastructure/repositorypostgres/sqlc"

	"github.com/google/uuid"
)

type PostgresCart struct {
	queries *sqlc.Queries
}

var _ domain.CartRepository = (*PostgresCart)(nil)

func ProvidePostgresCart(q *sqlc.Queries) *PostgresCart {
	return &PostgresCart{queries: q}
}

func (r *PostgresCart) Get(ctx context.Context, id uuid.UUID) (*domain.Cart, error) {
	panic("implement me")
}

func (r *PostgresCart) Save(ctx context.Context, cart domain.Cart) error {
	panic("implement me")
}
