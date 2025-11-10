package product

import (
	"time"

	"backend/internal/domain/pagination"
)

type UploadImageURLModel struct {
	URL string
	Key string
}

type VariantOptionValueModel struct {
	ID    int
	Value string
}

type VariantImageModel struct {
	ID        int
	URL       string
	Order     int
	CreatedAt time.Time
}

type VariantModel struct {
	ID            int
	SKU           string
	Price         int64
	Quantity      int
	PurchaseCount int
	CreatedAt     time.Time
	DeletedAt     *time.Time
	OptionValue   []VariantOptionValueModel
	Images        *[]VariantImageModel
}

type ImageModel struct {
	ID               int
	URL              string
	Order            int
	CreatedAt        time.Time
	ProductVariantID *int
}

type Model struct {
	ID              int
	Name            string
	Description     string
	ViewsCount      int
	TotalPurchase   int
	TrendingScore   int64
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Category        Category
	AttributeValues []AttributeValue
	Variants        []VariantModel
}

// TODO: move to attribute package
type Attribute struct {
	ID              string
	Code            string
	Name            string
	AttributeValues []AttributeValue
}

type AttributeValue struct {
	ID    string
	Value string
}

// TODO: move to category package
type Category struct {
	ID          int
	Name        string
	Description string
	CreatedAt   time.Time
}

type Pagination struct {
	Metadata pagination.Metadata
	Products []Model
}

// TODO: Add params later
type QueryParams struct {
	PaginationParams pagination.Params
}
