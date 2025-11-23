package application

import (
	"context"

	"backend/internal/domain"

	"github.com/google/uuid"
)

// ReviewCache defines the interface for review caching operations
type ReviewCache interface {
	// GetReviewList retrieves a cached review list pagination result
	GetReviewList(ctx context.Context, cacheKey string) (*Pagination[domain.Review], error)

	// SetReviewList caches a review list pagination result with the specified TTL in seconds
	SetReviewList(ctx context.Context, cacheKey string, pagination *Pagination[domain.Review]) error

	// InvalidateReviewList removes all cached review list entries
	InvalidateReviewList(ctx context.Context) error

	// BuildListCacheKey builds a cache key for review list queries
	BuildListCacheKey(
		orderItemIDs *[]uuid.UUID,
		productVariantID *uuid.UUID,
		userIDs *[]uuid.UUID,
		deleted domain.DeletedParam,
		limit, page int,
	) string
}
