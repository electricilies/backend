package attribute

import (
	"context"

	"backend/internal/domain/attribute"
)

type RepositoryImpl struct{}

func NewRepository() attribute.Repository {
	return &RepositoryImpl{}
}

func ProvideRepository() *RepositoryImpl {
	return &RepositoryImpl{}
}

func (r *RepositoryImpl) List(
	ctx context.Context,
	queryParams *attribute.QueryParams,
) (*attribute.PaginationModel, error) {
	return &attribute.PaginationModel{}, nil
}

func (r *RepositoryImpl) Create(ctx context.Context, model *attribute.Model) (*attribute.Model, error) {
	return model, nil
}

func (r *RepositoryImpl) Update(ctx context.Context, model *attribute.Model, id int) (*attribute.Model, error) {
	return model, nil
}

func (r *RepositoryImpl) Delete(ctx context.Context, id int) error {
	return nil
}

func (r *RepositoryImpl) UpdateValues(
	ctx context.Context,
	id int,
	values *[]attribute.ValueModel,
) (*[]attribute.ValueModel, error) {
	return values, nil
}
