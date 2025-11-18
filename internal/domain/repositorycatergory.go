package domain

import (
	"context"

	"backend/internal/service"
)

type CategoryRepository interface {
	List(context.Context, service.ListCategoryParam) (Pagination[Category], error)
	Get(context.Context, service.GetCategoryParam) (Category, error)
	Create(context.Context, service.CreateCategoryParam) (Category, error)
	Update(context.Context, service.UpdateCategoryParam) (Category, error)
}
