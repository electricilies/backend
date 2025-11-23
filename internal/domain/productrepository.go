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
		min_price *int64,
		max_price *int64,
		rating *float64,
		category_ids *[]uuid.UUID,
		deleted DeletedParam,
		sort_rating *string,
		sort_price *string,
		limit int,
		offset int,
	) (*[]Product, error)

	Count(
		ctx context.Context,
		ids *[]uuid.UUID,
		min_price *int64,
		max_price *int64,
		rating *float64,
		category_ids *[]uuid.UUID,
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
