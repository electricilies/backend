package domain

import (
	"context"

	"github.com/google/uuid"
)

type CategoryRepository interface {
	List(
		ctx context.Context,
		params CategoryRepositoryListParam,
	) (*[]Category, error)

	Count(
		ctx context.Context,
		params CategoryRepositoryCountParam,
	) (*int, error)

	Get(
		ctx context.Context,
		params CategoryRepositoryGetParam,
	) (*Category, error)

	Save(
		ctx context.Context,
		params CategoryRepositorySaveParam,
	) error
}

type CategoryRepositoryListParam struct {
	IDs     []uuid.UUID
	Search  string
	Deleted DeletedParam
	Limit   int
	Offset  int
}

type CategoryRepositoryCountParam struct {
	IDs     []uuid.UUID
	Search  string
	Deleted DeletedParam
}

type CategoryRepositoryGetParam struct {
	ID uuid.UUID
}

type CategoryRepositorySaveParam struct {
	Category Category
}
