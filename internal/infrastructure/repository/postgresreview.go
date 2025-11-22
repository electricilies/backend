package repository

import (
	"context"

	"backend/internal/domain"
	"backend/internal/infrastructure/repository/postgres"

	"github.com/google/uuid"
)

type PostgresReviews struct {
	querier postgres.Querier
}

func NewPostgresReviews(q postgres.Querier) *PostgresReviews {
	return &PostgresReviews{querier: q}
}

func (r *PostgresReviews) List(ctx context.Context, orderItemIDs *[]uuid.UUID, productVariantID *uuid.UUID, userIDs *[]uuid.UUID, deleted domain.DeletedParam, limit int, offset int) (*[]domain.Review, error) {
	panic("implement me")
}

func (r *PostgresReviews) Count(ctx context.Context, orderItemIDs *[]uuid.UUID, productVariantID *uuid.UUID, userIDs *[]uuid.UUID, deleted domain.DeletedParam) (*int, error) {
	panic("implement me")
}

func (r *PostgresReviews) Get(ctx context.Context, id uuid.UUID) (*domain.Review, error) {
	panic("implement me")
}

func (r *PostgresReviews) Save(ctx context.Context, review domain.Review) (*domain.Review, error) {
	panic("implement me")
}
