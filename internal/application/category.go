package application

import (
	"context"

	"backend/internal/domain/category"
)

type Category interface {
	ListCategories(context.Context, *category.QueryParams) (*category.PaginationModel, error)
}

type CategoryImpl struct {
	categoryRepo category.Repository
}

func NewCategory(categoryRepo category.Repository) Category {
	return &CategoryImpl{
		categoryRepo: categoryRepo,
	}
}

func ProvideCategory(
	categoryRepo category.Repository,
) *CategoryImpl {
	return &CategoryImpl{
		categoryRepo: categoryRepo,
	}
}

func (a *CategoryImpl) ListCategories(
	ctx context.Context,
	queryParams *category.QueryParams,
) (*category.PaginationModel, error) {
	return a.categoryRepo.List(ctx, queryParams)
}
