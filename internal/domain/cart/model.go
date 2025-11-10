package cart

import (
	"backend/internal/domain/pagination"
	"backend/internal/domain/product"
)

type Model struct {
	ID    int
	Items []ItemsPaginationPagination
}

type ItemModel struct {
	ID       int
	Product  product.Model
	Quantity int
}

type ItemsPaginationPagination struct {
	Metadata pagination.Metadata
	Items    []ItemModel
}
