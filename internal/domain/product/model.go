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

type OptionModel struct {
	ID     *int
	Name   *string
	Values *[]VariantOptionValueModel
}

type VariantOptionValueModel struct {
	ID    *int
	Value *string
}

type VariantModel struct {
	ID            *int
	SKU           *string
	Price         *int64
	Quantity      *int
	PurchaseCount *int
	CreatedAt     *time.Time
	DeletedAt     *time.Time
	UpdatedAt     *time.Time
	OptionValue   *[]VariantOptionValueModel
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
	TrendingScore   *float32
	Rating          *float32
	Price           *int64
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
	DeletedAt       *time.Time
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
	IDs              *[]int
	CategoryIDs      *[]int
	Rating           *float32
	Deleted          *param.Deleted
	SortPrice        *param.SortPrice
	SortRating       *param.SortRating
}
