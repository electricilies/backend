package application

import (
	"context"

	"backend/internal/domain/review"
)

type Review interface {
	ListReviewsByProductID(ctx context.Context, productID int, queryParams *review.QueryParams) (*review.Pagination, error)
}

type ReviewApp struct {
	reviewRepo review.Repository
}

func NewReview(reviewRepo review.Repository) Review {
	return &ReviewApp{
		reviewRepo: reviewRepo,
	}
}

func ProvideReview(
	reviewRepo review.Repository,
) *ReviewApp {
	return &ReviewApp{
		reviewRepo: reviewRepo,
	}
}

func (a *ReviewApp) ListReviewsByProductID(ctx context.Context, productID int, queryParams *review.QueryParams) (*review.Pagination, error) {
	return a.reviewRepo.ListByProduct(ctx, productID, queryParams)
}
