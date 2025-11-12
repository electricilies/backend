package attribute

import "context"

type Repository interface {
	List(ctx context.Context, queryParams *QueryParams) (*PaginationModel, error)
}
