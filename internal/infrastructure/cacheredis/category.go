package cacheredis

import (
	"context"
	"encoding/json"
	"time"

	"backend/internal/application"
	"backend/internal/delivery/http"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

// Category implements application.Category interface using Redis
type Category struct {
	redisClient *redis.Client
}

// ProvideCategory creates a new CategoryCache instance
func ProvideCategory(redisClient *redis.Client) *Category {
	return &Category{
		redisClient: redisClient,
	}
}

var _ application.CategoryCache = (*Category)(nil)

// GetCategory retrieves a cached category by ID
func (c *Category) GetCategory(ctx context.Context, categoryID uuid.UUID) (*http.CategoryResponseDto, error) {
	if c.redisClient == nil {
		return nil, redis.Nil
	}

	cacheKey := CategoryGetKey(categoryID)
	cachedData, err := c.redisClient.Get(ctx, cacheKey).Result()
	if err != nil {
		return nil, err
	}

	if cachedData == "" {
		return nil, redis.Nil
	}

	var category http.CategoryResponseDto
	if err := json.Unmarshal([]byte(cachedData), &category); err != nil {
		return nil, err
	}

	return &category, nil
}

// SetCategory caches a category with the specified TTL in seconds
func (c *Category) SetCategory(ctx context.Context, categoryID uuid.UUID, category *http.CategoryResponseDto) error {
	if c.redisClient == nil {
		return nil
	}

	cacheKey := CategoryGetKey(categoryID)
	data, err := json.Marshal(category)
	if err != nil {
		return err
	}

	return c.redisClient.Set(ctx, cacheKey, data, time.Duration(CacheTTLCategory)*time.Second).Err()
}

// GetCategoryList retrieves a cached category list pagination result
func (c *Category) GetCategoryList(ctx context.Context, cacheKey string) (*http.PaginationResponseDto[http.CategoryResponseDto], error) {
	if c.redisClient == nil {
		return nil, redis.Nil
	}

	cachedData, err := c.redisClient.Get(ctx, cacheKey).Result()
	if err != nil {
		return nil, err
	}

	if cachedData == "" {
		return nil, redis.Nil
	}

	var pagination http.PaginationResponseDto[http.CategoryResponseDto]
	if err := json.Unmarshal([]byte(cachedData), &pagination); err != nil {
		return nil, err
	}

	return &pagination, nil
}

// SetCategoryList caches a category list pagination result with the specified TTL in seconds
func (c *Category) SetCategoryList(ctx context.Context, cacheKey string, pagination *http.PaginationResponseDto[http.CategoryResponseDto]) error {
	if c.redisClient == nil {
		return nil
	}

	data, err := json.Marshal(pagination)
	if err != nil {
		return err
	}

	return c.redisClient.Set(ctx, cacheKey, data, time.Duration(CacheTTLCategory)*time.Second).Err()
}

// InvalidateCategory removes the cached category by ID
func (c *Category) InvalidateCategory(ctx context.Context, categoryID uuid.UUID) error {
	if c.redisClient == nil {
		return nil
	}

	cacheKey := CategoryGetKey(categoryID)
	return c.redisClient.Del(ctx, cacheKey).Err()
}

// InvalidateCategoryList removes all cached category list entries
func (c *Category) InvalidateCategoryList(ctx context.Context) error {
	if c.redisClient == nil {
		return nil
	}

	iter := c.redisClient.Scan(ctx, 0, CategoryListPrefix+"*", 0).Iterator()
	for iter.Next(ctx) {
		if err := c.redisClient.Del(ctx, iter.Val()).Err(); err != nil {
			return err
		}
	}
	return iter.Err()
}

// BuildListCacheKey builds a cache key for category list queries
func (c *Category) BuildListCacheKey(
	search *string,
	limit, page int,
) string {
	return CategoryListKey(
		search,
		limit,
		page,
	)
}
