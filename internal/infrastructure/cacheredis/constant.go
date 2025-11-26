package cacheredis

import (
	"fmt"

	"backend/internal/domain"

	"github.com/google/uuid"
)

const (
	// Cache TTL durations (in seconds)
	CacheTTLCategory       = 3600 // 1 hour
	CacheTTLAttribute      = 3600 // 1 hour
	CacheTTLAttributeValue = 3600 // 1 hour
	CacheTTLReview         = 1800 // 30 minutes
	CacheTTLProduct        = 3600 // 1 hour
	CacheTTLCart           = 1800 // 30 minutes
)

// Category cache keys
const (
	CategoryListPrefix = "category:list:"
	CategoryGetPrefix  = "category:get:"
)

func CategoryListKey(
	search *string,
	limit, page int,
) string {
	var searchStr string
	if search != nil {
		searchStr = *search
	}
	return fmt.Sprintf("%s%s:%d:%d", CategoryListPrefix, searchStr, limit, page)
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

func AttributeListKey(
	ids *[]uuid.UUID,
	search *string,
	deleted domain.DeletedParam,
	limit, page int,
) string {
	var idsStr string
	if ids != nil && len(*ids) > 0 {
		for _, id := range *ids {
			idsStr += id.String() + ","
		}
	}
	var searchStr string
	if search != nil {
		searchStr = *search
	}
	return fmt.Sprintf(
		"%s%s:%s:%s:%d:%d",
		AttributeListPrefix,
		idsStr,
		searchStr,
		deleted,
		limit,
		page,
	)
}

func AttributeGetKey(id uuid.UUID) string {
	return fmt.Sprintf("%s%s", AttributeGetPrefix, id.String())
}

func AttributeValueListKey(
	attributeID uuid.UUID,
	valueIDs *[]uuid.UUID,
	search *string,
	limit, page int,
) string {
	var valueIDsStr string
	if valueIDs != nil && len(*valueIDs) > 0 {
		for _, id := range *valueIDs {
			valueIDsStr += id.String() + ","
		}
	}
	var searchStr string
	if search != nil {
		searchStr = *search
	}
	return fmt.Sprintf(
		"%s%s:%s:%s:%d:%d",
		AttributeValueListPrefix,
		attributeID.String(),
		valueIDsStr,
		searchStr,
		limit,
		page,
	)
}

// Review cache keys
const (
	ReviewListPrefix = "review:list:"
)

func ReviewListKey(
	orderItemIDs *[]uuid.UUID,
	productVariantID *uuid.UUID,
	userIDs *[]uuid.UUID,
	deleted domain.DeletedParam,
	limit, page int,
) string {
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

	return fmt.Sprintf(
		"%s%s:%s:%s:%s:%d:%d",
		ReviewListPrefix,
		orderItemIDsStr,
		productVariantIDStr,
		userIDsStr,
		deleted,
		limit,
		page,
	)
}

// Product cache keys
const (
	ProductListPrefix = "product:list:"
	ProductGetPrefix  = "product:get:"
)

func ProductListKey(
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
) string {
	var idsStr string
	if ids != nil && len(*ids) > 0 {
		for _, id := range *ids {
			idsStr += id.String() + ","
		}
	}

	var searchStr string
	if search != nil {
		searchStr = *search
	}

	var minPriceStr string
	if minPrice != nil {
		minPriceStr = fmt.Sprintf("%d", *minPrice)
	}

	var maxPriceStr string
	if maxPrice != nil {
		maxPriceStr = fmt.Sprintf("%d", *maxPrice)
	}

	var ratingStr string
	if rating != nil {
		ratingStr = fmt.Sprintf("%.2f", *rating)
	}

	var categoryIDsStr string
	if categoryIDs != nil && len(*categoryIDs) > 0 {
		for _, id := range *categoryIDs {
			categoryIDsStr += id.String() + ","
		}
	}

	var sortRatingStr string
	if sortRating != nil {
		sortRatingStr = *sortRating
	}

	var sortPriceStr string
	if sortPrice != nil {
		sortPriceStr = *sortPrice
	}

	return fmt.Sprintf(
		"%s%s:%s:%s:%s:%s:%s:%s:%s:%s:%d:%d",
		ProductListPrefix,
		idsStr,
		searchStr,
		minPriceStr,
		maxPriceStr,
		ratingStr,
		categoryIDsStr,
		deleted,
		sortRatingStr,
		sortPriceStr,
		limit,
		page,
	)
}

func ProductGetKey(id uuid.UUID) string {
	return fmt.Sprintf("%s%s", ProductGetPrefix, id.String())
}

// Cart cache keys
const (
	CartGetPrefix  = "cart:get:"
	CartUserPrefix = "cart:user:"
)

func CartGetKey(id uuid.UUID) string {
	return fmt.Sprintf("%s%s", CartGetPrefix, id.String())
}

func CartUserKey(userID uuid.UUID) string {
	return fmt.Sprintf("%s%s", CartUserPrefix, userID.String())
}
