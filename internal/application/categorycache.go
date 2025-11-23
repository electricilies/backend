package application

import (
	"context"

	"backend/internal/domain"

	"github.com/google/uuid"
)

// CategoryCache defines the interface for category caching operations
type CategoryCache interface {
	// GetCategory retrieves a cached category by ID
	GetCategory(ctx context.Context, categoryID uuid.UUID) (*domain.Category, error)

	// SetCategory caches a category with the specified TTL in seconds
	SetCategory(ctx context.Context, categoryID uuid.UUID, category *domain.Category) error

	// GetCategoryList retrieves a cached category list pagination result
	GetCategoryList(ctx context.Context, cacheKey string) (*Pagination[domain.Category], error)

	// SetCategoryList caches a category list pagination result with the specified TTL in seconds
	SetCategoryList(ctx context.Context, cacheKey string, pagination *Pagination[domain.Category]) error

	// InvalidateCategory removes the cached category by ID
	InvalidateCategory(ctx context.Context, categoryID uuid.UUID) error

	// InvalidateCategoryList removes all cached category list entries
	InvalidateCategoryList(ctx context.Context) error

	// BuildListCacheKey builds a cache key for category list queries
	BuildListCacheKey(
		search *string,
		limit, page int,
	) string
}
