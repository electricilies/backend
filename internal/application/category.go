package application

import (
	"context"

	"backend/internal/domain/category"
)

type Category interface {
	ListCategories(ctx context.Context, queryParams *category.QueryParams) (*category.PaginationModel, error)
}

type categoryApp struct {
	categoryRepo category.Repository
}

func NewCategory(categoryRepo category.Repository) Category {
	return &categoryApp{
		categoryRepo: categoryRepo,
	}
}

func (a *categoryApp) ListCategories(ctx context.Context, queryParams *category.QueryParams) (*category.PaginationModel, error) {
	return a.categoryRepo.List(ctx, queryParams)
}
