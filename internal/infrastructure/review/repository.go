package review

import (
	"context"

	"backend/internal/domain/review"
)

type Repository struct{}

func NewRepository() review.Repository {
	return &Repository{}
}

func (r *Repository) ListReviewsByProductID(ctx context.Context, productID int, queryParams *review.QueryParams) (*review.Pagination, error) {
	return &review.Pagination{}, nil
}
