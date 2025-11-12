package cart

import (
	"backend/internal/domain/pagination"
	"backend/internal/domain/product"
)

type Model struct {
	ID    int
	Items *[]ItemsPaginationModel
}

type ItemModel struct {
	ID             int
	ProductVariant *product.VariantModel
	Quantity       int
}

type ItemsPaginationModel struct {
	Metadata *pagination.Metadata
	Items    *[]ItemModel
}
