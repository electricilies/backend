package application

import (
	"context"

	"backend/internal/delivery/http"
	"backend/internal/domain"

	"github.com/google/uuid"
)

// AttributeCache defines the interface for attribute caching operations
type AttributeCache interface {
	// GetAttribute retrieves a cached attribute by ID
	GetAttribute(ctx context.Context, attributeID uuid.UUID) (*http.AttributeResponseDto, error)

	// SetAttribute caches an attribute with the specified TTL in seconds
	SetAttribute(ctx context.Context, attributeID uuid.UUID, attribute *http.AttributeResponseDto) error

	// GetAttributeList retrieves a cached attribute list pagination result
	GetAttributeList(ctx context.Context, cacheKey string) (*http.PaginationResponseDto[http.AttributeResponseDto], error)

	// SetAttributeList caches an attribute list pagination result with the specified TTL in seconds
	SetAttributeList(ctx context.Context, cacheKey string, pagination *http.PaginationResponseDto[http.AttributeResponseDto]) error

	// GetAttributeValueList retrieves a cached attribute value list pagination result
	GetAttributeValueList(ctx context.Context, cacheKey string) (*http.PaginationResponseDto[http.AttributeValueResponseDto], error)

	// SetAttributeValueList caches an attribute value list pagination result with the specified TTL in seconds
	SetAttributeValueList(ctx context.Context, cacheKey string, pagination *http.PaginationResponseDto[http.AttributeValueResponseDto]) error

	// InvalidateAttribute removes the cached attribute by ID
	InvalidateAttribute(ctx context.Context, attributeID uuid.UUID) error

	// InvalidateAttributeList removes all cached attribute list entries
	InvalidateAttributeList(ctx context.Context) error

	// InvalidateAttributeValueList removes all cached attribute value list entries
	InvalidateAttributeValueList(ctx context.Context) error

	// InvalidateAllAttributes removes all attribute-related caches (attribute, list, and value list)
	InvalidateAllAttributes(ctx context.Context) error

	// BuildListCacheKey builds a cache key for attribute list queries
	BuildListCacheKey(
		ids *[]uuid.UUID,
		search *string,
		deleted domain.DeletedParam,
		limit, page int,
	) string

	// BuildValueListCacheKey builds a cache key for attribute value list queries
	BuildValueListCacheKey(
		attributeID uuid.UUID,
		valueIDs *[]uuid.UUID,
		search *string,
		limit, page int,
	) string
}
