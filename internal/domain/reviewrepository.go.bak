package domain

import (
	"context"

	"github.com/google/uuid"
)

type ReviewRepository interface {
	List(
		ctx context.Context,
		params ReviewRepositoryListParam,
	) (*[]Review, error)

	Count(
		ctx context.Context,
		params ReviewRepositoryCountParam,
	) (*int, error)

	Get(
		ctx context.Context,
		params ReviewRepositoryGetParam,
	) (*Review, error)

	Save(
		ctx context.Context,
		params ReviewRepositorySaveParam,
	) error
}

type ReviewRepositoryListParam struct {
	OrderItemIDs     []uuid.UUID
	ProductVariantID uuid.UUID
	UserIDs          []uuid.UUID
	Deleted          DeletedParam
	Limit            int
	Offset           int
}

type ReviewRepositoryCountParam struct {
	OrderItemIDs     []uuid.UUID
	ProductVariantID uuid.UUID
	UserIDs          []uuid.UUID
	Deleted          DeletedParam
}

type ReviewRepositoryGetParam struct {
	ID uuid.UUID
}

type ReviewRepositorySaveParam struct {
	Review Review
}
