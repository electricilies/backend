package application

import (
	"context"

	"backend/internal/domain"
)

type CategoryImpl struct {
	categoryRepo    domain.CategoryRepository
	categoryService domain.CategoryService
	categoryCache   CategoryCache
}

func ProvideCategory(categoryRepo domain.CategoryRepository, categoryService domain.CategoryService, categoryCache CategoryCache) *CategoryImpl {
	return &CategoryImpl{
		categoryRepo:    categoryRepo,
		categoryService: categoryService,
		categoryCache:   categoryCache,
	}
}

var _ Category = &CategoryImpl{}

func (c *CategoryImpl) Create(ctx context.Context, param CreateCategoryParam) (*domain.Category, error) {
	category, err := c.categoryService.Create(param.Data.Name)
	if err != nil {
		return nil, err
	}

	err = c.categoryRepo.Save(ctx, *category)
	if err != nil {
		return nil, err
	}

	_ = c.categoryCache.InvalidateCategoryList(ctx)

	return category, nil
}

func (c *CategoryImpl) List(ctx context.Context, param ListCategoryParam) (*Pagination[domain.Category], error) {
	cacheKey := c.categoryCache.BuildListCacheKey(
		param.Search,
		param.Limit,
		param.Page,
	)

	// Try to get from cache
	if cachedPagination, err := c.categoryCache.GetCategoryList(ctx, cacheKey); err == nil {
		return cachedPagination, nil
	}

	categories, err := c.categoryRepo.List(
		ctx,
		param.Search,
		param.Limit,
		param.Page,
	)
	if err != nil {
		return nil, err
	}

	// CategoryRepository.Count doesn't take search parameter
	count, err := c.categoryRepo.Count(ctx)
	if err != nil {
		return nil, err
	}

	pagination := newPagination(
		*categories,
		*count,
		param.Page,
		param.Limit,
	)

	_ = c.categoryCache.SetCategoryList(ctx, cacheKey, pagination)

	return pagination, nil
}

func (c *CategoryImpl) Get(ctx context.Context, param GetCategoryParam) (*domain.Category, error) {
	if cachedCategory, err := c.categoryCache.GetCategory(ctx, param.CategoryID); err == nil {
		return cachedCategory, nil
	}

	category, err := c.categoryRepo.Get(ctx, param.CategoryID)
	if err != nil {
		return nil, err
	}

	_ = c.categoryCache.SetCategory(ctx, param.CategoryID, category)

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

	err = c.categoryRepo.Save(ctx, *category)
	if err != nil {
		return nil, err
	}

	_ = c.categoryCache.InvalidateCategory(ctx, param.CategoryID)
	_ = c.categoryCache.InvalidateCategoryList(ctx)

	return category, nil
}
