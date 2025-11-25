package application

import (
	"context"

	"backend/internal/delivery/http"
	"backend/internal/domain"

	"github.com/google/uuid"
)

// ProductCache defines the interface for product caching operations
type ProductCache interface {
	// GetProduct retrieves a cached product by ID
	GetProduct(ctx context.Context, productID uuid.UUID) (*http.ProductResponseDto, error)

	// SetProduct caches a product with the specified TTL in seconds
	SetProduct(ctx context.Context, productID uuid.UUID, product *http.ProductResponseDto) error

	// GetProductList retrieves a cached product list pagination result
	GetProductList(ctx context.Context, cacheKey string) (*http.PaginationResponseDto[http.ProductResponseDto], error)

	// SetProductList caches a product list pagination result with the specified TTL in seconds
	SetProductList(ctx context.Context, cacheKey string, pagination *http.PaginationResponseDto[http.ProductResponseDto]) error

	// InvalidateProduct removes the cached product by ID
	InvalidateProduct(ctx context.Context, productID uuid.UUID) error

	// InvalidateProductList removes all cached product list entries
	InvalidateProductList(ctx context.Context) error

	// InvalidateAllProducts removes all product-related caches (both get and list)
	InvalidateAllProducts(ctx context.Context) error

	// BuildListCacheKey builds a cache key for product list queries
	BuildListCacheKey(
		ids *[]uuid.UUID,
		search *string,
		minPrice *int64,
		maxPrice *int64,
		rating *float64,
		categoryIDs *[]uuid.UUID,
		deleted domain.DeletedParam,
		sortRating *string,
		sortPrice *string,
		limit, page int,
	) string
}
