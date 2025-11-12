package attribute

import (
	"context"

	"backend/internal/domain/attribute"
)

type Repository struct{}

func NewRepository() attribute.Repository {
	return Repository{}
}

func (r Repository) List(ctx context.Context, queryParams *attribute.QueryParams) (*attribute.PaginationModel, error) {
	return &attribute.PaginationModel{}, nil
}
