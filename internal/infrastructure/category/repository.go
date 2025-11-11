package category

import (
	"context"

	"backend/internal/domain/category"
)

type Repository struct{}

func NewRepository() category.Repository {
	return &Repository{}
}

func (r *Repository) ListCategories(ctx context.Context, queryParams *category.QueryParams) (*category.PaginationModel, error) {
	return &category.PaginationModel{}, nil
}
