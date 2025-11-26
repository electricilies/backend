package application

import (
	"context"
	"fmt"

	"backend/internal/delivery/http"
	"backend/internal/domain"
	"backend/internal/helper/ptr"
	"backend/internal/helper/slice"

	"github.com/google/uuid"
)

type Product struct {
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
) *Product {
	return &Product{
		attributeRepo:        attributeRepo,
		attributeService:     attributeService,
		categoryRepo:         categoryRepo,
		productCache:         productCache,
		productObjectStorage: productObjectStorage,
		productRepo:          productRepo,
		productService:       productService,
	}
}

var _ http.ProductApplication = (*Product)(nil)

func (p *Product) List(ctx context.Context, param http.ListProductRequestDto) (*http.PaginationResponseDto[http.ProductResponseDto], error) {
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
		nil,
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

	categoryIDs := make([]uuid.UUID, 0, len(*products))
	for _, product := range *products {
		categoryIDs = append(categoryIDs, product.CategoryID)
	}

	categories, err := p.categoryRepo.List(
		ctx,
		&categoryIDs,
		nil,
		0, 0,
	)
	if err != nil {
		return nil, err
	}

	categoryMap := make(map[uuid.UUID]*domain.Category)
	for i := range *categories {
		categoryMap[(*categories)[i].ID] = &(*categories)[i]
	}

	attributeValuesIDs := make([]uuid.UUID, 0, len(*products))
	for _, product := range *products {
		attributeValuesIDs = append(attributeValuesIDs, product.AttributeValueIDs...)
	}

	attributes, err := p.attributeRepo.List(
		ctx,
		nil,
		&attributeValuesIDs,
		nil,
		domain.DeletedExcludeParam,
		0, 0,
	)
	if err != nil {
		return nil, err
	}

	// Map domain models to response DTOs
	productDtos := make([]http.ProductResponseDto, 0, len(*products))
	for _, product := range *products {
		dto := http.ToProductResponseDto(&product)
		if dto != nil {
			// Merge category
			if category, exists := categoryMap[product.CategoryID]; exists {
				dto.WithCategory(category)
			}
			// Merge attributes
			if attributes != nil {
				dto.WithAttributes(*attributes, product.AttributeValueIDs)
			}
			productDtos = append(productDtos, *dto)
		}
	}

	pagination := newPaginationResponseDto(
		productDtos,
		*count,
		param.Page,
		param.Limit,
	)

	// Cache the result
	_ = p.productCache.SetProductList(ctx, cacheKey, pagination)

	return pagination, nil
}

func (p *Product) Get(ctx context.Context, param http.GetProductRequestDto) (*http.ProductResponseDto, error) {
	if cachedProduct, err := p.productCache.GetProduct(ctx, param.ProductID); err == nil {
		return cachedProduct, nil
	}
	product, err := p.productRepo.Get(ctx, param.ProductID)
	if err != nil {
		return nil, err
	}

	category, err := p.categoryRepo.Get(ctx, product.CategoryID)
	if err != nil {
		return nil, err
	}

	attributes, err := p.attributeRepo.List(
		ctx,
		nil,
		&product.AttributeValueIDs,
		nil,
		domain.DeletedExcludeParam,
		0, 0,
	)
	if err != nil {
		return nil, err
	}

	productDto := http.ToProductResponseDto(product)
	productDto.WithCategory(category)
	if attributes != nil {
		productDto.WithAttributes(
			*attributes,
			product.AttributeValueIDs,
		)
	}
	_ = p.productCache.SetProduct(ctx, param.ProductID, productDto)

	return productDto, nil
}

func (p *Product) Create(ctx context.Context, param http.CreateProductRequestDto) (*http.ProductResponseDto, error) {
	product, err := domain.NewProduct(
		param.Data.Name,
		param.Data.Description,
		param.Data.CategoryID,
	)
	if err != nil {
		return nil, err
	}

	// Get and validate category
	category, err := p.categoryRepo.Get(ctx, param.Data.CategoryID)
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

		// Validate that attributes exist
		attributes, err := p.attributeRepo.List(
			ctx,
			&attributeIDs,
			nil,
			nil,
			domain.DeletedExcludeParam,
			0, 0,
		)
		if err != nil {
			return nil, err
		}
		if len(*attributes) != len(attributeIDs) {
			return nil, domain.ErrNotFound
		}

		product.AddAttributeIDs(attributeIDs...)
		product.AddAttributeValueIDs(attributeValueIDs...)
	}
	var options *[]domain.Option
	if param.Data.Options != nil {
		optionsWithOptionValues := make(map[string][]string, len(*param.Data.Options))
		for _, optionData := range *param.Data.Options {
			optionValues := make([]string, 0, len(optionData.Values))
			optionValues = append(optionValues, optionData.Values...)
			optionsWithOptionValues[optionData.Name] = optionValues
		}
		options, err = p.productService.CreateOptionsWithOptionValues(optionsWithOptionValues)
		if err != nil {
			return nil, err
		}
		product.AddOptions(*options...)
	}
	productImages := make([]domain.ProductImage, 0, len(param.Data.Images))
	for _, imgData := range param.Data.Images {
		image, err := domain.NewProductImage(
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
		variant, err := domain.NewVariant(
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
				image, err := domain.NewProductImage(
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

	// Fetch the created product with all relations
	var attributes *[]domain.Attribute
	if len(product.AttributeIDs) > 0 {
		attributes, err = p.attributeRepo.List(
			ctx,
			&product.AttributeIDs,
			nil,
			nil,
			domain.DeletedExcludeParam,
			0, 0,
		)
		if err != nil {
			return nil, err
		}
	}

	productDto := http.ToProductResponseDto(product)
	productDto.WithCategory(category)
	if attributes != nil {
		productDto.WithAttributes(*attributes, product.AttributeValueIDs)
	}
	return productDto, nil
}

func (p *Product) Update(ctx context.Context, param http.UpdateProductRequestDto) (*http.ProductResponseDto, error) {
	product, err := p.productRepo.Get(ctx, param.ProductID)
	if err != nil {
		return nil, err
	}

	var category *domain.Category
	if param.Data.CategoryID != nil {
		cat, err := p.categoryRepo.Get(ctx, *param.Data.CategoryID)
		if err != nil {
			return nil, err
		}
		category = cat
	} else {
		// Get existing category
		cat, err := p.categoryRepo.Get(ctx, product.CategoryID)
		if err != nil {
			return nil, err
		}
		category = cat
	}
	err = p.productService.Validate(*product)
	if err != nil {
		return nil, err
	}
	product.Update(
		param.Data.Name,
		param.Data.Description,
		ptr.To(category.ID),
	)
	err = p.productRepo.Save(ctx, *product)
	if err != nil {
		return nil, err
	}
	_ = p.productCache.InvalidateProduct(ctx, param.ProductID)
	_ = p.productCache.InvalidateProductList(ctx)

	// Fetch attributes if any
	var attributes *[]domain.Attribute
	if len(product.AttributeIDs) > 0 {
		attributes, err = p.attributeRepo.List(
			ctx,
			&product.AttributeIDs,
			nil,
			nil,
			domain.DeletedExcludeParam,
			0, 0,
		)
		if err != nil {
			return nil, err
		}
	}

	productDto := http.ToProductResponseDto(product)
	productDto.WithCategory(category)
	if attributes != nil {
		productDto.WithAttributes(*attributes, product.AttributeValueIDs)
	}
	return productDto, nil
}

func (p *Product) Delete(ctx context.Context, param http.DeleteProductRequestDto) error {
	product, err := p.productRepo.Get(ctx, param.ProductID)
	if err != nil {
		return err
	}
	err = p.productService.Validate(*product)
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

func (p *Product) AddVariants(ctx context.Context, param http.AddProductVariantsRequestDto) (*[]http.ProductVariantResponseDto, error) {
	product, err := p.productRepo.Get(ctx, param.ProductID)
	if err != nil {
		return nil, err
	}
	variants := make([]domain.ProductVariant, 0, len(param.Data))
	for _, variantData := range param.Data {
		variant, err := domain.NewVariant(
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
	variantDtos := http.ToProductVariantResponseDtoList(variants)
	return &variantDtos, nil
}

func (p *Product) UpdateVariant(ctx context.Context, param http.UpdateProductVariantRequestDto) (*http.ProductVariantResponseDto, error) {
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
	return http.ToProductVariantResponseDto(variant), nil
}

func (p *Product) AddImages(ctx context.Context, param http.AddProductImagesRequestDto) (*[]http.ProductImageResponseDto, error) {
	product, err := p.productRepo.Get(ctx, param.ProductID)
	if err != nil {
		return nil, err
	}
	images := make([]domain.ProductImage, 0, len(param.Data))
	for _, imgData := range param.Data {
		image, err := domain.NewProductImage(
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
	imageDtos := http.ToProductImageResponseDtoList(images)
	return &imageDtos, nil
}

func (p *Product) DeleteImages(ctx context.Context, param http.DeleteProductImagesRequestDto) error {
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

func (p *Product) UpdateOptions(ctx context.Context, param http.UpdateProductOptionsRequestDto) (*[]http.ProductOptionResponseDto, error) {
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
	optionDtos := http.ToProductOptionResponseDtoList(*slice.SlicePtrToPtrSlice(options))
	return &optionDtos, nil
}

func (p *Product) UpdateOptionValues(ctx context.Context, param http.UpdateProductOptionValuesRequestDto) (*[]http.ProductOptionValueResponseDto, error) {
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
	optionValueDtos := http.ToProductOptionValueResponseDtoList(*slice.SlicePtrToPtrSlice(optionValues))
	return &optionValueDtos, nil
}

func (p *Product) GetUploadImageURL(ctx context.Context) (*http.UploadImageURLResponseDto, error) {
	url, err := p.productObjectStorage.GetUploadImageURL(ctx)
	if err != nil {
		return nil, err
	}
	return url, nil
}

func (p *Product) GetDeleteImageURL(ctx context.Context, imageID uuid.UUID) (*http.DeleteImageURLResponseDto, error) {
	url, err := p.productObjectStorage.GetDeleteImageURL(ctx, imageID)
	if err != nil {
		return nil, err
	}
	return url, nil
}

func linkProductVariantsToOptionValues(
	product *domain.Product,
	options []domain.Option,
	param http.CreateProductRequestDto,
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
