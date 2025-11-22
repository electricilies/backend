package repository

import (
	"context"

	"backend/internal/domain"
	"backend/internal/infrastructure/repository/postgres"

	"github.com/google/uuid"
)

type PostgresCategory struct {
	queries *postgres.Queries
}

var _ domain.CategoryRepository = (*PostgresCategory)(nil)

func ProvidePostgresCategory(q *postgres.Queries) *PostgresCategory {
	return &PostgresCategory{queries: q}
}

func (r *PostgresCategory) List(ctx context.Context, search *string, limit int, offset int) (*[]domain.Category, error) {
	panic("implement me")
}

func (r *PostgresCategory) Count(ctx context.Context) (*int, error) {
	panic("implement me")
}

func (r *PostgresCategory) Get(ctx context.Context, id uuid.UUID) (*domain.Category, error) {
	panic("implement me")
}

func (r *PostgresCategory) Save(ctx context.Context, category domain.Category) error {
	panic("implement me")
}
