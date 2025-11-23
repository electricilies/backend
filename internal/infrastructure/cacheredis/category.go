package cacheredis

import (
	"context"
	"encoding/json"
	"time"

	"backend/internal/application"
	"backend/internal/domain"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

// CategoryCache implements application.CategoryCache interface using Redis
type CategoryCache struct {
	redisClient *redis.Client
}

// ProvideCategoryCache creates a new CategoryCache instance
func ProvideCategoryCache(redisClient *redis.Client) *CategoryCache {
	return &CategoryCache{
		redisClient: redisClient,
	}
}

var _ application.CategoryCache = (*CategoryCache)(nil)

// GetCategory retrieves a cached category by ID
func (c *CategoryCache) GetCategory(ctx context.Context, categoryID uuid.UUID) (*domain.Category, error) {
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

	var category domain.Category
	if err := json.Unmarshal([]byte(cachedData), &category); err != nil {
		return nil, err
	}

	return &category, nil
}

// SetCategory caches a category with the specified TTL in seconds
func (c *CategoryCache) SetCategory(ctx context.Context, categoryID uuid.UUID, category *domain.Category) error {
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
func (c *CategoryCache) GetCategoryList(ctx context.Context, cacheKey string) (*application.Pagination[domain.Category], error) {
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

	var pagination application.Pagination[domain.Category]
	if err := json.Unmarshal([]byte(cachedData), &pagination); err != nil {
		return nil, err
	}

	return &pagination, nil
}

// SetCategoryList caches a category list pagination result with the specified TTL in seconds
func (c *CategoryCache) SetCategoryList(ctx context.Context, cacheKey string, pagination *application.Pagination[domain.Category]) error {
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
func (c *CategoryCache) InvalidateCategory(ctx context.Context, categoryID uuid.UUID) error {
	if c.redisClient == nil {
		return nil
	}

	cacheKey := CategoryGetKey(categoryID)
	return c.redisClient.Del(ctx, cacheKey).Err()
}

// InvalidateCategoryList removes all cached category list entries
func (c *CategoryCache) InvalidateCategoryList(ctx context.Context) error {
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
func (c *CategoryCache) BuildListCacheKey(
	search *string,
	limit, page int,
) string {
	return CategoryListKey(
		search,
		limit,
		page,
	)
}
