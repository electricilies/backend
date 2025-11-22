package application

import (
	"context"

	"backend/internal/domain"
)

type Review interface {
	Create(ctx context.Context, param CreateReviewParam) (*domain.Review, error)
	List(ctx context.Context, param ListReviewsParam) (*Pagination[domain.Review], error)
	Get(ctx context.Context, param GetReviewParam) (*domain.Review, error)
	Update(ctx context.Context, param UpdateReviewParam) (*domain.Review, error)
	Delete(ctx context.Context, param DeleteReviewParam) error
}
