package constant

import (
	"fmt"

	"github.com/google/uuid"
)

const (
	// Cache TTL durations (in seconds)
	CacheTTLCategory       = 3600 // 1 hour
	CacheTTLAttribute      = 3600 // 1 hour
	CacheTTLAttributeValue = 3600 // 1 hour
	CacheTTLReview         = 1800 // 30 minutes
	CacheTTLProduct        = 3600 // 1 hour
)

// Category cache keys
const (
	CategoryListPrefix = "category:list:"
	CategoryGetPrefix  = "category:get:"
)

func CategoryListKey(search string, limit, page int) string {
	return fmt.Sprintf("%s%s:%d:%d", CategoryListPrefix, search, limit, page)
}

func CategoryGetKey(id uuid.UUID) string {
	return fmt.Sprintf("%s%s", CategoryGetPrefix, id.String())
}

// Attribute cache keys
const (
	AttributeListPrefix      = "attribute:list:"
	AttributeGetPrefix       = "attribute:get:"
	AttributeValueListPrefix = "attribute_value:list:"
)

func AttributeListKey(ids *[]uuid.UUID, search string, deleted string, limit, page int) string {
	idsStr := ""
	if ids != nil && len(*ids) > 0 {
		for _, id := range *ids {
			idsStr += id.String() + ","
		}
	}
	return fmt.Sprintf("%s%s:%s:%s:%d:%d", AttributeListPrefix, idsStr, search, deleted, limit, page)
}

func AttributeGetKey(id uuid.UUID) string {
	return fmt.Sprintf("%s%s", AttributeGetPrefix, id.String())
}

func AttributeValueListKey(attributeID uuid.UUID, valueIDs *[]uuid.UUID, search string, limit, page int) string {
	valueIDsStr := ""
	if valueIDs != nil && len(*valueIDs) > 0 {
		for _, id := range *valueIDs {
			valueIDsStr += id.String() + ","
		}
	}
	return fmt.Sprintf("%s%s:%s:%s:%d:%d", AttributeValueListPrefix, attributeID.String(), valueIDsStr, search, limit, page)
}

// Review cache keys
const (
	ReviewListPrefix = "review:list:"
)

func ReviewListKey(orderItemIDs *[]uuid.UUID, productVariantID *uuid.UUID, userIDs *[]uuid.UUID, deleted string, limit, page int) string {
	orderItemIDsStr := ""
	if orderItemIDs != nil && len(*orderItemIDs) > 0 {
		for _, id := range *orderItemIDs {
			orderItemIDsStr += id.String() + ","
		}
	}

	productVariantIDStr := ""
	if productVariantID != nil {
		productVariantIDStr = productVariantID.String()
	}

	userIDsStr := ""
	if userIDs != nil && len(*userIDs) > 0 {
		for _, id := range *userIDs {
			userIDsStr += id.String() + ","
		}
	}

	return fmt.Sprintf("%s%s:%s:%s:%s:%d:%d", ReviewListPrefix, orderItemIDsStr, productVariantIDStr, userIDsStr, deleted, limit, page)
}

// Product cache keys
const (
	ProductListPrefix = "product:list:"
	ProductGetPrefix  = "product:get:"
)

func ProductListKey(ids *[]uuid.UUID, search string, minPrice *int64, maxPrice *int64, rating *float64, categoryIDs *[]uuid.UUID, deleted string, sortRating string, sortPrice string, limit, page int) string {
	idsStr := ""
	if ids != nil && len(*ids) > 0 {
		for _, id := range *ids {
			idsStr += id.String() + ","
		}
	}

	minPriceStr := ""
	if minPrice != nil {
		minPriceStr = fmt.Sprintf("%d", *minPrice)
	}

	maxPriceStr := ""
	if maxPrice != nil {
		maxPriceStr = fmt.Sprintf("%d", *maxPrice)
	}

	ratingStr := ""
	if rating != nil {
		ratingStr = fmt.Sprintf("%.2f", *rating)
	}

	categoryIDsStr := ""
	if categoryIDs != nil && len(*categoryIDs) > 0 {
		for _, id := range *categoryIDs {
			categoryIDsStr += id.String() + ","
		}
	}

	return fmt.Sprintf("%s%s:%s:%s:%s:%s:%s:%s:%s:%s:%d:%d", ProductListPrefix, idsStr, search, minPriceStr, maxPriceStr, ratingStr, categoryIDsStr, deleted, sortRating, sortPrice, limit, page)
}

func ProductGetKey(id uuid.UUID) string {
	return fmt.Sprintf("%s%s", ProductGetPrefix, id.String())
}
