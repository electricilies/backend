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

type Category struct {
	redisClient *redis.Client
}

func ProvideCategory(redisClient *redis.Client) *Category {
	return &Category{
		redisClient: redisClient,
	}
}

var _ application.CategoryCache = (*Category)(nil)

func (c *Category) Get(
	ctx context.Context,
	param application.CategoryCacheParam,
) (*http.CategoryResponseDto, error) {
	key := c.getKey(param)
	data, err := c.redisClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	if data == "" {
		return nil, redis.Nil
	}
	var result http.CategoryResponseDto
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Category) Set(
	ctx context.Context,
	param application.CategoryCacheParam,
	category *http.CategoryResponseDto,
) error {
	key := c.getKey(param)
	data, err := json.Marshal(category)
	if err != nil {
		return err
	}
	return c.redisClient.Set(ctx, key, data, time.Duration(CacheTTLCategory)*time.Second).Err()
}

func (c *Category) Invalidate(
	ctx context.Context,
	param application.CategoryCacheParam,
) error {
	key := c.getKey(param)
	return c.redisClient.Del(ctx, key).Err()
}

func (c *Category) GetList(
	ctx context.Context,
	param application.CategoryCacheListParam,
) (*http.PaginationResponseDto[http.CategoryResponseDto], error) {
	key := c.getListKey(param)
	data, err := c.redisClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	if data == "" {
		return nil, redis.Nil
	}
	var result http.PaginationResponseDto[http.CategoryResponseDto]
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Category) SetList(
	ctx context.Context,
	param application.CategoryCacheListParam,
	pagination *http.PaginationResponseDto[http.CategoryResponseDto],
) error {
	key := c.getListKey(param)
	data, err := json.Marshal(pagination)
	if err != nil {
		return err
	}
	return c.redisClient.Set(ctx, key, data, time.Duration(CacheTTLCategory)*time.Second).Err()
}

func (c *Category) InvalidateList(
	ctx context.Context,
	param application.CategoryCacheListParam,
) error {
	key := c.getListKey(param)
	return c.redisClient.Del(ctx, key).Err()
}

func (c *Category) InvalidateAlls(
	ctx context.Context,
) error {
	patterns := []string{
		CategoryGetPrefix + "*",
		CategoryListPrefix + "*",
	}
	for _, pattern := range patterns {
		iter := c.redisClient.Scan(ctx, 0, pattern, 0).Iterator()
		for iter.Next(ctx) {
			c.redisClient.Del(ctx, iter.Val())
		}
		if err := iter.Err(); err != nil {
			return err
		}
	}
	return nil
}

func (c *Category) getKey(param application.CategoryCacheParam) string {
	return fmt.Sprintf("%s%s", CategoryGetPrefix, param.ID.String())
}

func (c *Category) getListKey(param application.CategoryCacheListParam) string {
	var parts string
	if param.Search != "" {
		parts = fmt.Sprintf("search:%s:", param.Search)
	}
	parts += fmt.Sprintf("limit:%d:page:%d", param.Limit, param.Page)
	return fmt.Sprintf("%s%s", CategoryListPrefix, parts)
}
