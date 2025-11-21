package domain

import (
	"context"
)

type ProductRepository interface {
	List(
		ctx context.Context,
		ids *[]int,
		search *string,
		min_price *int64,
		max_price *int64,
		rating *float64,
		category_ids *[]int,
		deleted string,
		sort_rating *string,
		sort_price *string,
		limit int,
		offset int,
	) (*[]Product, error)

	Count(
		ctx context.Context,
		ids *[]int,
		min_price *int64,
		max_price *int64,
		rating *float64,
		category_ids *[]int,
		deleted string,
		sort_rating *string,
		sort_price *string,
		limit int,
		offset int,
	) (*int, error)

	Get(
		ctx context.Context,
		productID int,
	) (*Product, error)

	Save(
		ctx context.Context,
		product *Product,
	) (*Product, error)

	Remove(
		ctx context.Context,
		productID int,
	) error
}
