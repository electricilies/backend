package category

import (
	"context"

	"backend/internal/domain/category"
)

type repositoryImpl struct{}

func NewRepository() category.Repository {
	return &repositoryImpl{}
}

func (r *repositoryImpl) List(ctx context.Context, queryParams *category.QueryParams) (*category.PaginationModel, error) {
	return &category.PaginationModel{}, nil
}

func (r *repositoryImpl) Create(ctx context.Context, categoryModel *category.Model) (*category.Model, error) {
	return &category.Model{}, nil
}

func (r *repositoryImpl) Update(ctx context.Context, categoryModel *category.Model, id int) (*category.Model, error) {
	return &category.Model{}, nil
}
