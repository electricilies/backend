package cacheredis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"backend/internal/application"
	"backend/internal/delivery/http"

	"github.com/redis/go-redis/v9"
)

type Cart struct {
	redisClient *redis.Client
}

func ProvideCart(redisClient *redis.Client) *Cart {
	return &Cart{
		redisClient: redisClient,
	}
}

var _ application.CartCache = (*Cart)(nil)

func (c *Cart) Get(
	ctx context.Context,
	param application.CartCacheParam,
) (*http.CartResponseDto, error) {
	if c.redisClient == nil {
		return nil, toDomainError(ErrClientNil)
	}
	key := c.getKey(param)
	data, err := c.redisClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	if data == "" {
		return nil, redis.Nil
	}
	var result http.CartResponseDto
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Cart) Set(
	ctx context.Context,
	param application.CartCacheParam,
	cart *http.CartResponseDto,
) error {
	if c.redisClient == nil {
		return toDomainError(ErrClientNil)
	}
	key := c.getKey(param)
	data, err := json.Marshal(cart)
	if err != nil {
		return err
	}
	return c.redisClient.Set(ctx, key, data, time.Duration(CacheTTLCart)*time.Second).Err()
}

func (c *Cart) Invalidate(
	ctx context.Context,
	param application.CartCacheParam,
) error {
	if c.redisClient == nil {
		return toDomainError(ErrClientNil)
	}
	key := c.getKey(param)
	return c.redisClient.Del(ctx, key).Err()
}

func (c *Cart) InvalidateAlls(
	ctx context.Context,
) error {
	if c.redisClient == nil {
		return toDomainError(ErrClientNil)
	}
	pattern := CartGetPrefix + "*"
	iter := c.redisClient.Scan(ctx, 0, pattern, 0).Iterator()
	for iter.Next(ctx) {
		c.redisClient.Del(ctx, iter.Val())
	}
	if err := iter.Err(); err != nil {
		return err
	}
	return nil
}

func (c *Cart) getKey(param application.CartCacheParam) string {
	return fmt.Sprintf("%s%s", CartGetPrefix, param.ID.String())
}
