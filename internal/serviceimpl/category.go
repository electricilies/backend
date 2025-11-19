package serviceimpl

import (
	"context"

	"backend/internal/domain"
)

type Category struct{}

func ProvideCategory() *Category {
	return &Category{}
}

var _ domain.CategoryService = &Category{}

func (s *Category) Create(ctx context.Context, param domain.CreateCategoryParam) (*domain.Category, error) {
	panic("implement me")
}

func (s *Category) Update(ctx context.Context, param domain.UpdateCategoryParam) (*domain.Category, error) {
	panic("implement me")
}

func (s *Category) Get(ctx context.Context, param domain.GetCategoryParam) (*domain.Category, error) {
	panic("implement me")
}
