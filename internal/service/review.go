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
	Limit   int
	Offset  int
	Deleted string
}

type GetReviewParam struct {
	ReviewID int `json:"reviewId" binding:"required"`
}

type DeleteReviewParam struct {
	ReviewID int `json:"reviewId" binding:"required"`
}

type Review interface {
	CreateReview(context.Context, CreateReviewParam) (*domain.Review, error)
	UpdateReview(context.Context, UpdateReviewParam) (*domain.Review, error)
	ListReviews(context.Context, ListReviewsParam) (*domain.DataPagination, error)
	GetReview(context.Context, int) (*domain.Review, error)
	DeleteReview(context.Context, int) error
}

type ReviewImpl struct{}

func ProvideReview() *ReviewImpl {
	return &ReviewImpl{}
}

var _ Review = &ReviewImpl{}

func (s *ReviewImpl) CreateReview(ctx context.Context, param CreateReviewParam) (*domain.Review, error) {
	return nil, nil
}

func (s *ReviewImpl) UpdateReview(ctx context.Context, param UpdateReviewParam) (*domain.Review, error) {
	return nil, nil
}

func (s *ReviewImpl) ListReviews(ctx context.Context, param ListReviewsParam) (*domain.DataPagination, error) {
	return nil, nil
}

func (s *ReviewImpl) GetReview(ctx context.Context, reviewID int) (*domain.Review, error) {
	return nil, nil
}

func (s *ReviewImpl) DeleteReview(ctx context.Context, reviewID int) error {
	return nil
}

