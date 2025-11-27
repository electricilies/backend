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

func (c *Category) Create(ctx context.Context, param http.CreateCategoryRequestDto) (*http.CategoryResponseDto, error) {
	category, err := domain.NewCategory(param.Data.Name)
	if err != nil {
		return nil, err
	}
	if err := c.categoryService.Validate(*category); err != nil {
		return nil, err
	}

	err = c.categoryRepo.Save(ctx, domain.CategoryRepositorySaveParam{Category: *category})
	if err != nil {
		return nil, err
	}

	_ = c.categoryCache.InvalidateCategoryList(ctx)

	return http.ToCategoryResponseDto(category), nil
}

func (c *Category) List(ctx context.Context, param http.ListCategoryRequestDto) (*http.PaginationResponseDto[http.CategoryResponseDto], error) {
	cacheKey := c.categoryCache.BuildListCacheKey(
		&param.Search,
		param.Limit,
		param.Page,
	)

	// Try to get from cache
	if cachedPagination, err := c.categoryCache.GetCategoryList(ctx, cacheKey); err == nil {
		return cachedPagination, nil
	}

	categories, err := c.categoryRepo.List(
		ctx,
		domain.CategoryRepositoryListParam{
			Search: param.Search,
			Limit:  param.Limit,
			Offset: (param.Page - 1) * param.Limit,
		},
	)
	if err != nil {
		return nil, err
	}

	// CategoryRepository.Count doesn't take search parameter
	count, err := c.categoryRepo.Count(ctx)
	if err != nil {
		return nil, err
	}

	// Map domain models to response DTOs
	categoryDtos := http.ToCategoryResponseDtoList(*categories)

	pagination := newPaginationResponseDto(
		categoryDtos,
		*count,
		param.Page,
		param.Limit,
	)

	_ = c.categoryCache.SetCategoryList(ctx, cacheKey, pagination)

	return pagination, nil
}

func (c *Category) Get(ctx context.Context, param http.GetCategoryRequestDto) (*http.CategoryResponseDto, error) {
	if cachedCategory, err := c.categoryCache.GetCategory(ctx, param.CategoryID); err == nil {
		return cachedCategory, nil
	}

	category, err := c.categoryRepo.Get(ctx, domain.CategoryRepositoryGetParam{ID: param.CategoryID})
	if err != nil {
		return nil, err
	}

	categoryDto := http.ToCategoryResponseDto(category)
	_ = c.categoryCache.SetCategory(ctx, param.CategoryID, categoryDto)

	return categoryDto, nil
}

func (c *Category) Update(ctx context.Context, param http.UpdateCategoryRequestDto) (*http.CategoryResponseDto, error) {
	category, err := c.categoryRepo.Get(ctx, domain.CategoryRepositoryGetParam{ID: param.CategoryID})
	if err != nil {
		return nil, err
	}

	category.Update(param.Data.Name)
	if err := c.categoryService.Validate(*category); err != nil {
		return nil, err
	}

	err = c.categoryRepo.Save(ctx, domain.CategoryRepositorySaveParam{Category: *category})
	if err != nil {
		return nil, err
	}

	_ = c.categoryCache.InvalidateCategory(ctx, param.CategoryID)
	_ = c.categoryCache.InvalidateCategoryList(ctx)

	return http.ToCategoryResponseDto(category), nil
}
