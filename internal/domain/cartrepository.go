package domain

import (
	"context"
)

type CartRepository interface {
	Get(
		ctx context.Context,
		id int,
	) (*Cart, error)

	Save(
		ctx context.Context,
		cart *Cart,
	) (*Cart, error)

	Remove(
		ctx context.Context,
		id int,
	) error
}
