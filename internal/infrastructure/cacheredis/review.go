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

// Review implements application.Review interface using Redis
type Review struct {
	redisClient *redis.Client
}

// ProvideReview creates a new ReviewCache instance
func ProvideReview(redisClient *redis.Client) *Review {
	return &Review{
		redisClient: redisClient,
	}
}

var _ application.ReviewCache = (*Review)(nil)

// GetReviewList retrieves a cached review list pagination result
func (c *Review) GetReviewList(ctx context.Context, cacheKey string) (*application.Pagination[domain.Review], error) {
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
func (c *Review) SetReviewList(ctx context.Context, cacheKey string, pagination *application.Pagination[domain.Review]) error {
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
func (c *Review) InvalidateReviewList(ctx context.Context) error {
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
func (c *Review) BuildListCacheKey(
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
