package domain

import (
	"context"
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
		limit int,
		offset int,
	) (*int, error)

	Get(
		ctx context.Context,
		id int,
	) (*Category, error)

	Save(
		ctx context.Context,
		category *Category,
	) (*Category, error)

	Remove(
		ctx context.Context,
		id int,
	) error
}
