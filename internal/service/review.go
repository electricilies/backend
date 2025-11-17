package service

import (
	"context"

	"backend/internal/domain"
)

type CreateReviewParam struct {
	OrderItemID int    `json:"orderItemId" binding:"required"`
	UserID      string `json:"userId" binding:"required"`
	Rating      int    `json:"rating" binding:"required,min=1,max=5"`
	Content     string `json:"content,omitempty"`
	ImageURL    string `json:"imageUrl,omitempty"`
}

type UpdateReviewParam struct {
	Rating   int    `json:"rate,omitempty"`
	Content  string `json:"content,omitempty"`
	ImageURL string `json:"imageUrl,omitempty"`
}

type ListReviewsParam struct {
	PaginationParam
	Deleted string
}

type GetReviewParam struct {
	ReviewID int `json:"reviewId" binding:"required"`
}

type DeleteReviewParam struct {
	ReviewID int `json:"reviewId" binding:"required"`
}

type Review interface {
	Create(context.Context, CreateReviewParam) (*domain.Review, error)
	Update(context.Context, UpdateReviewParam) (*domain.Review, error)
	List(context.Context, ListReviewsParam) (*domain.Pagination[domain.Review], error)
	Get(context.Context, int) (*domain.Review, error)
	Delete(context.Context, int) error
}

type ReviewImpl struct{}

func ProvideReview() *ReviewImpl {
	return &ReviewImpl{}
}

var _ Review = &ReviewImpl{}

func (s *ReviewImpl) Create(ctx context.Context, param CreateReviewParam) (*domain.Review, error) {
	return nil, nil
}

func (s *ReviewImpl) Update(ctx context.Context, param UpdateReviewParam) (*domain.Review, error) {
	return nil, nil
}

func (s *ReviewImpl) List(ctx context.Context, param ListReviewsParam) (*domain.Pagination[domain.Review], error) {
	return nil, nil
}

func (s *ReviewImpl) Get(ctx context.Context, reviewID int) (*domain.Review, error) {
	return nil, nil
}

func (s *ReviewImpl) Delete(ctx context.Context, reviewID int) error {
	return nil
}
