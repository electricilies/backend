package application

import (
	"context"

	"backend/internal/domain"
)

type Category interface {
	Create(ctx context.Context, param CreateCategoryParam) (*domain.Category, error)
	List(ctx context.Context, param ListCategoryParam) (*Pagination[domain.Category], error)
	Get(ctx context.Context, param GetCategoryParam) (*Category, error)
	Update(ctx context.Context, param UpdateCategoryParam) (*domain.Category, error)
}
