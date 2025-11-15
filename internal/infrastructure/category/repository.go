package category

import (
	"context"

	"backend/internal/domain/category"
)

type RepositoryImpl struct{}

func NewRepository() category.Repository {
	return &RepositoryImpl{}
}

func ProvideRepository() *RepositoryImpl {
	return &RepositoryImpl{}
}

func (r *RepositoryImpl) List(ctx context.Context, queryParams *category.QueryParams) (*category.PaginationModel, error) {
	return &category.PaginationModel{}, nil
}

func (r *RepositoryImpl) Create(ctx context.Context, categoryModel *category.Model) (*category.Model, error) {
	return &category.Model{}, nil
}

func (r *RepositoryImpl) Update(ctx context.Context, categoryModel *category.Model, id int) (*category.Model, error) {
	return &category.Model{}, nil
}
