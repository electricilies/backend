package domain

import (
	"context"

	"github.com/google/uuid"
)

type OrderRepository interface {
	List(
		ctx context.Context,
		ids *[]uuid.UUID,
		search *string,
		deleted DeletedParam,
		limit int,
		offset int,
	) (*[]Order, error)

	Count(
		ctx context.Context,
		ids *[]uuid.UUID,
		deleted DeletedParam,
	) (*int, error)

	Get(
		ctx context.Context,
		id uuid.UUID,
	) (*Order, error)

	Save(
		ctx context.Context,
		order Order,
	) error
}
