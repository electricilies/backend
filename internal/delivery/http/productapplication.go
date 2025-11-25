package http

import (
	"context"

	"github.com/google/uuid"
)

type ProductApplication interface {
	Create(context.Context, CreateProductRequestDto) (*ProductResponseDto, error)
	AddVariants(context.Context, AddProductVariantsRequestDto) (*[]ProductVariantResponseDto, error)
	List(context.Context, ListProductRequestDto) (*PaginationResponseDto[ProductResponseDto], error)
	GetDeleteImageURL(context.Context, uuid.UUID) (*DeleteImageURLResponseDto, error)
	GetUploadImageURL(context.Context) (*UploadImageURLResponseDto, error)
	Get(context.Context, GetProductRequestDto) (*ProductResponseDto, error)
	AddImages(context.Context, AddProductImagesRequestDto) (*[]ProductImageResponseDto, error)
	Update(context.Context, UpdateProductRequestDto) (*ProductResponseDto, error)
	UpdateVariant(context.Context, UpdateProductVariantRequestDto) (*ProductVariantResponseDto, error)
	UpdateOptions(context.Context, UpdateProductOptionsRequestDto) (*[]ProductOptionResponseDto, error)
	UpdateOptionValues(context.Context, UpdateProductOptionValuesRequestDto) (*[]ProductOptionValueResponseDto, error)
	Delete(context.Context, DeleteProductRequestDto) error
	DeleteImages(context.Context, DeleteProductImagesRequestDto) error
}
