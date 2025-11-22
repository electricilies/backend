package repository

import (
	"context"

	"backend/internal/domain"
	"backend/internal/infrastructure/repository/postgres"

	"github.com/google/uuid"
)

type PostgresCarts struct {
	querier postgres.Querier
}

func NewPostgresCarts(q postgres.Querier) *PostgresCarts {
	return &PostgresCarts{querier: q}
}

func (r *PostgresCarts) Get(ctx context.Context, id uuid.UUID) (*domain.Cart, error) {
	panic("implement me")
}

func (r *PostgresCarts) Save(ctx context.Context, cart domain.Cart) (*domain.Cart, error) {
	panic("implement me")
}
