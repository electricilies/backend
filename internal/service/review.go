package service

import (
	"backend/internal/domain"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/hashicorp/go-multierror"
)

type Review struct {
	validate *validator.Validate
}

func ProvideReview(
	validate *validator.Validate,
) *Review {
	return &Review{
		validate: validate,
	}
}

var _ domain.ReviewService = &Review{}

func (r *Review) Create(
	orderItemID uuid.UUID,
	userID uuid.UUID,
	rating int,
	content *string,
	imageURL *string,
) (*domain.Review, error) {
	orderItem := domain.OrderItem{ID: orderItemID}
	id, err := uuid.NewV7()
	if err != nil {
		return nil, multierror.Append(domain.ErrInternal, err)
	}
	review := &domain.Review{
		ID:        id,
		Rating:    rating,
		Content:   content,
		OrderItem: &orderItem,
		ImageURL:  imageURL,
	}
	if err := r.validate.Struct(review); err != nil {
		return nil, multierror.Append(domain.ErrInvalid, err)
	}
	return review, nil
}

func (r *Review) Update(
	review *domain.Review,
	rating *int,
	content *string,
	imageURL *string,
) error {
	if rating != nil {
		review.Rating = *rating
	}
	if content != nil {
		review.Content = content
	}
	if imageURL != nil {
		review.ImageURL = imageURL
	}
	if err := r.validate.Struct(review); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}

