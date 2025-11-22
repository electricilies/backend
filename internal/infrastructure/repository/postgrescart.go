package repository

import (
	"context"

	"backend/internal/domain"
	"backend/internal/infrastructure/repository/postgres"

	"github.com/google/uuid"
)

type PostgresCart struct {
	querier postgres.Querier
}

var _ domain.CartRepository = (*PostgresCart)(nil)

func ProvidePostgresCart(q postgres.Querier) *PostgresCart {
	return &PostgresCart{querier: q}
}

func (r *PostgresCart) Get(ctx context.Context, id uuid.UUID) (*domain.Cart, error) {
	panic("implement me")
}

func (r *PostgresCart) Save(ctx context.Context, cart domain.Cart) error {
	panic("implement me")
}
