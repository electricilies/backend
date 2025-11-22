package domain

import (
	"context"

	"github.com/google/uuid"
)

type CartRepository interface {
	Get(
		ctx context.Context,
		id uuid.UUID,
	) (*Cart, error)

	Save(
		ctx context.Context,
		cart Cart,
	) (*Cart, error)
}
