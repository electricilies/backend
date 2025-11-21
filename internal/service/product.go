package service

import (
	"backend/internal/domain"
)

type Product struct{}

func ProvideProduct() *Product {
	return &Product{}
}

var _ domain.ProductService = &Product{}
