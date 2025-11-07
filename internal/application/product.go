package application

import (
	"context"

	"backend/internal/domain/product"
)

type Product interface {
	GetUploadImageURL(ctx context.Context) (string, error)
}

type productApp struct {
	productRepo product.Repository
}

func NewProduct(productRepo product.Repository) Product {
	return &productApp{
		productRepo: productRepo,
	}
}

func (a *productApp) GetUploadImageURL(ctx context.Context) (string, error) {
	return a.productRepo.GetUploadImageURL(ctx)
}
