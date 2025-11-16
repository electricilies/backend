package category

import (
	"context"
)

type Repository interface {
	List(context.Context, *QueryParams) (*PaginationModel, error)
	Create(context.Context, *Model) (*Model, error)
	Update(context.Context, *Model, int) (*Model, error)
}
