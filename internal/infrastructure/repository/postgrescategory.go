package repository

import (
	"context"

	"backend/internal/domain"
	"backend/internal/infrastructure/repository/postgres"

	"github.com/google/uuid"
)

type PostgresCategories struct {
	querier postgres.Querier
}

func NewPostgresCategories(q postgres.Querier) *PostgresCategories {
	return &PostgresCategories{querier: q}
}

func (r *PostgresCategories) List(ctx context.Context, search *string, limit int, offset int) (*[]domain.Category, error) {
	panic("implement me")
}

func (r *PostgresCategories) Count(ctx context.Context) (*int, error) {
	panic("implement me")
}

func (r *PostgresCategories) Get(ctx context.Context, id uuid.UUID) (*domain.Category, error) {
	panic("implement me")
}

func (r *PostgresCategories) Save(ctx context.Context, category domain.Category) (*domain.Category, error) {
	panic("implement me")
}
