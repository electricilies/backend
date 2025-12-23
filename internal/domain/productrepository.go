package domain

import (
	"context"

	"github.com/google/uuid"
)

type ProductRepository interface {
	List(
		ctx context.Context,
		params ProductRepositoryListParam,
	) (*[]Product, error)

	Count(
		ctx context.Context,
		params ProductRepositoryCountParam,
	) (*int, error)

	Get(
		ctx context.Context,
		params ProductRepositoryGetParam,
	) (*Product, error)

	Save(
		ctx context.Context,
		params ProductRepositorySaveParam,
	) error
}

type ProductRepositoryListParam struct {
	IDs         []uuid.UUID
	Search      string
	MinPrice    int64
	MaxPrice    int64
	Rating      float64
	VariantIDs  []uuid.UUID
	CategoryIDs []uuid.UUID
	Deleted     DeletedParam
	SortRating  string
	SortPrice   string
	Limit       int
	Offset      int
}

type ProductRepositoryCountParam struct {
	IDs         []uuid.UUID
	Search      string
	MinPrice    int64
	MaxPrice    int64
	Rating      float64
	VariantIDs  []uuid.UUID
	CategoryIDs []uuid.UUID
	Deleted     DeletedParam
}

type ProductRepositoryGetParam struct {
	ProductID uuid.UUID
}

type ProductRepositorySaveParam struct {
	Product Product
}
