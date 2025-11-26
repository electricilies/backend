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

// Cart implements application.CartCache interface using Redis
type Cart struct {
	redisClient *redis.Client
}

// ProvideCart creates a new CartCache instance
func ProvideCart(redisClient *redis.Client) *Cart {
	return &Cart{
		redisClient: redisClient,
	}
}

var _ application.CartCache = (*Cart)(nil)

// GetCart retrieves a cached cart by ID
func (c *Cart) GetCart(ctx context.Context, cartID uuid.UUID) (*http.CartResponseDto, error) {
	if c.redisClient == nil {
		return nil, redis.Nil
	}

	cacheKey := CartGetKey(cartID)
	cachedData, err := c.redisClient.Get(ctx, cacheKey).Result()
	if err != nil {
		return nil, err
	}

	if cachedData == "" {
		return nil, redis.Nil
	}

	var cart http.CartResponseDto
	if err := json.Unmarshal([]byte(cachedData), &cart); err != nil {
		return nil, err
	}

	return &cart, nil
}

// SetCart caches a cart with the specified TTL in seconds
func (c *Cart) SetCart(ctx context.Context, cartID uuid.UUID, cart *http.CartResponseDto) error {
	if c.redisClient == nil {
		return nil
	}

	cacheKey := CartGetKey(cartID)
	data, err := json.Marshal(cart)
	if err != nil {
		return err
	}

	return c.redisClient.Set(ctx, cacheKey, data, time.Duration(CacheTTLCart)*time.Second).Err()
}

// InvalidateCart removes the cached cart by ID
func (c *Cart) InvalidateCart(ctx context.Context, cartID uuid.UUID) error {
	if c.redisClient == nil {
		return nil
	}

	cacheKey := CartGetKey(cartID)
	return c.redisClient.Del(ctx, cacheKey).Err()
}

// InvalidateUserCart removes the cached cart for a specific user
func (c *Cart) InvalidateUserCart(ctx context.Context, userID uuid.UUID) error {
	if c.redisClient == nil {
		return nil
	}

	// Scan and delete all cart entries for this user
	iter := c.redisClient.Scan(ctx, 0, CartUserPrefix+userID.String()+"*", 0).Iterator()
	for iter.Next(ctx) {
		if err := c.redisClient.Del(ctx, iter.Val()).Err(); err != nil {
			return err
		}
	}
	return iter.Err()
}
