package repository

import (
	"context"

	"backend/internal/domain"
	"backend/internal/infrastructure/repository/postgres"

	"github.com/google/uuid"
)

type PostgresReview struct {
	queries *postgres.Queries
}

var _ domain.ReviewRepository = (*PostgresReview)(nil)

func ProvidePostgresReview(q *postgres.Queries) *PostgresReview {
	return &PostgresReview{queries: q}
}

func (r *PostgresReview) List(
	ctx context.Context,
	orderItemIDs *[]uuid.UUID,
	productVariantID *uuid.UUID,
	userIDs *[]uuid.UUID,
	deleted domain.DeletedParam,
	limit int,
	offset int,
) (*[]domain.Review, error) {
	panic("implement me")
}

func (r *PostgresReview) Count(ctx context.Context, orderItemIDs *[]uuid.UUID, productVariantID *uuid.UUID, userIDs *[]uuid.UUID, deleted domain.DeletedParam) (*int, error) {
	panic("implement me")
}

func (r *PostgresReview) Get(ctx context.Context, id uuid.UUID) (*domain.Review, error) {
	panic("implement me")
}

func (r *PostgresReview) Save(ctx context.Context, review domain.Review) error {
	panic("implement me")
}
