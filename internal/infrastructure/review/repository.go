package review

import (
	"context"

	"backend/internal/domain/review"
	"backend/internal/infrastructure/persistence/postgres"
)

type repositoryImpl struct {
	db *postgres.Queries
}

func NewRepository(
	db *postgres.Queries,
) review.Repository {
	return &repositoryImpl{
		db: db,
	}
}

func (r *repositoryImpl) ListReviews(ctx context.Context, productID int, queryParams *review.QueryParams) (*review.Pagination, error) {
	return &review.Pagination{}, nil
}
