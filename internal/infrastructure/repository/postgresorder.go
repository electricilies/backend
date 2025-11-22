package repository

import (
	"context"

	"backend/internal/domain"
	"backend/internal/infrastructure/repository/postgres"

	"github.com/google/uuid"
)

type PostgresOrder struct {
	querier postgres.Querier
}

var _ domain.OrderRepository = (*PostgresOrder)(nil)


func ProvidePostgresOrder(q postgres.Querier) *PostgresOrder {
	return &PostgresOrder{querier: q}
}

func (r *PostgresOrder) List(ctx context.Context, ids *[]uuid.UUID, search *string, deleted domain.DeletedParam, limit int, offset int) (*[]domain.Order, error) {
	panic("implement me")
}

func (r *PostgresOrder) Count(ctx context.Context, ids *[]uuid.UUID, deleted domain.DeletedParam) (*int, error) {
	panic("implement me")
}

func (r *PostgresOrder) Get(ctx context.Context, id uuid.UUID) (*domain.Order, error) {
	panic("implement me")
}

func (r *PostgresOrder) Save(ctx context.Context, order domain.Order) (*domain.Order, error) {
	panic("implement me")
}
