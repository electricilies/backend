package review

import "context"

type Repository interface {
	ListByProduct(ctx context.Context, productID int, queryParams *QueryParams) (*Pagination, error)
	Get(ctx context.Context, id int) (*Model, error)
	Create(ctx context.Context, review *Model) (*Model, error)
	Update(ctx context.Context, review *Model, id int) (*Model, error)
	Delete(ctx context.Context, id int) error
}
