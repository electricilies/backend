package service

import (
	"context"

	"backend/internal/domain"
)

type CategoryImpl struct{}

func ProvideCategory() *CategoryImpl {
	return &CategoryImpl{}
}

var _ Category = &CategoryImpl{}

func (s *CategoryImpl) List(ctx context.Context, param ListCategoryParam) (*Pagination[domain.Category], error) {
	panic("implement me")
}

func (s *CategoryImpl) Create(ctx context.Context, param CreateCategoryParam) (*domain.Category, error) {
	panic("implement me")
}

func (s *CategoryImpl) Update(ctx context.Context, param UpdateCategoryParam) (*domain.Category, error) {
	panic("implement me")
}

func (s *CategoryImpl) Get(ctx context.Context, param GetCategoryParam) (*domain.Category, error) {
	panic("implement me")
}
