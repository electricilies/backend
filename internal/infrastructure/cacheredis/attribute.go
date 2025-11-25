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

// Attribute implements application.Attribute interface using Redis
type Attribute struct {
	redisClient *redis.Client
}

// ProvideAttribute creates a new AttributeCache instance
func ProvideAttribute(redisClient *redis.Client) *Attribute {
	return &Attribute{
		redisClient: redisClient,
	}
}

var _ application.AttributeCache = (*Attribute)(nil)

// GetAttribute retrieves a cached attribute by ID
func (c *Attribute) GetAttribute(ctx context.Context, attributeID uuid.UUID) (*domain.Attribute, error) {
	if c.redisClient == nil {
		return nil, redis.Nil
	}

	cacheKey := AttributeGetKey(attributeID)
	cachedData, err := c.redisClient.Get(ctx, cacheKey).Result()
	if err != nil {
		return nil, err
	}

	if cachedData == "" {
		return nil, redis.Nil
	}

	var attribute domain.Attribute
	if err := json.Unmarshal([]byte(cachedData), &attribute); err != nil {
		return nil, err
	}

	return &attribute, nil
}

// SetAttribute caches an attribute with the specified TTL in seconds
func (c *Attribute) SetAttribute(ctx context.Context, attributeID uuid.UUID, attribute *domain.Attribute) error {
	if c.redisClient == nil {
		return nil
	}

	cacheKey := AttributeGetKey(attributeID)
	data, err := json.Marshal(attribute)
	if err != nil {
		return err
	}

	return c.redisClient.Set(ctx, cacheKey, data, time.Duration(CacheTTLAttribute)*time.Second).Err()
}

// GetAttributeList retrieves a cached attribute list pagination result
func (c *Attribute) GetAttributeList(ctx context.Context, cacheKey string) (*http.PaginationResponseDto[domain.Attribute], error) {
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

	var pagination http.PaginationResponseDto[domain.Attribute]
	if err := json.Unmarshal([]byte(cachedData), &pagination); err != nil {
		return nil, err
	}

	return &pagination, nil
}

// SetAttributeList caches an attribute list pagination result with the specified TTL in seconds
func (c *Attribute) SetAttributeList(ctx context.Context, cacheKey string, pagination *http.PaginationResponseDto[domain.Attribute]) error {
	if c.redisClient == nil {
		return nil
	}

	data, err := json.Marshal(pagination)
	if err != nil {
		return err
	}

	return c.redisClient.Set(ctx, cacheKey, data, time.Duration(CacheTTLAttribute)*time.Second).Err()
}

// GetAttributeValueList retrieves a cached attribute value list pagination result
func (c *Attribute) GetAttributeValueList(ctx context.Context, cacheKey string) (*http.PaginationResponseDto[domain.AttributeValue], error) {
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

	var pagination http.PaginationResponseDto[domain.AttributeValue]
	if err := json.Unmarshal([]byte(cachedData), &pagination); err != nil {
		return nil, err
	}

	return &pagination, nil
}

// SetAttributeValueList caches an attribute value list pagination result with the specified TTL in seconds
func (c *Attribute) SetAttributeValueList(ctx context.Context, cacheKey string, pagination *http.PaginationResponseDto[domain.AttributeValue]) error {
	if c.redisClient == nil {
		return nil
	}

	data, err := json.Marshal(pagination)
	if err != nil {
		return err
	}

	return c.redisClient.Set(ctx, cacheKey, data, time.Duration(CacheTTLAttributeValue)*time.Second).Err()
}

// InvalidateAttribute removes the cached attribute by ID
func (c *Attribute) InvalidateAttribute(ctx context.Context, attributeID uuid.UUID) error {
	if c.redisClient == nil {
		return nil
	}

	cacheKey := AttributeGetKey(attributeID)
	return c.redisClient.Del(ctx, cacheKey).Err()
}

// InvalidateAttributeList removes all cached attribute list entries
func (c *Attribute) InvalidateAttributeList(ctx context.Context) error {
	if c.redisClient == nil {
		return nil
	}

	iter := c.redisClient.Scan(ctx, 0, AttributeListPrefix+"*", 0).Iterator()
	for iter.Next(ctx) {
		if err := c.redisClient.Del(ctx, iter.Val()).Err(); err != nil {
			return err
		}
	}
	return iter.Err()
}

// InvalidateAttributeValueList removes all cached attribute value list entries
func (c *Attribute) InvalidateAttributeValueList(ctx context.Context) error {
	if c.redisClient == nil {
		return nil
	}

	iter := c.redisClient.Scan(ctx, 0, AttributeValueListPrefix+"*", 0).Iterator()
	for iter.Next(ctx) {
		if err := c.redisClient.Del(ctx, iter.Val()).Err(); err != nil {
			return err
		}
	}
	return iter.Err()
}

// InvalidateAllAttributes removes all attribute-related caches (attribute, list, and value list)
func (c *Attribute) InvalidateAllAttributes(ctx context.Context) error {
	if c.redisClient == nil {
		return nil
	}

	// Invalidate attribute get caches
	iter := c.redisClient.Scan(ctx, 0, AttributeGetPrefix+"*", 0).Iterator()
	for iter.Next(ctx) {
		if err := c.redisClient.Del(ctx, iter.Val()).Err(); err != nil {
			return err
		}
	}
	if err := iter.Err(); err != nil {
		return err
	}

	// Invalidate attribute list caches
	if err := c.InvalidateAttributeList(ctx); err != nil {
		return err
	}

	// Invalidate attribute value list caches
	return c.InvalidateAttributeValueList(ctx)
}

// BuildListCacheKey builds a cache key for attribute list queries
func (c *Attribute) BuildListCacheKey(
	ids *[]uuid.UUID,
	search *string,
	deleted domain.DeletedParam,
	limit, page int,
) string {
	return AttributeListKey(
		ids,
		search,
		deleted,
		limit,
		page,
	)
}

// BuildValueListCacheKey builds a cache key for attribute value list queries
func (c *Attribute) BuildValueListCacheKey(
	attributeID uuid.UUID,
	valueIDs *[]uuid.UUID,
	search *string,
	limit, page int,
) string {
	return AttributeValueListKey(
		attributeID,
		valueIDs,
		search,
		limit,
		page,
	)
}
