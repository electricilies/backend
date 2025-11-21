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
	orderItem domain.OrderItem,
	rating int,
	content *string,
	ImageURL *string,
) (*domain.Review, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, multierror.Append(domain.ErrInternal, err)
	}
	review := &domain.Review{
		ID:        id,
		Rating:    rating,
		Content:   content,
		OrderItem: &orderItem,
		ImageURL:  ImageURL,
	}
	if err := r.validate.Struct(review); err != nil {
		return nil, multierror.Append(domain.ErrInvalid, err)
	}
	return review, nil
}
