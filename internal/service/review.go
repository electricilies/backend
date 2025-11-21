package serviceimpl

import (
	"context"

	"backend/internal/domain"
)

type Review struct{}

func ProvideReview() *Review {
	return &Review{}
}

var _ domain.ReviewService = &Review{}

func (s *Review) Create(ctx context.Context, param domain.CreateReviewParam) (*domain.Review, error) {
	panic("not implemented")
}

func (s *Review) Update(ctx context.Context, param domain.UpdateReviewParam) (*domain.Review, error) {
	panic("not implemented")
}

func (s *Review) Get(ctx context.Context, reviewID int) (*domain.Review, error) {
	panic("not implemented")
}

func (s *Review) Delete(ctx context.Context, reviewID int) error {
	panic("not implemented")
}
