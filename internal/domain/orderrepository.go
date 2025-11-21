package domain

import (
	"context"
)

type OrderRepository interface {
	Count(
		ctx context.Context,
		ids *[]int,
		deleted string,
		limit int,
		offset int,
	) (*int, error)

	Get(
		ctx context.Context,
		id int,
	) (*Order, error)

	List(
		ctx context.Context,
		ids *[]int,
		search *string,
		deleted string,
		limit int,
		offset int,
	) (*[]Order, error)

	Save(
		ctx context.Context,
		order *Order,
	) (*Order, error)

	Remove(
		ctx context.Context,
		id int,
	) error
}
