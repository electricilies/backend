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

type Attribute struct {
	redisClient *redis.Client
}

func ProvideAttribute(redisClient *redis.Client) *Attribute {
	return &Attribute{
		redisClient: redisClient,
	}
}

var _ application.AttributeCache = (*Attribute)(nil)

func (a *Attribute) Get(
	ctx context.Context,
	param application.AttributeCacheParam,
) (*http.AttributeResponseDto, error) {
	key := a.getKey(param)
	data, err := a.redisClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	if data == "" {
		return nil, redis.Nil
	}
	var result http.AttributeResponseDto
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (a *Attribute) Set(
	ctx context.Context,
	param application.AttributeCacheParam,
	attribute *http.AttributeResponseDto,
) error {
	key := a.getKey(param)
	data, err := json.Marshal(attribute)
	if err != nil {
		return err
	}
	return a.redisClient.Set(ctx, key, data, time.Duration(CacheTTLAttribute)*time.Second).Err()
}

func (a *Attribute) Invalidate(
	ctx context.Context,
	param application.AttributeCacheParam,
) error {
	key := a.getKey(param)
	return a.redisClient.Del(ctx, key).Err()
}

func (a *Attribute) GetList(
	ctx context.Context,
	param application.AttributeCacheListParam,
) (*http.PaginationResponseDto[http.AttributeResponseDto], error) {
	key := a.getListKey(param)
	data, err := a.redisClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	if data == "" {
		return nil, redis.Nil
	}
	var result http.PaginationResponseDto[http.AttributeResponseDto]
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (a *Attribute) SetList(
	ctx context.Context,
	param application.AttributeCacheListParam,
	pagination *http.PaginationResponseDto[http.AttributeResponseDto],
) error {
	key := a.getListKey(param)
	data, err := json.Marshal(pagination)
	if err != nil {
		return err
	}
	return a.redisClient.Set(ctx, key, data, time.Duration(CacheTTLAttribute)*time.Second).Err()
}

func (a *Attribute) InvalidateList(
	ctx context.Context,
	param application.AttributeCacheListParam,
) error {
	key := a.getListKey(param)
	return a.redisClient.Del(ctx, key).Err()
}

func (a *Attribute) GetValueList(
	ctx context.Context,
	param application.AttributeCacheValueListParam,
) (*http.PaginationResponseDto[http.AttributeValueResponseDto], error) {
	key := a.getValueListKey(param)
	data, err := a.redisClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	if data == "" {
		return nil, redis.Nil
	}
	var result http.PaginationResponseDto[http.AttributeValueResponseDto]
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (a *Attribute) SetValueList(
	ctx context.Context,
	param application.AttributeCacheValueListParam,
	pagination *http.PaginationResponseDto[http.AttributeValueResponseDto],
) error {
	key := a.getValueListKey(param)
	data, err := json.Marshal(pagination)
	if err != nil {
		return err
	}
	return a.redisClient.Set(ctx, key, data, time.Duration(CacheTTLAttributeValue)*time.Second).Err()
}

func (a *Attribute) InvalidateValueList(
	ctx context.Context,
	param application.AttributeCacheValueListParam,
) error {
	key := a.getValueListKey(param)
	return a.redisClient.Del(ctx, key).Err()
}

func (a *Attribute) InvalidateAlls(
	ctx context.Context,
) error {
	patterns := []string{
		AttributeGetPrefix + "*",
		AttributeListPrefix + "*",
		AttributeValueListPrefix + "*",
	}
	for _, pattern := range patterns {
		iter := a.redisClient.Scan(ctx, 0, pattern, 0).Iterator()
		for iter.Next(ctx) {
			a.redisClient.Del(ctx, iter.Val())
		}
		if err := iter.Err(); err != nil {
			return err
		}
	}
	return nil
}

func (a *Attribute) getKey(param application.AttributeCacheParam) string {
	return fmt.Sprintf("%s%s", AttributeGetPrefix, param.ID.String())
}

func (a *Attribute) getListKey(param application.AttributeCacheListParam) string {
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
	parts = append(parts, fmt.Sprintf("deleted:%s", param.Deleted))
	parts = append(parts, fmt.Sprintf("limit:%d", param.Limit))
	parts = append(parts, fmt.Sprintf("page:%d", param.Page))
	return fmt.Sprintf("%s%s", AttributeListPrefix, strings.Join(parts, ":"))
}

func (a *Attribute) getValueListKey(param application.AttributeCacheValueListParam) string {
	var parts []string
	parts = append(parts, fmt.Sprintf("attr:%s", param.ID.String()))
	if len(param.ValueIDs) > 0 {
		ids := make([]string, len(param.ValueIDs))
		for i, id := range param.ValueIDs {
			ids[i] = id.String()
		}
		sort.Strings(ids)
		parts = append(parts, fmt.Sprintf("ids:%s", strings.Join(ids, ",")))
	}
	if param.Search != "" {
		parts = append(parts, fmt.Sprintf("search:%s", param.Search))
	}
	parts = append(parts, fmt.Sprintf("limit:%d", param.Limit))
	parts = append(parts, fmt.Sprintf("page:%d", param.Page))
	return fmt.Sprintf("%s%s", AttributeValueListPrefix, strings.Join(parts, ":"))
}
