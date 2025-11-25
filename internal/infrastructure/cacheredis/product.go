package cacheredis

import (
	"context"
	"encoding/json"
	"time"

	"backend/internal/application"
	"backend/internal/delivery/http"
	"backend/internal/domain"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

// Product implements application.Product interface using Redis
type Product struct {
	redisClient *redis.Client
}

// ProvideProduct creates a new ProductCache instance
func ProvideProduct(redisClient *redis.Client) *Product {
	return &Product{
		redisClient: redisClient,
	}
}

var _ application.ProductCache = (*Product)(nil)

// GetProduct retrieves a cached product by ID
func (c *Product) GetProduct(ctx context.Context, productID uuid.UUID) (*domain.Product, error) {
	if c.redisClient == nil {
		return nil, redis.Nil
	}

	cacheKey := ProductGetKey(productID)
	cachedData, err := c.redisClient.Get(ctx, cacheKey).Result()
	if err != nil {
		return nil, err
	}

	if cachedData == "" {
		return nil, redis.Nil
	}

	var product domain.Product
	if err := json.Unmarshal([]byte(cachedData), &product); err != nil {
		return nil, err
	}

	return &product, nil
}

// SetProduct caches a product with the specified TTL in seconds
func (c *Product) SetProduct(ctx context.Context, productID uuid.UUID, product *domain.Product) error {
	if c.redisClient == nil {
		return nil
	}

	cacheKey := ProductGetKey(productID)
	data, err := json.Marshal(product)
	if err != nil {
		return err
	}

	return c.redisClient.Set(ctx, cacheKey, data, time.Duration(CacheTTLProduct)*time.Second).Err()
}

// GetProductList retrieves a cached product list pagination result
func (c *Product) GetProductList(ctx context.Context, cacheKey string) (*http.PaginationResponseDto[domain.Product], error) {
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

	var pagination http.PaginationResponseDto[domain.Product]
	if err := json.Unmarshal([]byte(cachedData), &pagination); err != nil {
		return nil, err
	}

	return &pagination, nil
}

// SetProductList caches a product list pagination result with the specified TTL in seconds
func (c *Product) SetProductList(ctx context.Context, cacheKey string, pagination *http.PaginationResponseDto[domain.Product]) error {
	if c.redisClient == nil {
		return nil
	}

	data, err := json.Marshal(pagination)
	if err != nil {
		return err
	}

	return c.redisClient.Set(ctx, cacheKey, data, time.Duration(CacheTTLProduct)*time.Second).Err()
}

// InvalidateProduct removes the cached product by ID
func (c *Product) InvalidateProduct(ctx context.Context, productID uuid.UUID) error {
	if c.redisClient == nil {
		return nil
	}

	cacheKey := ProductGetKey(productID)
	return c.redisClient.Del(ctx, cacheKey).Err()
}

// InvalidateProductList removes all cached product list entries
func (c *Product) InvalidateProductList(ctx context.Context) error {
	if c.redisClient == nil {
		return nil
	}

	iter := c.redisClient.Scan(ctx, 0, ProductListPrefix+"*", 0).Iterator()
	for iter.Next(ctx) {
		if err := c.redisClient.Del(ctx, iter.Val()).Err(); err != nil {
			return err
		}
	}
	return iter.Err()
}

// InvalidateAllProducts removes all product-related caches (both get and list)
func (c *Product) InvalidateAllProducts(ctx context.Context) error {
	if c.redisClient == nil {
		return nil
	}

	// Invalidate all product get caches
	iter := c.redisClient.Scan(ctx, 0, ProductGetPrefix+"*", 0).Iterator()
	for iter.Next(ctx) {
		if err := c.redisClient.Del(ctx, iter.Val()).Err(); err != nil {
			return err
		}
	}
	if err := iter.Err(); err != nil {
		return err
	}

	// Invalidate all product list caches
	return c.InvalidateProductList(ctx)
}

// BuildListCacheKey builds a cache key for product list queries
func (c *Product) BuildListCacheKey(
	ids *[]uuid.UUID,
	search *string,
	minPrice *int64,
	maxPrice *int64,
	rating *float64,
	categoryIDs *[]uuid.UUID,
	deleted domain.DeletedParam,
	sortRating *string,
	sortPrice *string,
	limit, page int,
) string {
	return ProductListKey(
		ids,
		search,
		minPrice,
		maxPrice,
		rating,
		categoryIDs,
		deleted,
		sortRating,
		sortPrice,
		limit,
		page,
	)
}
