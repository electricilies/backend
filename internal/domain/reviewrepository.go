package domain

import (
	"context"

	"github.com/google/uuid"
)

type ReviewRepository interface {
	Get(
		ctx context.Context,
		id uuid.UUID,
	) (*Review, error)

	List(
		ctx context.Context,
		orderItemIDs *[]uuid.UUID,
		productVariantID *uuid.UUID,
		userIDs *[]uuid.UUID,
		deleted DeletedParam,
		limit int,
		offset int,
	) (*[]Review, error)

	Count(
		ctx context.Context,
		orderItemIDs *[]uuid.UUID,
		productVariantID *uuid.UUID,
		userIDs *[]uuid.UUID,
		deleted DeletedParam,
	) (*int, error)

	Save(
		ctx context.Context,
		review *Review,
	) (*Review, error)

	Remove(
		ctx context.Context,
		id uuid.UUID,
	) error
}
