package domain

import (
	"context"
)

type ReviewRepository interface {
	Get(
		ctx context.Context,
		id int,
	) (*Review, error)

	List(
		ctx context.Context,
		orderItemIDs *[]int,
		productVariantID *int,
		userIDs *[]int,
		deleted string,
		limit int,
		offset int,
	) (*[]Review, error)

	Count(
		ctx context.Context,
		orderItemIDs *[]int,
		productVariantID *int,
		userIDs *[]int,
		deleted string,
		limit int,
		offset int,
	) (*int, error)

	Save(
		ctx context.Context,
		review *Review,
	) (*Review, error)

	Remove(
		ctx context.Context,
		id int,
	) error
}
