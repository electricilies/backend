package cacheredis

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"

	"backend/internal/application"
	"backend/internal/delivery/http"

	"github.com/redis/go-redis/v9"
)

type Product struct {
	redisClient *redis.Client
}

func ProvideProduct(redisClient *redis.Client) *Product {
	return &Product{
		redisClient: redisClient,
	}
}

var _ application.ProductCache = (*Product)(nil)

func (p *Product) Get(
	ctx context.Context,
	param application.ProductCacheParam,
) (*http.ProductResponseDto, error) {
	key := p.getKey(param)
	data, err := p.redisClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	if data == "" {
		return nil, redis.Nil
	}
	var result http.ProductResponseDto
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (p *Product) Set(
	ctx context.Context,
	param application.ProductCacheParam,
	product *http.ProductResponseDto,
) error {
	key := p.getKey(param)
	data, err := json.Marshal(product)
	if err != nil {
		return err
	}
	return p.redisClient.Set(ctx, key, data, time.Duration(CacheTTLProduct)*time.Second).Err()
}

func (p *Product) Invalidate(
	ctx context.Context,
	param application.ProductCacheParam,
) error {
	key := p.getKey(param)
	return p.redisClient.Del(ctx, key).Err()
}

func (p *Product) GetList(
	ctx context.Context,
	param application.ProductCacheListParam,
) (*http.PaginationResponseDto[http.ProductResponseDto], error) {
	key := p.getListKey(param)
	data, err := p.redisClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	if data == "" {
		return nil, redis.Nil
	}
	var result http.PaginationResponseDto[http.ProductResponseDto]
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (p *Product) SetList(
	ctx context.Context,
	param application.ProductCacheListParam,
	pagination *http.PaginationResponseDto[http.ProductResponseDto],
) error {
	key := p.getListKey(param)
	data, err := json.Marshal(pagination)
	if err != nil {
		return err
	}
	return p.redisClient.Set(ctx, key, data, time.Duration(CacheTTLProduct)*time.Second).Err()
}

func (p *Product) InvalidateList(
	ctx context.Context,
	param application.ProductCacheListParam,
) error {
	key := p.getListKey(param)
	return p.redisClient.Del(ctx, key).Err()
}

func (p *Product) InvalidateAlls(
	ctx context.Context,
) error {
	patterns := []string{
		ProductGetPrefix + "*",
		ProductListPrefix + "*",
	}
	for _, pattern := range patterns {
		iter := p.redisClient.Scan(ctx, 0, pattern, 0).Iterator()
		for iter.Next(ctx) {
			p.redisClient.Del(ctx, iter.Val())
		}
		if err := iter.Err(); err != nil {
			return err
		}
	}
	return nil
}

func (p *Product) getKey(param application.ProductCacheParam) string {
	return fmt.Sprintf("%s%s", ProductGetPrefix, param.ID.String())
}

func (p *Product) getListKey(param application.ProductCacheListParam) string {
	var parts []string
	if len(param.IDs) > 0 {
		ids := make([]string, len(param.IDs))
		for i, id := range param.IDs {
			ids[i] = id.String()
		}
		sort.Strings(ids)
		parts = append(parts, fmt.Sprintf("ids:%s", strings.Join(ids, ",")))
	}
	if param.Search != "" {
		parts = append(parts, fmt.Sprintf("search:%s", param.Search))
	}
	if param.MinPrice != 0 {
		parts = append(parts, fmt.Sprintf("min_price:%d", param.MinPrice))
	}
	if param.MaxPrice != 0 {
		parts = append(parts, fmt.Sprintf("max_price:%d", param.MaxPrice))
	}
	if param.Rating != 0 {
		parts = append(parts, fmt.Sprintf("rating:%.2f", param.Rating))
	}
	if len(param.CategoryIDs) > 0 {
		ids := make([]string, len(param.CategoryIDs))
		for i, id := range param.CategoryIDs {
			ids[i] = id.String()
		}
		sort.Strings(ids)
		parts = append(parts, fmt.Sprintf("category_ids:%s", strings.Join(ids, ",")))
	}
	parts = append(parts, fmt.Sprintf("deleted:%s", param.Deleted))
	if param.SortRating != "" {
		parts = append(parts, fmt.Sprintf("sort_rating:%s", param.SortRating))
	}
	if param.SortPrice != "" {
		parts = append(parts, fmt.Sprintf("sort_price:%s", param.SortPrice))
	}
	parts = append(parts, fmt.Sprintf("limit:%d", param.Limit))
	parts = append(parts, fmt.Sprintf("page:%d", param.Page))
	return fmt.Sprintf("%s%s", ProductListPrefix, strings.Join(parts, ":"))
}
