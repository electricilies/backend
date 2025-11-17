package service

import (
	"context"

	"backend/internal/domain"
)

type CreateProductParam struct {
	Name              string                        `json:"name" binding:"required"`
	Description       string                        `json:"description,omitempty"`
	CategoryIDs       []int                         `json:"categoryIds,omitempty"`
	AttributeValueIDs []int                         `json:"attributeValueIds,omitempty"`
	ProductOption     []CreateProductOptionParam  `json:"productOption,omitempty"`
	Category          int                           `json:"category" binding:"required"`
	ProductVariants   []CreateProductVariantParam `json:"productVariants" binding:"required"`
	ProductImages     []CreateProductImageParam   `json:"productImages" binding:"required"`
}

type CreateProductOptionParam struct {
	Option string   `json:"option" binding:"required"`
	Value  []string `json:"value" binding:"required"`
}

type UpdateProductOptionParam struct {
	Name *string `json:"name" binding:"required"`
}

type CreateProductVariantParam struct {
	SKU                 string   `json:"sku" binding:"required"`
	Price               int64    `json:"price" binding:"required"`
	Quantity            int      `json:"quantity" binding:"required"`
	ProductOptionValues []string `json:"productOptionValues,omitempty"`
}

type UpdateProductVariantParam struct {
	Price    *int64 `json:"price,omitempty"`
	Quantity *int   `json:"quantity,omitempty"`
}

type CreateProductImageParam struct {
	URL              string `json:"url" binding:"required"`
	Order            int    `json:"order,omitempty"`
	ProductVariantID int    `json:"productVariantId,omitempty"`
	ProductID        int    `json:"productId,omitempty"`
}

type UpdateProductParam struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	CategoryID  *int    `json:"categoryId,omitempty"`
}

type ListProductParam struct {
	Limit       int
	Offset      int
	CategoryIDs []int
	MinPrice    int64
	MaxPrice    int64
	SortPrice   string
	SortRating  string
	Search      string
	Deleted     string
}

type GetProductParam struct {
	ProductID int `json:"productId" binding:"required"`
}

type DeleteProductParam struct {
	ProductID int `json:"productId" binding:"required"`
}

type Product interface {
	CreateProduct(context.Context, CreateProductParam) (*domain.Product, error)
	UpdateProduct(context.Context, UpdateProductParam) (*domain.Product, error)
	ListProducts(context.Context, ListProductParam) (*domain.DataPagination, error)
	GetProduct(context.Context, GetProductParam) (*domain.Product, error)
	DeleteProduct(context.Context, DeleteProductParam) error
	CreateProductOptions(context.Context, []CreateProductOptionParam) (*domain.ProductOption, error)
	CreateProductVariants(context.Context, []CreateProductVariantParam) (*domain.ProductVariant, error)
	UpdateProductVariants(context.Context, []UpdateProductVariantParam) (*domain.ProductVariant, error)
	UpdateProductOption(context.Context, UpdateProductOptionParam) (*domain.ProductOption, error)
	GetDeleteImageURL(context.Context, int) (string, error)
	GetUploadImageURL(context.Context) (*domain.ProductUploadURLImage, error)
	CreateProductImages(context.Context, []CreateProductImageParam) ([]domain.ProductImage, error)
}

type ProductImpl struct{}

func ProvideProduct() *ProductImpl {
	return &ProductImpl{}
}

var _ Product = &ProductImpl{}

func (s *ProductImpl) CreateProduct(ctx context.Context, param CreateProductParam) (*domain.Product, error) {
	return nil, nil
}

func (s *ProductImpl) UpdateProduct(ctx context.Context, param UpdateProductParam) (*domain.Product, error) {
	return nil, nil
}

func (s *ProductImpl) ListProducts(ctx context.Context, param ListProductParam) (*domain.DataPagination, error) {
	return nil, nil
}

func (s *ProductImpl) GetProduct(ctx context.Context, param GetProductParam) (*domain.Product, error) {
	return nil, nil
}

func (s *ProductImpl) DeleteProduct(ctx context.Context, param DeleteProductParam) error {
	return nil
}

func (s *ProductImpl) CreateProductOptions(ctx context.Context, param []CreateProductOptionParam) (*domain.ProductOption, error) {
	return nil, nil
}

func (s *ProductImpl) CreateProductVariants(ctx context.Context, param []CreateProductVariantParam) (*domain.ProductVariant, error) {
	return nil, nil
}

func (s *ProductImpl) UpdateProductVariants(ctx context.Context, param []UpdateProductVariantParam) (*domain.ProductVariant, error) {
	return nil, nil
}

func (s *ProductImpl) UpdateProductOption(ctx context.Context, param UpdateProductOptionParam) (*domain.ProductOption, error) {
	return nil, nil
}

func (s *ProductImpl) GetDeleteImageURL(ctx context.Context, imageID int) (string, error) {
	return "", nil
}

func (s *ProductImpl) GetUploadImageURL(ctx context.Context) (*domain.ProductUploadURLImage, error) {
	return nil, nil
}

func (s *ProductImpl) CreateProductImages(ctx context.Context, param []CreateProductImageParam) ([]domain.ProductImage, error) {
	return nil, nil
}

