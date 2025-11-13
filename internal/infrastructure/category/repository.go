package category

import (
	"context"

	"backend/internal/domain/category"
)

type repositoryImpl struct{}

func NewRepository() category.Repository {
	return &repositoryImpl{}
}

func (r *repositoryImpl) ListCategories(ctx context.Context, queryParams *category.QueryParams) (*category.PaginationModel, error) {
	return &category.PaginationModel{}, nil
}
