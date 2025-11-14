package attribute

import "context"

type Repository interface {
	List(ctx context.Context, queryParams *QueryParams) (*PaginationModel, error)
	Create(ctx context.Context, model *Model) (*Model, error)
	Update(ctx context.Context, model *Model, id int) (*Model, error)
	Delete(ctx context.Context, id int) error
	UpdateValues(ctx context.Context, id int, values *[]ValueModel) (*[]ValueModel, error)
}
