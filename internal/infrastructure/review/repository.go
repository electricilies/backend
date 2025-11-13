package review

import (
	"context"

	"backend/internal/domain/review"
)

type repositoryImpl struct{}

func NewRepository() review.Repository {
	return &repositoryImpl{}
}

func (r *repositoryImpl) ListReviewsByProductID(ctx context.Context, productID int, queryParams *review.QueryParams) (*review.Pagination, error) {
	return &review.Pagination{}, nil
}
