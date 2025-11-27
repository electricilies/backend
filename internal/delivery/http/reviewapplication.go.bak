package http

import (
	"context"

	"backend/internal/domain"
)

type ReviewApplication interface {
	Create(ctx context.Context, param CreateReviewRequestDto) (*domain.Review, error)
	List(ctx context.Context, param ListReviewsRequestDto) (*PaginationResponseDto[domain.Review], error)
	Get(ctx context.Context, param GetReviewRequestDto) (*domain.Review, error)
	Update(ctx context.Context, param UpdateReviewRequestDto) (*domain.Review, error)
	Delete(ctx context.Context, param DeleteReviewRequestDto) error
}
