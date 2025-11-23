package application

import (
	"context"
	"encoding/json"
	"time"

	"backend/internal/constant"
	"backend/internal/domain"

	"github.com/redis/go-redis/v9"
)

type CategoryImpl struct {
	categoryRepo    domain.CategoryRepository
	categoryService domain.CategoryService
	redisClient     *redis.Client
}

func ProvideCategory(categoryRepo domain.CategoryRepository, categoryService domain.CategoryService, redisClient *redis.Client) *CategoryImpl {
	return &CategoryImpl{
		categoryRepo:    categoryRepo,
		categoryService: categoryService,
		redisClient:     redisClient,
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

	// Invalidate list cache
	if c.redisClient != nil {
		iter := c.redisClient.Scan(ctx, 0, constant.CategoryListPrefix+"*", 0).Iterator()
		for iter.Next(ctx) {
			c.redisClient.Del(ctx, iter.Val())
		}
	}

	return category, nil
}

func (c *CategoryImpl) List(ctx context.Context, param ListCategoryParam) (*Pagination[domain.Category], error) {
	cacheKey := constant.CategoryListKey(
		param.Search,
		param.Limit,
		param.Page,
	)

	// Try to get from cache
	if c.redisClient != nil {
		cachedData, err := c.redisClient.Get(ctx, cacheKey).Result()
		if err == nil && cachedData != "" {
			var pagination Pagination[domain.Category]
			if err := json.Unmarshal([]byte(cachedData), &pagination); err == nil {
				return &pagination, nil
			}
		}
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

	// Cache the result
	if c.redisClient != nil {
		if data, err := json.Marshal(pagination); err == nil {
			c.redisClient.Set(ctx, cacheKey, data, time.Duration(constant.CacheTTLCategory)*time.Second)
		}
	}

	return pagination, nil
}

func (c *CategoryImpl) Get(ctx context.Context, param GetCategoryParam) (*domain.Category, error) {
	// Build cache key
	cacheKey := constant.CategoryGetKey(param.CategoryID)

	// Try to get from cache
	if c.redisClient != nil {
		cachedData, err := c.redisClient.Get(ctx, cacheKey).Result()
		if err == nil && cachedData != "" {
			var category domain.Category
			if err := json.Unmarshal([]byte(cachedData), &category); err == nil {
				return &category, nil
			}
		}
	}

	category, err := c.categoryRepo.Get(ctx, param.CategoryID)
	if err != nil {
		return nil, err
	}

	// Cache the result
	if c.redisClient != nil {
		if data, err := json.Marshal(category); err == nil {
			c.redisClient.Set(ctx, cacheKey, data, time.Duration(constant.CacheTTLCategory)*time.Second)
		}
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

	err = c.categoryRepo.Save(ctx, *category)
	if err != nil {
		return nil, err
	}

	// Invalidate cache
	if c.redisClient != nil {
		// Delete specific category cache
		c.redisClient.Del(ctx, constant.CategoryGetKey(param.CategoryID))
		// Delete list cache (use pattern matching)
		iter := c.redisClient.Scan(ctx, 0, constant.CategoryListPrefix+"*", 0).Iterator()
		for iter.Next(ctx) {
			c.redisClient.Del(ctx, iter.Val())
		}
	}

	return category, nil
}
