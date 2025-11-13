package attribute

import (
	"context"

	"backend/internal/domain/attribute"
)

type repositoryImpl struct{}

func NewRepository() attribute.Repository {
	return &repositoryImpl{}
}

func (r *repositoryImpl) List(ctx context.Context, queryParams *attribute.QueryParams) (*attribute.PaginationModel, error) {
	return &attribute.PaginationModel{}, nil
}
