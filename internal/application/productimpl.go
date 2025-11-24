package application

import (
	"context"
	"fmt"

	"backend/internal/domain"
	"backend/internal/helper/slice"

	"github.com/google/uuid"
)

type ProductImpl struct {
	attributeRepo        domain.AttributeRepository
	attributeService     domain.AttributeService
	categoryRepo         domain.CategoryRepository
	productCache         ProductCache
	productObjectStorage ProductObjectStorage
	productRepo          domain.ProductRepository
	productService       domain.ProductService
}

func ProvideProduct(
	attributeRepo domain.AttributeRepository,
	attributeService domain.AttributeService,
	categoryRepo domain.CategoryRepository,
	productCache ProductCache,
	productObjectStorage ProductObjectStorage,
	productRepo domain.ProductRepository,
	productService domain.ProductService,
) *ProductImpl {
	return &ProductImpl{
		attributeRepo:        attributeRepo,
		attributeService:     attributeService,
		categoryRepo:         categoryRepo,
		productCache:         productCache,
		productObjectStorage: productObjectStorage,
		productRepo:          productRepo,
		productService:       productService,
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
		(param.Page-1)*param.Limit,
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
	if cachedProduct, err := p.productCache.GetProduct(ctx, param.ProductID); err == nil {
		return cachedProduct, nil
	}
	// HACK: we assume that product repo also get the category :v
	product, err := p.productRepo.Get(ctx, param.ProductID)
	if err != nil {
		return nil, err
	}
	_ = p.productCache.SetProduct(ctx, param.ProductID, product)

	return product, nil
}

func (p *ProductImpl) Create(ctx context.Context, param CreateProductParam) (*domain.Product, error) {
	category, err := p.categoryRepo.Get(ctx, param.Data.CategoryID)
	if err != nil {
		return nil, err
	}
	product, err := domain.CreateProduct(
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
		attributeValues := domain.FilterAttributeValuesFromAttributes(*attributes, attributeValueIDs)
		product.AddAttributeValues(attributeValues...)
	}
	var options *[]domain.Option
	if param.Data.Options != nil {
		optionsWithOptionValues := make(map[string][]string, len(*param.Data.Options))
		for _, optionData := range *param.Data.Options {
			optionValues := make([]string, 0, len(optionData.Values))
			optionValues = append(optionValues, optionData.Values...)
			optionsWithOptionValues[optionData.Name] = optionValues
		}
		options, err = domain.CreateOptionsWithOptionValues(optionsWithOptionValues)
		if err != nil {
			return nil, err
		}
		product.AddOptions(*options...)
	}
	productImages := make([]domain.ProductImage, 0, len(param.Data.Images))
	for _, imgData := range param.Data.Images {
		image, err := domain.CreateImage(
			imgData.URL,
			imgData.Order,
		)
		if err != nil {
			return nil, err
		}
		productImages = append(productImages, *image)
	}
	product.AddImages(productImages...)
	for _, variantData := range param.Data.Variants {
		variant, err := domain.CreateVariant(
			variantData.SKU,
			variantData.Price,
			variantData.Quantity,
		)
		if err != nil {
			return nil, err
		}
		product.AddVariants(*variant)
		if variantData.Images != nil {
			variantImages := make([]domain.ProductImage, 0, len(*variantData.Images))
			for _, imgData := range *variantData.Images {
				image, err := domain.CreateImage(
					imgData.URL,
					imgData.Order,
				)
				if err != nil {
					return nil, err
				}
				variantImages = append(variantImages, *image)
			}
			if err := product.AddVariantImages(variant.ID, variantImages...); err != nil {
				return nil, err
			}
		}
		if options != nil {
			err = linkProductVariantsToOptionValues(product, *options, param)
			if err != nil {
				return nil, err
			}
		}
	}
	product.UpdateMinPrice()
	err = p.productService.Validate(*product)
	if err != nil {
		return nil, err
	}
	err = p.productRepo.Save(ctx, *product)
	if err != nil {
		return nil, err
	}
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
	product.Update(
		param.Data.Name,
		param.Data.Description,
		category,
	)
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
	product.Remove()
	err = p.productRepo.Save(ctx, *product)
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
		variant, err := domain.CreateVariant(
			variantData.SKU,
			variantData.Price,
			variantData.Quantity,
		)
		if err != nil {
			return nil, err
		}
		variants = append(variants, *variant)
	}
	product.AddVariants(variants...)
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
	if err := product.UpdateVariant(
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
	if err := p.productService.Validate(*product); err != nil {
		return nil, err
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
		image, err := domain.CreateImage(
			imgData.URL,
			imgData.Order,
		)
		if err != nil {
			return nil, err
		}
		if imgData.ProductVariantID != nil {
			if err := product.AddVariantImages(*imgData.ProductVariantID, *image); err != nil {
				return nil, err
			}
		} else {
			images = append(images, *image)
		}
	}
	product.AddImages(images...)
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
	for i := range product.Images {
		if _, exists := imageIDs[product.Images[i].ID]; exists {
			product.Images[i].Remove()
		}
	}
	for i := range product.Variants {
		for j := range product.Variants[i].Images {
			if _, exists := imageIDs[product.Variants[i].Images[j].ID]; exists {
				product.Variants[i].Images[j].Remove()
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
		if err := product.UpdateOption(
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
		if err := product.UpdateOptionValue(
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
	url, err := p.productObjectStorage.GetUploadImageURL(ctx)
	if err != nil {
		return nil, err
	}
	return url, nil
}

func (p *ProductImpl) GetDeleteImageURL(ctx context.Context, imageID uuid.UUID) (*DeleteImageURL, error) {
	url, err := p.productObjectStorage.GetDeleteImageURL(ctx, imageID)
	if err != nil {
		return nil, err
	}
	return url, nil
}

func linkProductVariantsToOptionValues(
	product *domain.Product,
	options []domain.Option,
	param CreateProductParam,
) error {
	optionsMap := make(map[string]domain.OptionValue, 0)
	for _, option := range options {
		for _, optionValue := range option.Values {
			optionsMap[fmt.Sprintf("%s/%s", option.Name, optionValue.Value)] = optionValue
		}
	}
	for i := range product.Variants {
		optionValues := make([]domain.OptionValue, len(*param.Data.Variants[i].Options))
		for j, optionData := range *param.Data.Variants[i].Options {
			optionValue, exists := optionsMap[fmt.Sprintf("%s/%s", optionData.Name, optionData.Value)]
			if !exists {
				return domain.ErrInvalid
			}
			optionValues[j] = optionValue
		}
		product.Variants[i].OptionValues = append(product.Variants[i].OptionValues, optionValues...)
	}
	return nil
}
