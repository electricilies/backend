package application

import (
	"context"

	"backend/internal/domain/review"
)

type Review interface {
	ListReviewsByProductID(ctx context.Context, productID int, queryParams *review.QueryParams) (*review.Pagination, error)
}

type reviewApp struct {
	reviewRepo review.Repository
}

func NewReview(reviewRepo review.Repository) Review {
	return &reviewApp{
		reviewRepo: reviewRepo,
	}
}

func (a *reviewApp) ListReviewsByProductID(ctx context.Context, productID int, queryParams *review.QueryParams) (*review.Pagination, error) {
	return a.reviewRepo.ListReviewsByProductID(ctx, productID, queryParams)
}
