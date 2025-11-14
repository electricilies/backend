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

func (r *repositoryImpl) Create(ctx context.Context, model *attribute.Model) (*attribute.Model, error)
func (r *repositoryImpl) Update(ctx context.Context, model *attribute.Model, id int) (*attribute.Model, error)
func (r *repositoryImpl) Delete(ctx context.Context, id int) error
func (r *repositoryImpl) UpdateValues(ctx context.Context, id int, values *[]attribute.ValueModel) (*[]attribute.ValueModel, error)
