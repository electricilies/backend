package application

import (
	"context"

	"backend/internal/domain"
)

type Product interface {
	Create(context.Context, CreateProductParam) (*domain.Product, error)
	AddVariants(context.Context, AddProductVariantsParam) (*domain.ProductVariant, error)
	List(context.Context, ListProductParam) (*Pagination[domain.Product], error)
	GetDeleteImageURL(context.Context, int) (*DeleteImageURL, error)
	GetUploadImageURL(context.Context) (*UploadImageURL, error)
	Get(context.Context, GetProductParam) (*domain.Product, error)
	AddImages(context.Context, AddProductImagesParam) (*[]domain.ProductImage, error)
	Update(context.Context, UpdateProductParam) (*Product, error)
	UpdateVariant(context.Context, UpdateProductVariantParam) (*domain.ProductVariant, error)
	UpdateOptions(context.Context, UpdateProductOptionsParam) error
	UpdateOptionValues(context.Context, UpdateProductOptionValuesParam) error
	Delete(context.Context, DeleteProductParam) error
	DeleteImages(context.Context, DeleteProductImagesParam) error
}
