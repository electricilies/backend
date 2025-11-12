package cart

import (
	"backend/internal/domain/params"
	"backend/internal/domain/product"
)

type Model struct {
	ID    int
	Items *[]ItemsPaginationModel
}

type ItemModel struct {
	ID             string
	ProductVariant *product.VariantModel
	Quantity       int
}

type ItemsPaginationModel struct {
	Metadata *params.Metadata
	Items    *[]ItemModel
}
