package application

import (
	"context"

	"backend/internal/domain"
)

type Review interface {
	Create(context.Context, CreateReviewParam) (*domain.Review, error)
	List(context.Context, ListReviewsParam) (*Pagination[domain.Review], error)
	Get(context.Context, int) (*domain.Review, error)
	Update(context.Context, UpdateReviewParam) (*domain.Review, error)
	Delete(context.Context, int) error
}
