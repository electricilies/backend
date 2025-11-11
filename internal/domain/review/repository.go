package review

import "context"

type Repository interface {
	ListReviewsByProductID(ctx context.Context, productID int, queryParams *QueryParams) (*Pagination, error)
}
