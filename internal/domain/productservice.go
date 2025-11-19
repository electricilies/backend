package domain

import (
	"context"
)

type ProductService interface {
	Create(context.Context, CreateProductParam) (*Product, error)
	Update(context.Context, UpdateProductParam) (*Product, error)
	Get(context.Context, GetProductParam) (*Product, error)
	Delete(context.Context, DeleteProductParam) error
	AddVariants(context.Context, []CreateProductVariantParam) (*ProductVariant, error)
	UpdateVariant(context.Context, UpdateProductVariantParam) (*ProductVariant, error)
	UpdateOption(context.Context, UpdateProductOptionParam) (*ProductOption, error)
	CreateImages(context.Context, []CreateProductImageParam) ([]ProductImage, error)
}
