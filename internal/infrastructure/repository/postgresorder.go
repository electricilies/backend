package repository

import (
	"context"

	"backend/internal/domain"
	"backend/internal/infrastructure/repository/postgres"

	"github.com/google/uuid"
)

type PostgresOrders struct {
	querier postgres.Querier
}

func NewPostgresOrders(q postgres.Querier) *PostgresOrders {
	return &PostgresOrders{querier: q}
}

func (r *PostgresOrders) List(ctx context.Context, ids *[]uuid.UUID, search *string, deleted domain.DeletedParam, limit int, offset int) (*[]domain.Order, error) {
	panic("implement me")
}

func (r *PostgresOrders) Count(ctx context.Context, ids *[]uuid.UUID, deleted domain.DeletedParam) (*int, error) {
	panic("implement me")
}

func (r *PostgresOrders) Get(ctx context.Context, id uuid.UUID) (*domain.Order, error) {
	panic("implement me")
}

func (r *PostgresOrders) Save(ctx context.Context, order domain.Order) (*domain.Order, error) {
	panic("implement me")
}
