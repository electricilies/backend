package application

import (
	"context"

	"backend/internal/domain/product"
)

type Product interface {
	GetUploadImageURL(context.Context) (*product.UploadImageURLModel, error)
	GetDeleteImageURL(context.Context, int) (string, error)
	ListProducts(context.Context, *product.QueryParams) (*product.PaginationModel, error)
}

type ProductImpl struct {
	productRepo product.Repository
}

func NewProduct(productRepo product.Repository) Product {
	return &ProductImpl{
		productRepo: productRepo,
	}
}

func ProvideProduct(
	productRepo product.Repository,
) *ProductImpl {
	return &ProductImpl{
		productRepo: productRepo,
	}
}

func (a *ProductImpl) GetUploadImageURL(ctx context.Context) (*product.UploadImageURLModel, error) {
	return a.productRepo.GetUploadImageURL(ctx)
}

func (a *ProductImpl) GetDeleteImageURL(ctx context.Context, id int) (string, error) {
	return a.productRepo.GetDeleteImageURL(ctx, id)
}

func (a *ProductImpl) ListProducts(ctx context.Context, queryParams *product.QueryParams) (*product.PaginationModel, error) {
	return a.productRepo.List(ctx, queryParams)
}
