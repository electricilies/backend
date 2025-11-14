package review

import "context"

type Repository interface {
	ListReviews(ctx context.Context, productID int, queryParams *QueryParams) (*Pagination, error)
}
