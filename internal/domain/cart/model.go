package cart

import (
	"backend/internal/domain/pagination"
	"backend/internal/domain/product"
)

type Model struct {
	ID    int
	Items *[]ItemsPagination
}

type ItemModel struct {
	ID       int
	Product  *product.Model
	Quantity int
}

type ItemsPagination struct {
	Metadata *pagination.Metadata
	Items    *[]ItemModel
}
