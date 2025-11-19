package application

import (
	"context"

	"backend/internal/domain"
)

type Category interface {
	List(ctx context.Context, param ListCategoryParam) (*Pagination[domain.Category], error)
}
