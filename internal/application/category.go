package application

import (
	"context"

	"backend/internal/domain/category"
)

type Category interface {
	ListCategories(ctx context.Context, queryParams *category.QueryParams) (*category.PaginationModel, error)
}

type CategoryApp struct {
	categoryRepo category.Repository
}

func NewCategory(categoryRepo category.Repository) Category {
	return &CategoryApp{
		categoryRepo: categoryRepo,
	}
}

func ProvideCategory(
	categoryRepo category.Repository,
) *CategoryApp {
	return &CategoryApp{
		categoryRepo: categoryRepo,
	}
}

func (a *CategoryApp) ListCategories(ctx context.Context, queryParams *category.QueryParams) (*category.PaginationModel, error) {
	return a.categoryRepo.List(ctx, queryParams)
}
