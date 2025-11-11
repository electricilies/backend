package attribute

import "context"

type Repository interface {
	List(ctx context.Context, QueryParams *QueryParams) (*PaginationModel, error)
}
