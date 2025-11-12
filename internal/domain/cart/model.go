package cart

import (
	"backend/internal/domain/param"
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
	Metadata *param.Metadata
	Items    *[]ItemModel
}
