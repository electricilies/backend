package domain

import (
	"context"

	"github.com/google/uuid"
)

type CategoryRepository interface {
	List(
		ctx context.Context,
		search *string,
		limit int,
		offset int,
	) (*[]Category, error)

	Count(
		ctx context.Context,
	) (*int, error)

	Get(
		ctx context.Context,
		id uuid.UUID,
	) (*Category, error)

	Save(
		ctx context.Context,
		category *Category,
	) (*Category, error)

	Remove(
		ctx context.Context,
		id uuid.UUID,
	) error
}
