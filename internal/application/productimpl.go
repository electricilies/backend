package application

import (
	"context"

	"backend/internal/domain"
	"backend/internal/helper/slice"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

type ProductImpl struct {
	productRepo      domain.ProductRepository
	productService   domain.ProductService
	categoryRepo     domain.CategoryRepository
	attributeRepo    domain.AttributeRepository
	attributeService domain.AttributeService
	productCache     ProductCache
	s3Client         *s3.Client
}

func ProvideProduct(
	productRepo domain.ProductRepository,
	productService domain.ProductService,
	productCache ProductCache,
	s3Client *s3.Client,
) *ProductImpl {
	return &ProductImpl{
		productRepo:    productRepo,
		productService: productService,
		productCache:   productCache,
		s3Client:       s3Client,
	}
}

var _ Product = (*ProductImpl)(nil)

func (p *ProductImpl) List(ctx context.Context, param ListProductParam) (*Pagination[domain.Product], error) {
	cacheKey := p.productCache.BuildListCacheKey(
		param.ProductIDs,
		param.Search,
		param.MinPrice,
		param.MaxPrice,
		param.Rating,
		param.CategoryIDs,
		param.Deleted,
		param.SortRating,
		param.SortPrice,
		param.Limit,
		param.Page,
	)

	// Try to get from cache
	if cachedPagination, err := p.productCache.GetProductList(ctx, cacheKey); err == nil {
		return cachedPagination, nil
	}

	products, err := p.productRepo.List(
		ctx,
		param.ProductIDs,
		param.Search,
		param.MinPrice,
		param.MaxPrice,
		param.Rating,
		param.CategoryIDs,
		param.Deleted,
		param.SortRating,
		param.SortPrice,
		param.Limit,
		param.Page*param.Limit,
	)
	if err != nil {
		return nil, err
	}

	count, err := p.productRepo.Count(
		ctx,
		param.ProductIDs,
		param.MinPrice,
		param.MaxPrice,
		param.Rating,
		param.CategoryIDs,
		param.Deleted,
	)
	if err != nil {
		return nil, err
	}

	pagination := newPagination(
		*products,
		*count,
		param.Page,
		param.Limit,
	)

	// Cache the result
	_ = p.productCache.SetProductList(ctx, cacheKey, pagination)

	return pagination, nil
}

func (p *ProductImpl) Get(ctx context.Context, param GetProductParam) (*domain.Product, error) {
	// Try to get from cache
	if cachedProduct, err := p.productCache.GetProduct(ctx, param.ProductID); err == nil {
		return cachedProduct, nil
	}

	product, err := p.productRepo.Get(ctx, param.ProductID)
	if err != nil {
		return nil, err
	}

	// Cache the result
	_ = p.productCache.SetProduct(ctx, param.ProductID, product)

	return product, nil
}

func (p *ProductImpl) Create(ctx context.Context, param CreateProductParam) (*domain.Product, error) {
	category, err := p.categoryRepo.Get(ctx, param.Data.CategoryID)
	if err != nil {
		return nil, err
	}
	product, err := p.productService.Create(
		param.Data.Name,
		param.Data.Description,
		*category,
	)
	if err != nil {
		return nil, err
	}
	if param.Data.AttributeValueIDs != nil {
		attributeIDs := make([]uuid.UUID, 0, len(*param.Data.AttributeValueIDs))
		attributeValueIDs := make([]uuid.UUID, 0, len(*param.Data.AttributeValueIDs))
		for _, a := range *param.Data.AttributeValueIDs {
			attributeIDs = append(attributeIDs, a.AttributeID)
			attributeValueIDs = append(attributeValueIDs, a.ValueID)
		}
		attributes, err := p.attributeRepo.List(
			ctx,
			&attributeIDs,
			nil,
			domain.DeletedExcludeParam,
			0, 0,
		)
		if err != nil {
			return nil, err
		}
		attributeValues := p.attributeService.FilterAttributeValuesFromAttributes(*attributes, attributeValueIDs)
		err = p.productService.AddAttributeValues(product, attributeValues...)
		if err != nil {
			return nil, err
		}
	}
	optionsWithOptionValues := make(map[string][]string, len(param.Data.Options))
	options, err := p.productService.CreateOptionsWithOptionValues(optionsWithOptionValues)
	if err != nil {
		return nil, err
	}
	if err := p.productService.AddOptions(product, *options...); err != nil {
		return nil, err
	}
	productImages := make([]domain.ProductImage, 0, len(param.Data.Images))
	for _, imgData := range param.Data.Images {
		image, err := p.productService.CreateImage(
			imgData.URL,
			imgData.Order,
		)
		if err != nil {
			return nil, err
		}
		productImages = append(productImages, *image)
	}
	if err := p.productService.AddImages(product, productImages...); err != nil {
		return nil, err
	}
	productVariants := make([]domain.ProductVariant, 0, len(param.Data.Variants))
	for _, variantData := range param.Data.Variants {
		variant, err := p.productService.CreateVariant(
			variantData.SKU,
			variantData.Price,
			variantData.Quantity,
		)
		if err != nil {
			return nil, err
		}
		productVariants = append(productVariants, *variant)
		if variantData.Images != nil {
			variantImages := make([]domain.ProductImage, 0, len(*variantData.Images))
			for _, imgData := range *variantData.Images {
				image, err := p.productService.CreateImage(
					imgData.URL,
					imgData.Order,
				)
				if err != nil {
					return nil, err
				}
				variantImages = append(variantImages, *image)
			}
			if err := p.productService.AddVariantImages(product, variant.ID, variantImages...); err != nil {
				return nil, err
			}
		}
	}
	if err := p.productService.AddVariants(product, productVariants...); err != nil {
		return nil, err
	}
	err = p.productRepo.Save(ctx, *product)
	if err != nil {
		return nil, err
	}
	// Invalidate cache after create
	_ = p.productCache.InvalidateProductList(ctx)
	return product, nil
}

func (p *ProductImpl) Update(ctx context.Context, param UpdateProductParam) (*domain.Product, error) {
	var category *domain.Category
	if param.Data.CategoryID != nil {
		cat, err := p.categoryRepo.Get(ctx, *param.Data.CategoryID)
		if err != nil {
			return nil, err
		}
		category = cat
	}
	product, err := p.productRepo.Get(ctx, param.ProductID)
	if err != nil {
		return nil, err
	}
	if err := p.productService.Update(
		product,
		param.Data.Name,
		param.Data.Description,
		category,
	); err != nil {
		return nil, err
	}
	err = p.productRepo.Save(ctx, *product)
	if err != nil {
		return nil, err
	}
	_ = p.productCache.InvalidateProduct(ctx, param.ProductID)
	_ = p.productCache.InvalidateProductList(ctx)
	return product, nil
}

func (p *ProductImpl) Delete(ctx context.Context, param DeleteProductParam) error {
	product, err := p.productRepo.Get(ctx, param.ProductID)
	if err != nil {
		return err
	}
	err = p.productService.Remove(product)
	if err != nil {
		return err
	}
	_ = p.productCache.InvalidateProduct(ctx, param.ProductID)
	_ = p.productCache.InvalidateProductList(ctx)
	return nil
}

func (p *ProductImpl) AddVariants(ctx context.Context, param AddProductVariantsParam) (*[]domain.ProductVariant, error) {
	product, err := p.productRepo.Get(ctx, param.ProductID)
	if err != nil {
		return nil, err
	}
	variants := make([]domain.ProductVariant, 0, len(param.Data))
	for _, variantData := range param.Data {
		variant, err := p.productService.CreateVariant(
			variantData.SKU,
			variantData.Price,
			variantData.Quantity,
		)
		if err != nil {
			return nil, err
		}
		variants = append(variants, *variant)
	}
	if err := p.productService.AddVariants(product, variants...); err != nil {
		return nil, err
	}
	err = p.productRepo.Save(ctx, *product)
	if err != nil {
		return nil, err
	}
	_ = p.productCache.InvalidateProduct(ctx, param.ProductID)
	_ = p.productCache.InvalidateProductList(ctx)
	return &variants, nil
}

func (p *ProductImpl) UpdateVariant(ctx context.Context, param UpdateProductVariantParam) (*domain.ProductVariant, error) {
	product, err := p.productRepo.Get(ctx, param.ProductID)
	if err != nil {
		return nil, err
	}
	if err := p.productService.UpdateVariant(
		product,
		param.ProductVariantID,
		param.Data.Price,
		param.Data.Quantity,
	); err != nil {
		return nil, err
	}
	variant := product.GetVariantByID(param.ProductVariantID)
	if variant == nil {
		return nil, domain.ErrNotFound
	}
	err = p.productRepo.Save(ctx, *product)
	if err != nil {
		return nil, err
	}
	// Invalidate product list caches
	_ = p.productCache.InvalidateProductList(ctx)
	return variant, nil
}

func (p *ProductImpl) AddImages(ctx context.Context, param AddProductImagesParam) (*[]domain.ProductImage, error) {
	product, err := p.productRepo.Get(ctx, param.ProductID)
	if err != nil {
		return nil, err
	}
	images := make([]domain.ProductImage, 0, len(param.Data))
	for _, imgData := range param.Data {
		image, err := p.productService.CreateImage(
			imgData.URL,
			imgData.Order,
		)
		if err != nil {
			return nil, err
		}
		if imgData.ProductVariantID != nil {
			if err := p.productService.AddVariantImages(product, *imgData.ProductVariantID, *image); err != nil {
				return nil, err
			}
		} else {
			images = append(images, *image)
		}
	}
	if err := p.productService.AddImages(product, images...); err != nil {
		return nil, err
	}
	err = p.productRepo.Save(ctx, *product)
	if err != nil {
		return nil, err
	}
	// Invalidate product caches
	_ = p.productCache.InvalidateAllProducts(ctx)
	return &images, nil
}

func (p *ProductImpl) DeleteImages(ctx context.Context, param DeleteProductImagesParam) error {
	product, err := p.productRepo.Get(ctx, param.ProductID)
	if err != nil {
		return err
	}
	imageIDs := make(map[uuid.UUID]struct{}, len(param.ImageIDs))
	for _, id := range param.ImageIDs {
		imageIDs[id] = struct{}{}
	}
	for _, image := range product.Images {
		if _, exists := imageIDs[image.ID]; exists {
			err := p.productService.RemoveImage(&image)
			if err != nil {
				return err
			}
		}
	}
	for _, variant := range product.Variants {
		for _, image := range variant.Images {
			if _, exists := imageIDs[image.ID]; exists {
				err := p.productService.RemoveImage(&image)
				if err != nil {
					return err
				}
			}
		}
	}
	err = p.productRepo.Save(ctx, *product)
	if err != nil {
		return err
	}
	_ = p.productCache.InvalidateAllProducts(ctx)
	return nil
}

func (p *ProductImpl) UpdateOptions(ctx context.Context, param UpdateProductOptionsParam) (*[]domain.Option, error) {
	product, err := p.productRepo.Get(ctx, param.ProductID)
	if err != nil {
		return nil, err
	}
	optionIDs := make([]uuid.UUID, 0, len(param.Data))
	for _, data := range param.Data {
		if err := p.productService.UpdateOption(
			product,
			data.ID,
			data.Name,
		); err != nil {
			return nil, err
		}
		optionIDs = append(optionIDs, data.ID)
	}
	options := product.GetOptionsByIDs(optionIDs)
	if options == nil {
		return nil, domain.ErrNotFound
	}
	err = p.productRepo.Save(ctx, *product)
	if err != nil {
		return nil, err
	}
	_ = p.productCache.InvalidateAllProducts(ctx)
	return slice.SlicePtrToPtrSlice(options), nil
}

func (p *ProductImpl) UpdateOptionValues(ctx context.Context, param UpdateProductOptionValuesParam) (*[]domain.OptionValue, error) {
	product, err := p.productRepo.Get(ctx, param.ProductID)
	if err != nil {
		return nil, err
	}
	optionValueIDs := make([]uuid.UUID, 0, len(param.Data))
	for _, data := range param.Data {
		if err := p.productService.UpdateOptionValue(
			product,
			param.OptionID,
			data.ID,
			data.Value,
		); err != nil {
			return nil, err
		}
	}
	option := product.GetOptionByID(param.OptionID)
	if option == nil {
		return nil, domain.ErrNotFound
	}
	optionValues := option.GetValuesByIDs(optionValueIDs)
	err = p.productRepo.Save(ctx, *product)
	if err != nil {
		return nil, err
	}
	_ = p.productCache.InvalidateAllProducts(ctx)
	return slice.SlicePtrToPtrSlice(optionValues), nil
}

func (p *ProductImpl) GetUploadImageURL(ctx context.Context) (*UploadImageURL, error) {
	return nil, domain.ErrNotImplemented
}

func (p *ProductImpl) GetDeleteImageURL(ctx context.Context, imageID int) (*DeleteImageURL, error) {
	// TODO: Implement get delete image URL logic
	return nil, domain.ErrNotImplemented
}
