package review

import "context"

type Repository interface {
	ListByProduct(context.Context, int, *QueryParams) (*Pagination, error)
	Get(context.Context, int) (*Model, error)
	Create(context.Context, *Model) (*Model, error)
	Update(context.Context, *Model, int) (*Model, error)
	Delete(context.Context, int) error
}
