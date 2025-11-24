package repositorypostgres

import (
	"context"

	"backend/internal/domain"
	"backend/internal/infrastructure/repositorypostgres/sqlc"

	"github.com/google/uuid"
)

type Review struct {
	queries *sqlc.Queries
}

var _ domain.ReviewRepository = (*Review)(nil)

func ProvideReview(q *sqlc.Queries) *Review {
	return &Review{queries: q}
}

func (r *Review) List(
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

func (r *Review) Count(ctx context.Context, orderItemIDs *[]uuid.UUID, productVariantID *uuid.UUID, userIDs *[]uuid.UUID, deleted domain.DeletedParam) (*int, error) {
	panic("implement me")
}

func (r *Review) Get(ctx context.Context, id uuid.UUID) (*domain.Review, error) {
	panic("implement me")
}

func (r *Review) Save(ctx context.Context, review domain.Review) error {
	panic("implement me")
}
