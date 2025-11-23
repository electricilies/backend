package domain

import "github.com/google/uuid"

type ReviewService interface {
	Create(
		orderItemID uuid.UUID,
		userID uuid.UUID,
		rating int,
		content *string,
		ImageURL *string,
	) (*Review, error)

	Update(
		review *Review,
		userID uuid.UUID,
		rating *int,
		content *string,
		imageURL *string,
	) error
}
