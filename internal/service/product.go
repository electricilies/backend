package service

import (
	"context"

	"backend/internal/domain"
)

type Product interface {
	Create(context.Context, CreateProductParam) (*domain.Product, error)
	Update(context.Context, UpdateProductParam) (*domain.Product, error)
	List(context.Context, ListProductParam) (*Pagination[domain.Product], error)
	Get(context.Context, GetProductParam) (*domain.Product, error)
	Delete(context.Context, DeleteProductParam) error
	CreateOptions(context.Context, []CreateProductOptionParam) (*domain.ProductOption, error)
	CreateVariants(context.Context, []CreateProductVariantParam) (*domain.ProductVariant, error)
	UpdateVariant(context.Context, UpdateProductVariantParam) (*domain.ProductVariant, error)
	UpdateOption(context.Context, UpdateProductOptionParam) (*domain.ProductOption, error)
	GetDeleteImageURL(context.Context, int) (*domain.ProductImageDeleteURL, error)
	GetUploadImageURL(context.Context) (*domain.ProductUploadURLImage, error)
	CreateImages(context.Context, []CreateProductImageParam) ([]domain.ProductImage, error)
}
