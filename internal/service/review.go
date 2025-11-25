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
	orderID uuid.UUID,
	orderItemID uuid.UUID,
	rating int,
	content *string,
	imageURL *string,
) (*domain.Review, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, multierror.Append(domain.ErrInternal, err)
	}
	review := &domain.Review{
		ID:          id,
		Rating:      rating,
		Content:     content,
		OrderID:     orderID,
		OrderItemID: orderItemID,
		ImageURL:    imageURL,
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
	updated := false
	if rating != nil {
		review.Rating = *rating
		updated = true
	}
	if content != nil {
		review.Content = content
		updated = true
	}
	if imageURL != nil {
		review.ImageURL = imageURL
		updated = true
	}
	if !updated {
		return nil
	}
	if err := r.validate.Struct(review); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}
