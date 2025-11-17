package service

import (
	"context"

	"backend/internal/domain"
)

type ListCategoryParam struct {
	PaginationParam
}

type GetCategoryParam struct {
	CategoryID int `json:"categoryId" binding:"required"`
}

type CreateCategoryParam struct {
	Name string `json:"name" binding:"required"`
}

type UpdateCategoryParam struct {
	Name string `json:"name" binding:"required"`
}

type Category interface {
	List(context.Context, ListCategoryParam) (*domain.Pagination[domain.Category], error)
	Create(context.Context, CreateCategoryParam) (*domain.Category, error)
	Update(context.Context, UpdateCategoryParam) (*domain.Category, error)
	Get(context.Context, GetCategoryParam) (*domain.Category, error)
}

type CategoryImpl struct{}

func ProvideCategory() *CategoryImpl {
	return &CategoryImpl{}
}

var _ Category = &CategoryImpl{}

func (s *CategoryImpl) List(ctx context.Context, param ListCategoryParam) (*domain.Pagination[domain.Category], error) {
	return nil, nil
}

func (s *CategoryImpl) Create(ctx context.Context, param CreateCategoryParam) (*domain.Category, error) {
	return nil, nil
}

func (s *CategoryImpl) Update(ctx context.Context, param UpdateCategoryParam) (*domain.Category, error) {
	return nil, nil
}

func (s *CategoryImpl) Get(ctx context.Context, param GetCategoryParam) (*domain.Category, error) {
	return nil, nil
}
