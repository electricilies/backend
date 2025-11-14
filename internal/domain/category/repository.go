package category

import (
	"context"
)

type Repository interface {
	List(ctx context.Context, queryParams *QueryParams) (*PaginationModel, error)
	Create(ctx context.Context, category *Model) (*Model, error)
	Update(ctx context.Context, category *Model, id int) (*Model, error)
}
