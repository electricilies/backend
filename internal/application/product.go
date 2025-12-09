package application

import (
	"context"
	"fmt"

	"backend/config"
	"backend/internal/delivery/http"
	"backend/internal/domain"

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
	srvCfg               *config.Server
}

func ProvideProduct(
	attributeRepo domain.AttributeRepository,
	attributeService domain.AttributeService,
	categoryRepo domain.CategoryRepository,
	productCache ProductCache,
	productObjectStorage ProductObjectStorage,
	productRepo domain.ProductRepository,
	productService domain.ProductService,
	srvCfg *config.Server,
) *Product {
	return &Product{
		attributeRepo:        attributeRepo,
		attributeService:     attributeService,
		categoryRepo:         categoryRepo,
		productCache:         productCache,
		productObjectStorage: productObjectStorage,
		productRepo:          productRepo,
		productService:       productService,
		srvCfg:               srvCfg,
	}
}

var _ http.ProductApplication = (*Product)(nil)

func (p *Product) List(ctx context.Context, param http.ListProductRequestDto) (*http.PaginationResponseDto[http.ProductResponseDto], error) {
	cacheParam := ProductCacheListParam{
		IDs:         param.ProductIDs,
		Search:      param.Search,
		MinPrice:    param.MinPrice,
		MaxPrice:    param.MaxPrice,
		Rating:      param.Rating,
		CategoryIDs: param.CategoryIDs,
		Deleted:     param.Deleted,
		SortRating:  param.SortRating,
		SortPrice:   param.SortPrice,
		Limit:       param.Limit,
		Page:        param.Page,
	}

	if cachedPagination, err := p.productCache.GetList(ctx, cacheParam); err == nil {
		return cachedPagination, nil
	}

	products, err := p.productRepo.List(
		ctx,
		domain.ProductRepositoryListParam{
			IDs:         param.ProductIDs,
			Search:      param.Search,
			MinPrice:    param.MinPrice,
			MaxPrice:    param.MaxPrice,
			Rating:      param.Rating,
			CategoryIDs: param.CategoryIDs,
			Deleted:     param.Deleted,
			SortRating:  param.SortRating,
			SortPrice:   param.SortPrice,
			Limit:       param.Limit,
			Offset:      (param.Page - 1) * param.Limit,
		},
	)
	if err != nil {
		return nil, err
	}

	count, err := p.productRepo.Count(
		ctx,
		domain.ProductRepositoryCountParam{
			IDs:         param.ProductIDs,
			MinPrice:    param.MinPrice,
			MaxPrice:    param.MaxPrice,
			Rating:      param.Rating,
			CategoryIDs: param.CategoryIDs,
			Deleted:     param.Deleted,
		},
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
		domain.CategoryRepositoryListParam{
			IDs: categoryIDs,
		},
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
		domain.AttributeRepositoryListParam{
			AttributeValueIDs: attributeValuesIDs,
			Deleted:           domain.DeletedExcludeParam,
		},
	)
	if err != nil {
		return nil, err
	}

	productDtos := make([]http.ProductResponseDto, 0, len(*products))
	for _, product := range *products {
		dto := http.ToProductResponseDto(&product)
		if dto != nil {
			if category, exists := categoryMap[product.CategoryID]; exists {
				dto.WithCategory(category)
			}
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

	_ = p.productCache.SetList(ctx, cacheParam, pagination)

	return pagination, nil
}

func (p *Product) Get(ctx context.Context, param http.GetProductRequestDto) (*http.ProductResponseDto, error) {
	cacheParam := ProductCacheParam{ID: param.ProductID}

	if cachedProduct, err := p.productCache.Get(ctx, cacheParam); err == nil {
		return cachedProduct, nil
	}

	product, err := p.productRepo.Get(ctx, domain.ProductRepositoryGetParam{ProductID: param.ProductID})
	if err != nil {
		return nil, err
	}

	category, err := p.categoryRepo.Get(ctx, domain.CategoryRepositoryGetParam{ID: product.CategoryID})
	if err != nil {
		return nil, err
	}

	attributes, err := p.attributeRepo.List(
		ctx,
		domain.AttributeRepositoryListParam{
			AttributeValueIDs: product.AttributeValueIDs,
			Deleted:           domain.DeletedExcludeParam,
		},
	)
	if err != nil {
		return nil, err
	}

	productDto := http.ToProductResponseDto(product)
	productDto.WithCategory(category)
	productDto.WithAttributes(
		*attributes,
		product.AttributeValueIDs,
	)
	_ = p.productCache.Set(ctx, cacheParam, productDto)

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
	category, err := p.categoryRepo.Get(ctx, domain.CategoryRepositoryGetParam{ID: param.Data.CategoryID})
	if err != nil {
		return nil, err
	}
	if len(param.Data.AttributeValueIDs) > 0 {
		attributeIDs := make([]uuid.UUID, 0, len(param.Data.AttributeValueIDs))
		attributeValueIDs := make([]uuid.UUID, 0, len(param.Data.AttributeValueIDs))
		for _, a := range param.Data.AttributeValueIDs {
			attributeIDs = append(attributeIDs, a.AttributeID)
			attributeValueIDs = append(attributeValueIDs, a.ValueID)
		}

		// Validate that attributes exist
		attributes, err := p.attributeRepo.List(
			ctx,
			domain.AttributeRepositoryListParam{
				IDs:     attributeIDs,
				Deleted: domain.DeletedExcludeParam,
			},
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
	{
		optionsWithOptionValues := make(map[string][]string, len(param.Data.Options))
		for _, optionData := range param.Data.Options {
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
			imgData.Order,
			p.productObjectStorage.BuildImageURL,
		)
		if err != nil {
			return nil, err
		}
		productImages = append(productImages, *image)
		if err = p.productObjectStorage.PersistImageFromTemp(
			ctx,
			imgData.Key,
			image.ID,
		); err != nil {
			return nil, err
		}
	}
	product.AddImages(productImages...)
	imageKeyImageIDMap := make(map[string]uuid.UUID, len(param.Data.Images))
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
		variantImages := make([]domain.ProductImage, 0, len(variantData.Images))
		for _, imgData := range variantData.Images {
			image, err := domain.NewProductImage(
				imgData.Order,
				p.productObjectStorage.BuildImageURL,
			)
			if err != nil {
				return nil, err
			}
			variantImages = append(variantImages, *image)
			imageKeyImageIDMap[imgData.Key] = image.ID

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
	product.UpdateMinPrice()
	err = p.productService.Validate(*product)
	if err != nil {
		return nil, err
	}
	err = p.productRepo.Save(ctx, domain.ProductRepositorySaveParam{Product: *product})
	if err != nil {
		return nil, err
	}

	for key, imageID := range imageKeyImageIDMap {
		err = p.productObjectStorage.PersistImageFromTemp(
			ctx,
			key,
			imageID,
		)
		if err != nil {
			return nil, err
		}
	}

	_ = p.productCache.InvalidateAlls(ctx)

	// Fetch the created product with all relations
	var attributes *[]domain.Attribute
	if len(product.AttributeIDs) > 0 {
		attributes, err = p.attributeRepo.List(
			ctx,
			domain.AttributeRepositoryListParam{
				IDs:     product.AttributeIDs,
				Deleted: domain.DeletedExcludeParam,
			},
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
	product, err := p.productRepo.Get(ctx, domain.ProductRepositoryGetParam{ProductID: param.ProductID})
	if err != nil {
		return nil, err
	}

	categoryID := product.CategoryID
	if param.Data.CategoryID != uuid.Nil {
		categoryID = param.Data.CategoryID
	}
	category, err := p.categoryRepo.Get(ctx, domain.CategoryRepositoryGetParam{
		ID: categoryID,
	})
	if err != nil {
		return nil, err
	}
	err = p.productService.Validate(*product)
	if err != nil {
		return nil, err
	}
	product.Update(
		param.Data.Name,
		param.Data.Description,
		category.ID,
	)
	err = p.productRepo.Save(ctx, domain.ProductRepositorySaveParam{Product: *product})
	if err != nil {
		return nil, err
	}
	_ = p.productCache.InvalidateAlls(ctx)

	// Fetch attributes if any
	var attributes *[]domain.Attribute
	if len(product.AttributeIDs) > 0 {
		attributes, err = p.attributeRepo.List(
			ctx,
			domain.AttributeRepositoryListParam{
				IDs:     product.AttributeIDs,
				Deleted: domain.DeletedExcludeParam,
			},
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
	product, err := p.productRepo.Get(ctx, domain.ProductRepositoryGetParam{ProductID: param.ProductID})
	if err != nil {
		return err
	}
	err = p.productService.Validate(*product)
	if err != nil {
		return err
	}
	product.Remove()
	err = p.productRepo.Save(ctx, domain.ProductRepositorySaveParam{Product: *product})
	if err != nil {
		return err
	}
	_ = p.productCache.InvalidateAlls(ctx)
	return nil
}

func (p *Product) AddVariants(ctx context.Context, param http.AddProductVariantsRequestDto) (*[]http.ProductVariantResponseDto, error) {
	product, err := p.productRepo.Get(ctx, domain.ProductRepositoryGetParam{ProductID: param.ProductID})
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
	err = p.productRepo.Save(ctx, domain.ProductRepositorySaveParam{Product: *product})
	if err != nil {
		return nil, err
	}
	_ = p.productCache.InvalidateAlls(ctx)
	variantDtos := http.ToProductVariantResponseDtoList(variants)
	return &variantDtos, nil
}

func (p *Product) UpdateVariant(ctx context.Context, param http.UpdateProductVariantRequestDto) (*http.ProductVariantResponseDto, error) {
	product, err := p.productRepo.Get(ctx, domain.ProductRepositoryGetParam{ProductID: param.ProductID})
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
	err = p.productRepo.Save(ctx, domain.ProductRepositorySaveParam{Product: *product})
	if err != nil {
		return nil, err
	}
	// Invalidate product list caches
	_ = p.productCache.InvalidateAlls(ctx)
	return http.ToProductVariantResponseDto(variant), nil
}

func (p *Product) AddImages(ctx context.Context, param http.AddProductImagesRequestDto) (*[]http.ProductImageResponseDto, error) {
	product, err := p.productRepo.Get(ctx, domain.ProductRepositoryGetParam{ProductID: param.ProductID})
	if err != nil {
		return nil, err
	}
	images := make([]domain.ProductImage, 0, len(param.Data))
	imageKeyImageIDMap := make(map[string]uuid.UUID, len(param.Data))
	for _, imgData := range param.Data {
		image, err := domain.NewProductImage(
			imgData.Order,
			p.productObjectStorage.BuildImageURL,
		)
		if err != nil {
			return nil, err
		}
		if imgData.ProductVariantID != uuid.Nil {
			if err := product.AddVariantImages(imgData.ProductVariantID, *image); err != nil {
				return nil, err
			}
		} else {
			images = append(images, *image)
			imageKeyImageIDMap[imgData.Key] = image.ID
		}
	}
	product.AddImages(images...)
	err = p.productRepo.Save(ctx, domain.ProductRepositorySaveParam{Product: *product})
	if err != nil {
		return nil, err
	}
	for key, imageID := range imageKeyImageIDMap {
		err = p.productObjectStorage.PersistImageFromTemp(
			ctx,
			key,
			imageID,
		)
		if err != nil {
			return nil, err
		}
	}
	// Invalidate product caches
	_ = p.productCache.InvalidateAlls(ctx)
	imageDtos := http.ToProductImageResponseDtoList(images)
	return &imageDtos, nil
}

func (p *Product) DeleteImages(ctx context.Context, param http.DeleteProductImagesRequestDto) error {
	product, err := p.productRepo.Get(ctx, domain.ProductRepositoryGetParam{ProductID: param.ProductID})
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
	err = p.productRepo.Save(ctx, domain.ProductRepositorySaveParam{Product: *product})
	if err != nil {
		return err
	}
	_ = p.productCache.InvalidateAlls(ctx)
	return nil
}

func (p *Product) UpdateOptions(ctx context.Context, param http.UpdateProductOptionsRequestDto) (*[]http.ProductOptionResponseDto, error) {
	product, err := p.productRepo.Get(ctx, domain.ProductRepositoryGetParam{ProductID: param.ProductID})
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
	err = p.productRepo.Save(ctx, domain.ProductRepositorySaveParam{Product: *product})
	if err != nil {
		return nil, err
	}
	_ = p.productCache.InvalidateAlls(ctx)
	optionDtos := http.ToProductOptionResponseDtoList(options)
	return &optionDtos, nil
}

func (p *Product) UpdateOptionValues(ctx context.Context, param http.UpdateProductOptionValuesRequestDto) (*[]http.ProductOptionValueResponseDto, error) {
	product, err := p.productRepo.Get(ctx, domain.ProductRepositoryGetParam{ProductID: param.ProductID})
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
	err = p.productRepo.Save(ctx, domain.ProductRepositorySaveParam{Product: *product})
	if err != nil {
		return nil, err
	}
	_ = p.productCache.InvalidateAlls(ctx)
	optionValueDtos := http.ToProductOptionValueResponseDtoList(optionValues)
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
		optionValues := make([]domain.OptionValue, len(param.Data.Variants[i].Options))
		for j, optionData := range param.Data.Variants[i].Options {
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
