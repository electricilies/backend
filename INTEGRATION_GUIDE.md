# Redis Caching Integration Guide

## Quick Start

### 1. Update Dependency Injection

Your dependency injection code needs to be updated to pass the Redis client to the application layer provider functions.

**Before:**
```go
categoryApp := application.ProvideCategory(categoryRepo, categoryService)
attributeApp := application.ProvideAttribute(attributeRepo, attributeService)
reviewApp := application.ProvideReview(reviewRepo, reviewService)
```

**After:**
```go
// Initialize Redis client (already exists in your codebase)
redisClient := client.NewRedis(ctx, serverConfig)

// Pass Redis client to application providers
categoryApp := application.ProvideCategory(categoryRepo, categoryService, redisClient)
attributeApp := application.ProvideAttribute(attributeRepo, attributeService, redisClient)
reviewApp := application.ProvideReview(reviewRepo, reviewService, redisClient)
```

### 2. Files Modified

- `internal/application/categoryimpl.go` - Added Redis caching
- `internal/application/attributeimpl.go` - Added Redis caching
- `internal/application/reviewimpl.go` - Added Redis caching

### 3. Files Created

- `internal/constant/redis.go` - Redis key constants and TTL configurations

## Caching Behavior

### Cached Operations

| Module | Operation | Cache Key Pattern | TTL |
|--------|-----------|-------------------|-----|
| Category | List | `category:list:{search}:{limit}:{page}` | 1 hour |
| Category | Get | `category:get:{id}` | 1 hour |
| Attribute | List | `attribute:list:{ids}:{search}:{deleted}:{limit}:{page}` | 1 hour |
| Attribute | Get | `attribute:get:{id}` | 1 hour |
| Attribute Value | List | `attribute_value:list:{attributeId}:{valueIds}:{search}:{limit}:{page}` | 1 hour |
| Review | List | `review:list:{orderItemIds}:{variantId}:{userIds}:{deleted}:{limit}:{page}` | 30 min |

### Cache Invalidation

Cache is automatically invalidated on:
- **Create**: Invalidates all list caches
- **Update**: Invalidates specific get cache and all list caches
- **Delete**: Invalidates specific get cache (if applicable) and all list caches

## Configuration

### Adjusting TTL Values

Edit `internal/constant/redis.go`:

```go
const (
    CacheTTLCategory       = 3600  // 1 hour (in seconds)
    CacheTTLAttribute      = 3600  // 1 hour
    CacheTTLAttributeValue = 3600  // 1 hour
    CacheTTLReview         = 1800  // 30 minutes
)
```

### Disabling Cache

To disable caching temporarily without code changes, simply don't initialize or pass a nil Redis client:

```go
// This will work without caching
categoryApp := application.ProvideCategory(categoryRepo, categoryService, nil)
```

## Testing

### Unit Tests

When writing unit tests, mock the Redis client:

```go
import (
    "github.com/redis/go-redis/v9"
    "github.com/stretchr/testify/mock"
)

// Create a mock Redis client
mockRedis := &MockRedisClient{}
mockRedis.On("Get", mock.Anything, mock.Anything).Return(redis.NewStringResult("", redis.Nil))
mockRedis.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(redis.NewStatusResult("OK", nil))

// Pass to application
categoryApp := application.ProvideCategory(categoryRepo, categoryService, mockRedis)
```

### Integration Tests

Test cache behavior:

```go
// Test cache hit
result1, _ := categoryApp.List(ctx, param)
result2, _ := categoryApp.List(ctx, param) // Should hit cache

// Test cache invalidation
categoryApp.Create(ctx, createParam)
result3, _ := categoryApp.List(ctx, param) // Should miss cache, fetch fresh data
```

## Monitoring

### Cache Hit Rate

To monitor cache effectiveness, add metrics:

```go
// Example: Add to each cached method
cacheHits := 0
cacheMisses := 0

if cachedData != "" {
    cacheHits++
    // Return cached data
} else {
    cacheMisses++
    // Fetch from database
}

// Log or export metrics
hitRate := float64(cacheHits) / float64(cacheHits + cacheMisses)
```

### Redis Commands

Monitor cache keys in Redis CLI:

```bash
# List all category cache keys
redis-cli KEYS "category:*"

# List all attribute cache keys
redis-cli KEYS "attribute:*"

# List all review cache keys
redis-cli KEYS "review:*"

# Check TTL of a specific key
redis-cli TTL "category:get:123e4567-e89b-12d3-a456-426614174000"

# Manually clear all cache
redis-cli FLUSHDB
```

## Troubleshooting

### Issue: Stale Data

**Symptom**: Updates not reflected in API responses

**Solution**: Check cache invalidation logic is being called after save operations

```bash
# Verify cache is being invalidated
redis-cli MONITOR | grep DEL
```

### Issue: High Memory Usage

**Symptom**: Redis memory growing continuously

**Solution**: 
1. Verify TTL is set correctly
2. Reduce TTL values
3. Implement max memory policy in Redis config

```conf
# redis.conf
maxmemory 256mb
maxmemory-policy allkeys-lru
```

### Issue: Cache Miss on Every Request

**Symptom**: Cache never hits, always fetching from database

**Solution**: 
1. Check Redis connection
2. Verify cache keys are consistent
3. Check TTL hasn't expired immediately

```go
// Debug cache key generation
cacheKey := constant.CategoryListKey(searchStr, *param.Limit, *param.Page)
log.Printf("Cache key: %s", cacheKey)
```

## Performance Considerations

### Expected Improvements

- **Read-heavy workloads**: 50-90% reduction in database queries
- **Response time**: 10-50ms vs 100-500ms (database query)
- **Database load**: Significant reduction during traffic spikes

### Trade-offs

- **Consistency**: Eventual consistency (cache invalidation lag)
- **Memory**: Additional Redis memory usage
- **Complexity**: Cache invalidation logic to maintain

## Next Steps

1. ✅ Update dependency injection code
2. ✅ Deploy to staging environment
3. ⬜ Monitor cache hit rates
4. ⬜ Adjust TTL values based on usage patterns
5. ⬜ Add cache metrics/monitoring
6. ⬜ Consider adding cache warming for frequently accessed data
