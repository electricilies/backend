package application

import (
	"context"

	"backend/internal/domain"

	"github.com/google/uuid"
)

type Product interface {
	Create(context.Context, CreateProductParam) (*domain.Product, error)
	AddVariants(context.Context, AddProductVariantsParam) (*[]domain.ProductVariant, error)
	List(context.Context, ListProductParam) (*Pagination[domain.Product], error)
	GetDeleteImageURL(context.Context, uuid.UUID) (*DeleteImageURL, error)
	GetUploadImageURL(context.Context) (*UploadImageURL, error)
	Get(context.Context, GetProductParam) (*domain.Product, error)
	AddImages(context.Context, AddProductImagesParam) (*[]domain.ProductImage, error)
	Update(context.Context, UpdateProductParam) (*domain.Product, error)
	UpdateVariant(context.Context, UpdateProductVariantParam) (*domain.ProductVariant, error)
	UpdateOptions(context.Context, UpdateProductOptionsParam) (*[]domain.Option, error)
	UpdateOptionValues(context.Context, UpdateProductOptionValuesParam) (*[]domain.OptionValue, error)
	Delete(context.Context, DeleteProductParam) error
	DeleteImages(context.Context, DeleteProductImagesParam) error
}
