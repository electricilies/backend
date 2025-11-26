package domain

import (
	"context"

	"github.com/google/uuid"
)

type ProductRepository interface {
	List(
		ctx context.Context,
		ids *[]uuid.UUID,
		search *string,
		minPrice *int64,
		maxPrice *int64,
		rating *float64,
		variantIDs *[]uuid.UUID,
		categoryIDs *[]uuid.UUID,
		deleted DeletedParam,
		sortRating *string,
		sortPrice *string,
		limit int,
		offset int,
	) (*[]Product, error)

	Count(
		ctx context.Context,
		ids *[]uuid.UUID,
		minPrice *int64,
		maxPrice *int64,
		rating *float64,
		categoryIDs *[]uuid.UUID,
		deleted DeletedParam,
	) (*int, error)

	Get(
		ctx context.Context,
		productID uuid.UUID,
	) (*Product, error)

	Save(
		ctx context.Context,
		product Product,
	) error
}
