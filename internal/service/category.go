package service

import (
	"context"

	"backend/internal/domain"
)

type Category interface {
	List(ctx context.Context, param ListCategoryParam) (*Pagination[domain.Category], error)
	Create(ctx context.Context, param CreateCategoryParam) (*domain.Category, error)
	Update(ctx context.Context, param UpdateCategoryParam) (*domain.Category, error)
	Get(ctx context.Context, param GetCategoryParam) (*domain.Category, error)
}
