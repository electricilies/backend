package service

import (
	"context"

	"backend/internal/domain"
)

type ListCategoryParam struct {
	Limit  int
	Offset int
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
	ListCategories(context.Context, ListCategoryParam) (*domain.DataPagination, error)
	CreateCategory(context.Context, CreateCategoryParam) (*domain.Category, error)
	UpdateCategory(context.Context, UpdateCategoryParam) (*domain.Category, error)
	GetCategory(context.Context, GetCategoryParam) (*domain.Category, error)
}

type CategoryImpl struct{}

func ProvideCategory() *CategoryImpl {
	return &CategoryImpl{}
}

var _ Category = &CategoryImpl{}

func (s *CategoryImpl) ListCategories(ctx context.Context, param ListCategoryParam) (*domain.DataPagination, error) {
	return nil, nil
}

func (s *CategoryImpl) CreateCategory(ctx context.Context, param CreateCategoryParam) (*domain.Category, error) {
	return nil, nil
}

func (s *CategoryImpl) UpdateCategory(ctx context.Context, param UpdateCategoryParam) (*domain.Category, error) {
	return nil, nil
}

func (s *CategoryImpl) GetCategory(ctx context.Context, param GetCategoryParam) (*domain.Category, error) {
	return nil, nil
}

