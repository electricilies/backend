package domain

import (
	"context"
)

type CategoryService interface {
	Create(ctx context.Context, param CreateCategoryParam) (*Category, error)
	Update(ctx context.Context, param UpdateCategoryParam) (*Category, error)
	Get(ctx context.Context, param GetCategoryParam) (*Category, error)
}
