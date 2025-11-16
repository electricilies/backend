package category

import (
	"context"
	"backend/internal/domain/common"
)

type Repository interface {
	Create(ctx context.Context, category *Category) (*Category, error)
	List(ctx context.Context, param ListParam) (*common.Pagination[Repository], error)
	Get(ctx context.Context, param GetParam) (*Category, error)
	Update(ctx context.Context, category *Category) (*Category, error)
}

// TODO: add search by name support
type ListParam struct {
	common.PaginationParam
}

type GetParam struct {
	ID int
}
