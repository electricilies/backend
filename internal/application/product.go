package application

import (
	"context"

	"backend/internal/domain/product"
)

type Product interface {
	GetUploadImageURL(ctx context.Context) (*product.UploadImageURLModel, error)
	GetDeleteImageURL(ctx context.Context, id int) (string, error)
	ListProducts(ctx context.Context, queryParams *product.QueryParams) (*product.PaginationModel, error)
}

type ProductApp struct {
	productRepo product.Repository
}

func NewProduct(productRepo product.Repository) Product {
	return &ProductApp{
		productRepo: productRepo,
	}
}

func ProvideProduct(
	productRepo product.Repository,
) *ProductApp {
	return &ProductApp{
		productRepo: productRepo,
	}
}

func (a *ProductApp) GetUploadImageURL(ctx context.Context) (*product.UploadImageURLModel, error) {
	return a.productRepo.GetUploadImageURL(ctx)
}

func (a *ProductApp) GetDeleteImageURL(ctx context.Context, id int) (string, error) {
	return a.productRepo.GetDeleteImageURL(ctx, id)
}

func (a *ProductApp) ListProducts(ctx context.Context, queryParams *product.QueryParams) (*product.PaginationModel, error) {
	return a.productRepo.List(ctx, queryParams)
}
