package http

import (
	"context"

	"backend/internal/domain"

	"github.com/google/uuid"
)

type ProductApplication interface {
	Create(context.Context, CreateProductRequestDto) (*domain.Product, error)
	AddVariants(context.Context, AddProductVariantsRequestDto) (*[]domain.ProductVariant, error)
	List(context.Context, ListProductRequestDto) (*PaginationResponseDto[domain.Product], error)
	GetDeleteImageURL(context.Context, uuid.UUID) (*DeleteImageURLResponseDto, error)
	GetUploadImageURL(context.Context) (*UploadImageURLResponseDto, error)
	Get(context.Context, GetProductRequestDto) (*domain.Product, error)
	AddImages(context.Context, AddProductImagesRequestDto) (*[]domain.ProductImage, error)
	Update(context.Context, UpdateProductRequestDto) (*domain.Product, error)
	UpdateVariant(context.Context, UpdateProductVariantRequestDto) (*domain.ProductVariant, error)
	UpdateOptions(context.Context, UpdateProductOptionsRequestDto) (*[]domain.Option, error)
	UpdateOptionValues(context.Context, UpdateProductOptionValuesRequestDto) (*[]domain.OptionValue, error)
	Delete(context.Context, DeleteProductRequestDto) error
	DeleteImages(context.Context, DeleteProductImagesRequestDto) error
}
