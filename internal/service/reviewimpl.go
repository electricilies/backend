package service

import (
	"context"

	"backend/internal/domain"
)

type ReviewImpl struct{}

func ProvideReview() *ReviewImpl {
	return &ReviewImpl{}
}

var _ Review = &ReviewImpl{}

func (s *ReviewImpl) Create(ctx context.Context, param CreateReviewParam) (*domain.Review, error) {
	panic("not implemented")
}

func (s *ReviewImpl) Update(ctx context.Context, param UpdateReviewParam) (*domain.Review, error) {
	panic("not implemented")
}

func (s *ReviewImpl) List(ctx context.Context, param ListReviewsParam) (*Pagination[domain.Review], error) {
	panic("not implemented")
}

func (s *ReviewImpl) Get(ctx context.Context, reviewID int) (*domain.Review, error) {
	panic("not implemented")
}

func (s *ReviewImpl) Delete(ctx context.Context, reviewID int) error {
	panic("not implemented")
}
