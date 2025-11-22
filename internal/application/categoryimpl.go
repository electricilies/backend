package application

import (
	"context"

	"backend/internal/domain"
)

type CategoryImpl struct {
	categoryRepo    domain.CategoryRepository
	categoryService domain.CategoryService
}

func ProvideCategory(categoryRepo domain.CategoryRepository, categoryService domain.CategoryService) *CategoryImpl {
	return &CategoryImpl{
		categoryRepo:    categoryRepo,
		categoryService: categoryService,
	}
}

var _ Category = &CategoryImpl{}

func (c *CategoryImpl) Create(ctx context.Context, param CreateCategoryParam) (*domain.Category, error) {
	category, err := c.categoryService.Create(param.Data.Name)
	if err != nil {
		return nil, err
	}

	savedCategory, err := c.categoryRepo.Save(ctx, *category)
	if err != nil {
		return nil, err
	}

	return savedCategory, nil
}

func (c *CategoryImpl) List(ctx context.Context, param ListCategoryParam) (*Pagination[domain.Category], error) {
	categories, err := c.categoryRepo.List(
		ctx,
		param.Search,
		*param.Limit,
		*param.Page,
	)
	if err != nil {
		return nil, err
	}

	// CategoryRepository.Count doesn't take search parameter
	count, err := c.categoryRepo.Count(ctx)
	if err != nil {
		return nil, err
	}

	pagination := newPagination(*categories, *count, *param.Page, *param.Limit)
	return pagination, nil
}

func (c *CategoryImpl) Get(ctx context.Context, param GetCategoryParam) (*domain.Category, error) {
	category, err := c.categoryRepo.Get(ctx, param.CategoryID)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c *CategoryImpl) Update(ctx context.Context, param UpdateCategoryParam) (*domain.Category, error) {
	category, err := c.categoryRepo.Get(ctx, param.CategoryID)
	if err != nil {
		return nil, err
	}

	err = c.categoryService.Update(category, param.Data.Name)
	if err != nil {
		return nil, err
	}

	savedCategory, err := c.categoryRepo.Save(ctx, *category)
	if err != nil {
		return nil, err
	}

	return savedCategory, nil
}
