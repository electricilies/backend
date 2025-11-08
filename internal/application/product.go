package application

import (
	"backend/internal/domain/product"
	"context"
)

type Product interface {
	GetUploadImageURL(ctx context.Context) (string, error)
	GetDeleteImageURL(ctx context.Context, id int) (string, error)
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

func (a *productApp) GetDeleteImageURL(ctx context.Context, id int) (string, error) {
	return a.productRepo.GetDeleteImageURL(ctx, id)
}
