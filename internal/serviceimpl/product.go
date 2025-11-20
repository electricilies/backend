package serviceimpl

import "backend/internal/domain"

type Product struct{}

func ProvideProduct() *Product {
	return &Product{}
}

var _ domain.ProductService = &Product{}

func (p *Product) Create() {

// 	Create(context.Context, CreateProductParam) (*Product, error)
// 	Update(context.Context, UpdateProductParam) (*Product, error)
// 	Get(context.Context, GetProductParam) (*Product, error)
// 	Delete(context.Context, DeleteProductParam) error
// 	AddImages(context.Context, AddProductImagesParam) ([]ProductImage, error)
// 	DeleteImages(context.Context, DeleteProductImagesParam) error
// 	AddVariants(context.Context, AddProductVariantsParam) (*ProductVariant, error)
// 	UpdateVariant(context.Context, UpdateProductVariantParam) (*ProductVariant, error)
// 	UpdateOptions(context.Context, UpdateProductOptionsParam) error
// 	UpdateOptionValues(context.Context, UpdateProductOptionValuesParam) error
// }
