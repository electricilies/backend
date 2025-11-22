package repository

import (
	"context"

	"backend/internal/domain"
	"backend/internal/infrastructure/repository/postgres"

	"github.com/google/uuid"
)

type PostgresReview struct {
	querier postgres.Querier
}

var _ domain.ReviewRepository = (*PostgresReview)(nil)

func ProvidePostgresReview(q postgres.Querier) *PostgresReview {
	return &PostgresReview{querier: q}
}

func (r *PostgresReview) List(ctx context.Context, orderItemIDs *[]uuid.UUID, productVariantID *uuid.UUID, userIDs *[]uuid.UUID, deleted domain.DeletedParam, limit int, offset int) (*[]domain.Review, error) {
	panic("implement me")
}

func (r *PostgresReview) Count(ctx context.Context, orderItemIDs *[]uuid.UUID, productVariantID *uuid.UUID, userIDs *[]uuid.UUID, deleted domain.DeletedParam) (*int, error) {
	panic("implement me")
}

func (r *PostgresReview) Get(ctx context.Context, id uuid.UUID) (*domain.Review, error) {
	panic("implement me")
}

func (r *PostgresReview) Save(ctx context.Context, review domain.Review) (*domain.Review, error) {
	panic("implement me")
}
