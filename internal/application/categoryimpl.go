package application

import (
	"context"

	"backend/internal/delivery/http"
	"backend/internal/domain"
)

type Category struct {
	categoryRepo    domain.CategoryRepository
	categoryService domain.CategoryService
	categoryCache   CategoryCache
}

func ProvideCategory(categoryRepo domain.CategoryRepository, categoryService domain.CategoryService, categoryCache CategoryCache) *Category {
	return &Category{
		categoryRepo:    categoryRepo,
		categoryService: categoryService,
		categoryCache:   categoryCache,
	}
}

var _ http.CategoryApplication = &Category{}

func (c *Category) Create(ctx context.Context, param http.CreateCategoryRequestDto) (*domain.Category, error) {
	category, err := domain.NewCategory(param.Data.Name)
	if err != nil {
		return nil, err
	}
	if err := c.categoryService.Validate(*category); err != nil {
		return nil, err
	}

	err = c.categoryRepo.Save(ctx, *category)
	if err != nil {
		return nil, err
	}

	_ = c.categoryCache.InvalidateCategoryList(ctx)

	return category, nil
}

func (c *Category) List(ctx context.Context, param http.ListCategoryRequestDto) (*http.PaginationResponseDto[domain.Category], error) {
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
		(param.Page-1)*param.Limit,
	)
	if err != nil {
		return nil, err
	}

	// CategoryRepository.Count doesn't take search parameter
	count, err := c.categoryRepo.Count(ctx)
	if err != nil {
		return nil, err
	}

	pagination := newPaginationResponseDto(
		*categories,
		*count,
		param.Page,
		param.Limit,
	)

	_ = c.categoryCache.SetCategoryList(ctx, cacheKey, pagination)

	return pagination, nil
}

func (c *Category) Get(ctx context.Context, param http.GetCategoryRequestDto) (*domain.Category, error) {
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

func (c *Category) Update(ctx context.Context, param http.UpdateCategoryRequestDto) (*domain.Category, error) {
	category, err := c.categoryRepo.Get(ctx, param.CategoryID)
	if err != nil {
		return nil, err
	}

	category.Update(param.Data.Name)
	if err := c.categoryService.Validate(*category); err != nil {
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
