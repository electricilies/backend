package serviceimpl

import (
	"context"

	"backend/internal/domain"
)

type Product struct{}

func ProvideProduct() *Product {
	return &Product{}
}

var _ domain.ProductService = &Product{}

func (s *Product) Create(ctx context.Context, param domain.CreateProductParam) (*domain.Product, error) {
	panic("implement me")
}

func (s *Product) Update(ctx context.Context, param domain.UpdateProductParam) (*domain.Product, error) {
	panic("implement me")
}

func (s *Product) Get(ctx context.Context, param domain.GetProductParam) (*domain.Product, error) {
	panic("implement me")
}

func (s *Product) Delete(ctx context.Context, param domain.DeleteProductParam) error {
	panic("implement me")
}

func (s *Product) AddImages(ctx context.Context, param domain.AddProductImagesParam) (*[]domain.ProductImage, error) {
	panic("implement me")
}

func (s *Product) DeleteImages(ctx context.Context, param domain.DeleteProductImagesParam) error {
	panic("implement me")
}

func (s *Product) AddVariants(ctx context.Context, param domain.AddProductVariantsParam) (*domain.ProductVariant, error) {
	panic("implement me")
}

func (s *Product) UpdateVariant(ctx context.Context, param domain.UpdateProductVariantParam) (*domain.ProductVariant, error) {
	panic("implement me")
}

func (s *Product) UpdateOptions(ctx context.Context, param domain.UpdateProductOptionsParam) error {
	panic("implement me")
}

func (s *Product) UpdateOptionValues(ctx context.Context, param domain.UpdateProductOptionValuesParam) error {
	panic("implement me")
}
