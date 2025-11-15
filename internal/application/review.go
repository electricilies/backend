package application

import (
	"context"

	"backend/internal/domain/review"
)

type Review interface {
	ListReviewsByProductID(context.Context, int, *review.QueryParams) (*review.Pagination, error)
}

type ReviewImpl struct {
	reviewRepo review.Repository
}

func NewReview(reviewRepo review.Repository) Review {
	return &ReviewImpl{
		reviewRepo: reviewRepo,
	}
}

func ProvideReview(
	reviewRepo review.Repository,
) *ReviewImpl {
	return &ReviewImpl{
		reviewRepo: reviewRepo,
	}
}

func (a *ReviewImpl) ListReviewsByProductID(
	ctx context.Context,
	productID int,
	queryParams *review.QueryParams,
) (*review.Pagination, error) {
	return a.reviewRepo.ListByProduct(ctx, productID, queryParams)
}
