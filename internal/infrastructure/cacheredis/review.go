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

// ReviewCache implements application.ReviewCache interface using Redis
type ReviewCache struct {
	redisClient *redis.Client
}

// ProvideReviewCache creates a new ReviewCache instance
func ProvideReviewCache(redisClient *redis.Client) *ReviewCache {
	return &ReviewCache{
		redisClient: redisClient,
	}
}

var _ application.ReviewCache = (*ReviewCache)(nil)

// GetReviewList retrieves a cached review list pagination result
func (c *ReviewCache) GetReviewList(ctx context.Context, cacheKey string) (*application.Pagination[domain.Review], error) {
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

	var pagination application.Pagination[domain.Review]
	if err := json.Unmarshal([]byte(cachedData), &pagination); err != nil {
		return nil, err
	}

	return &pagination, nil
}

// SetReviewList caches a review list pagination result with the specified TTL in seconds
func (c *ReviewCache) SetReviewList(ctx context.Context, cacheKey string, pagination *application.Pagination[domain.Review]) error {
	if c.redisClient == nil {
		return nil
	}

	data, err := json.Marshal(pagination)
	if err != nil {
		return err
	}

	return c.redisClient.Set(ctx, cacheKey, data, time.Duration(CacheTTLReview)*time.Second).Err()
}

// InvalidateReviewList removes all cached review list entries
func (c *ReviewCache) InvalidateReviewList(ctx context.Context) error {
	if c.redisClient == nil {
		return nil
	}

	iter := c.redisClient.Scan(ctx, 0, ReviewListPrefix+"*", 0).Iterator()
	for iter.Next(ctx) {
		if err := c.redisClient.Del(ctx, iter.Val()).Err(); err != nil {
			return err
		}
	}
	return iter.Err()
}

// BuildListCacheKey builds a cache key for review list queries
func (c *ReviewCache) BuildListCacheKey(
	orderItemIDs *[]uuid.UUID,
	productVariantID *uuid.UUID,
	userIDs *[]uuid.UUID,
	deleted domain.DeletedParam,
	limit, page int,
) string {
	return ReviewListKey(
		orderItemIDs,
		productVariantID,
		userIDs,
		deleted,
		limit,
		page,
	)
}
