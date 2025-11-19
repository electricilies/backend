package domain

import (
	"context"
)

type ReviewService interface {
	Create(context.Context, CreateReviewParam) (*Review, error)
	Update(context.Context, UpdateReviewParam) (*Review, error)
	Get(context.Context, int) (*Review, error)
	Delete(context.Context, int) error
}
