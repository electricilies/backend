package domain

import "github.com/google/uuid"

type ReviewService interface {
	Create(
		orderID uuid.UUID,
		orderItemID uuid.UUID,
		rating int,
		content *string,
		ImageURL *string,
	) (*Review, error)

	Update(
		review *Review,
		rating *int,
		content *string,
		imageURL *string,
	) error
}
