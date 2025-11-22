package repository

import (
	"context"

	"backend/internal/domain"
	"backend/internal/infrastructure/repository/postgres"

	"github.com/google/uuid"
)

type PostgresProducts struct {
	querier postgres.Querier
}

func NewPostgresProducts(q postgres.Querier) *PostgresProducts {
	return &PostgresProducts{querier: q}
}

func (r *PostgresProducts) List(ctx context.Context, ids *[]uuid.UUID, search *string, min_price *int64, max_price *int64, rating *float64, category_ids *[]uuid.UUID, deleted domain.DeletedParam, sort_rating *string, sort_price *string, limit int, offset int) (*[]domain.Product, error) {
	panic("implement me")
}

func (r *PostgresProducts) Count(ctx context.Context, ids *[]uuid.UUID, min_price *int64, max_price *int64, rating *float64, category_ids *[]uuid.UUID, deleted domain.DeletedParam) (*int, error) {
	panic("implement me")
}

func (r *PostgresProducts) Get(ctx context.Context, productID uuid.UUID) (*domain.Product, error) {
	panic("implement me")
}

func (r *PostgresProducts) Save(ctx context.Context, product domain.Product) (*domain.Product, error) {
	panic("implement me")
}
