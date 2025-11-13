package product

import (
	"time"

	"backend/internal/domain/attribute"
	"backend/internal/domain/category"
	"backend/internal/domain/param"
)

type UploadImageURLModel struct {
	URL *string
	Key *string
}

type VariantOptionValueModel struct {
	ID    *int
	Value *string
}

type VariantImageModel struct {
	ID        *int
	URL       *string
	Order     *int
	CreatedAt *time.Time
}

type VariantModel struct {
	ID            *int
	SKU           *string
	Price         *int64
	Quantity      *int
	PurchaseCount *int
	CreatedAt     *time.Time
	DeletedAt     *time.Time
	OptionValue   *[]VariantOptionValueModel
	Images        *[]VariantImageModel
}

type ImageModel struct {
	ID               *int
	URL              *string
	Order            *int
	CreatedAt        *time.Time
	ProductVariantID *int
}

type Model struct {
	ID              *int
	Name            *string
	Description     *string
	ViewsCount      *int
	TotalPurchase   *int
	TrendingScore   *int64
	Rating          *float64
	Price           *int64
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
	Category        *category.Model
	AttributeValues *[]attribute.ValueModel
	Variants        *[]VariantModel
}

type PaginationModel struct {
	Metadata *param.PaginationMetadata
	Products *[]Model
}

// TODO: Add param later
type QueryParams struct {
	PaginationParams *param.Pagination
	Search           *string
	MinPrice         *int64
	MaxPrice         *int64
	CategoryIDs      *[]int
	Deleted          *param.Deleted
	SortPrice        *param.SortPrice
	SortRating       *param.SortRating
}
