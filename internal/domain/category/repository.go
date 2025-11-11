package category

import (
	"context"
)

type Repository interface {
	ListCategories(ctx context.Context, queryParams *QueryParams) (*PaginationModel, error)
}
