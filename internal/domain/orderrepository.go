package domain

import (
	"context"

	"github.com/google/uuid"
)

type OrderRepository interface {
	List(
		ctx context.Context,
		params OrderRepositoryListParam,
	) (*[]Order, error)

	Count(
		ctx context.Context,
		params OrderRepositoryCountParam,
	) (*int, error)

	Get(
		ctx context.Context,
		params OrderRepositoryGetParam,
	) (*Order, error)

	Save(
		ctx context.Context,
		params OrderRepositorySaveParam,
	) error
}

type OrderRepositoryListParam struct {
	IDs     []uuid.UUID
	Search  string
	Status  OrderStatus
	Deleted DeletedParam
	Limit   int
	Offset  int
}

type OrderRepositoryCountParam struct {
	IDs     []uuid.UUID
	Deleted DeletedParam
}

type OrderRepositoryGetParam struct {
	ID uuid.UUID
}

type OrderRepositorySaveParam struct {
	Order Order
}
