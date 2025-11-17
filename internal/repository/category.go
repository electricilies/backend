package repository

import (
	"context"

	"backend/internal/domain"
	"backend/internal/service"
)

type Category interface {
	List(context.Context, service.ListCategoryParam) (*domain.Pagination[domain.Category], error)
	Get(context.Context, service.GetCategoryParam) (*domain.Category, error)
	Create(context.Context, service.CreateCategoryParam) (*domain.Category, error)
	Update(context.Context, service.UpdateCategoryParam) (*domain.Category, error)
}
