package repositorypostgres

import (
	"context"
	"math/big"
	"time"

	"backend/internal/domain"
	"backend/internal/helper/ptr"
	"backend/internal/infrastructure/repositorypostgres/sqlc"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Product struct {
	queries *sqlc.Queries
	conn    *pgxpool.Pool
}

var _ domain.ProductRepository = (*Product)(nil)

func ProvideProduct(
	q *sqlc.Queries,
	conn *pgxpool.Pool,
) *Product {
	return &Product{
		queries: q,
		conn:    conn,
	}
}

func (r *Product) List(
	ctx context.Context,
	ids *[]uuid.UUID,
	search *string,
	minPrice *int64,
	maxPrice *int64,
	rating *float64,
	categoryIDs *[]uuid.UUID,
	deleted domain.DeletedParam,
	sortRating *string,
	sortPrice *string,
	limit int,
	offset int,
) (*[]domain.Product, error) {
	minPriceParam := pgtype.Numeric{
		Int:   big.NewInt(ptr.Deref(minPrice, 0)),
		Valid: minPrice != nil,
	}
	maxPriceParam := pgtype.Numeric{
		Int:   big.NewInt(ptr.Deref(maxPrice, 0)),
		Valid: maxPrice != nil,
	}
	var ratingParam *float32
	if rating != nil {
		r := float32(*rating)
		ratingParam = &r
	}
	productEntities, err := r.queries.ListProducts(ctx, sqlc.ListProductsParams{
		IDs:         ptr.Deref(ids, []uuid.UUID{}),
		Search:      search,
		MinPrice:    minPriceParam,
		MaxPrice:    maxPriceParam,
		Rating:      ratingParam,
		CategoryIDs: ptr.Deref(categoryIDs, []uuid.UUID{}),
		Deleted:     string(deleted),
		SortRating:  sortRating,
		SortPrice:   sortPrice,
		Limit:       int32(limit),
		Offset:      int32(offset),
	})
	if err != nil {
		return nil, ToDomainErrorFromPostgres(err)
	}
	products := make([]domain.Product, 0, len(productEntities))
	for _, productEntity := range productEntities {
		product, err := r.Get(ctx, productEntity.ID)
		if err != nil {
			return nil, ToDomainErrorFromPostgres(err)
		}
		products = append(products, *product)
	}
	return &products, nil
}

func (r *Product) Count(
	ctx context.Context,
	ids *[]uuid.UUID,
	minPrice *int64,
	maxPrice *int64,
	rating *float64,
	categoryIDs *[]uuid.UUID,
	deleted domain.DeletedParam,
) (*int, error) {
	minPriceParam := pgtype.Numeric{
		Int:   big.NewInt(ptr.Deref(minPrice, 0)),
		Valid: minPrice != nil,
	}
	maxPriceParam := pgtype.Numeric{
		Int:   big.NewInt(ptr.Deref(maxPrice, 0)),
		Valid: maxPrice != nil,
	}
	var ratingParam *float32
	if rating != nil {
		r := float32(*rating)
		ratingParam = &r
	}
	productEntities, err := r.queries.CountProducts(ctx, sqlc.CountProductsParams{
		IDs:         ptr.Deref(ids, []uuid.UUID{}),
		MinPrice:    minPriceParam,
		MaxPrice:    maxPriceParam,
		Rating:      ratingParam,
		CategoryIDs: ptr.Deref(categoryIDs, []uuid.UUID{}),
		Deleted:     string(deleted),
	})
	if err != nil {
		return nil, ToDomainErrorFromPostgres(err)
	}
	return ptr.To(int(productEntities)), nil
}

// FIXME: Missing deleted
// HACK: This should take only product
// Category and attribute need to take from those domains
func (r *Product) Get(ctx context.Context, productID uuid.UUID) (*domain.Product, error) {
	productEntity, err := r.queries.GetProduct(ctx, sqlc.GetProductParams{
		ID: productID,
	})
	if err != nil {
		return nil, ToDomainErrorFromPostgres(err)
	}
	product := &domain.Product{
		ID:            productEntity.ID,
		Name:          productEntity.Name,
		Description:   productEntity.Description,
		ViewsCount:    int(productEntity.ViewsCount),
		TotalPurchase: int(productEntity.TotalPurchase),
		TrendingScore: int64(productEntity.TrendingScore),
		Price:         productEntity.Price.Int.Int64(),
		Rating:        float64(productEntity.Rating),
		CreatedAt:     productEntity.CreatedAt.Time,
		UpdatedAt:     productEntity.UpdatedAt.Time,
		DeletedAt:     fromPgValidToPtr(productEntity.DeletedAt.Time, productEntity.DeletedAt.Valid),
	}
	if err := getCategory(ctx, *r.queries, product, productEntity); err != nil {
		return nil, ToDomainErrorFromPostgres(err)
	}
	if err := GetAttributeAndValue(ctx, *r.queries, product); err != nil {
		return nil, ToDomainErrorFromPostgres(err)
	}
	if err := getOptionsAndValues(ctx, *r.queries, product); err != nil {
		return nil, ToDomainErrorFromPostgres(err)
	}
	if err := getVariants(ctx, *r.queries, product); err != nil {
		return nil, ToDomainErrorFromPostgres(err)
	}
	if err := getImages(ctx, *r.queries, product); err != nil {
		return nil, ToDomainErrorFromPostgres(err)
	}
	return product, nil
}

func getCategory(
	ctx context.Context,
	queries sqlc.Queries,
	product *domain.Product,
	productEntity sqlc.Product,
) error {
	if product == nil {
		return nil
	}
	categoryEntity, err := queries.GetCategory(ctx, sqlc.GetCategoryParams{
		ID: productEntity.CategoryID,
	})
	if err != nil {
		return err
	}
	category := domain.Category{
		ID:        categoryEntity.ID,
		Name:      categoryEntity.Name,
		CreatedAt: categoryEntity.CreatedAt.Time,
		UpdatedAt: categoryEntity.UpdatedAt.Time,
		DeletedAt: fromPgValidToPtr(categoryEntity.DeletedAt.Time, categoryEntity.DeletedAt.Valid),
	}
	product.Category = &category
	return nil
}

func GetAttributeAndValue(
	ctx context.Context,
	queries sqlc.Queries,
	product *domain.Product,
) error {
	productsAttributeValuesEntity, err := queries.ListProductsAttributeValues(ctx, sqlc.ListProductsAttributeValuesParams{
		ProductID: pgtype.UUID{
			Bytes: product.ID,
			Valid: true,
		},
	})
	if err != nil {
		return err
	}
	attributeValueIDs := make([]uuid.UUID, 0, len(productsAttributeValuesEntity))
	for _, pav := range productsAttributeValuesEntity {
		attributeValueIDs = append(attributeValueIDs, pav.AttributeValueID)
	}
	attributeValueEntities, err := queries.ListAttributeValues(ctx, sqlc.ListAttributeValuesParams{
		IDs: attributeValueIDs,
	})
	if err != nil {
		return err
	}
	attributeIDs := make([]uuid.UUID, 0, len(attributeValueEntities))
	for _, av := range attributeValueEntities {
		attributeIDs = append(attributeIDs, av.AttributeID)
	}
	attributeEntities, err := queries.ListAttributes(ctx, sqlc.ListAttributesParams{
		IDs: attributeIDs,
	})
	if err != nil {
		return err
	}
	attributes := make([]domain.Attribute, 0, len(attributeEntities))
	for _, attr := range attributeEntities {
		attributes = append(attributes, domain.Attribute{
			ID:        attr.ID,
			Name:      attr.Name,
			Code:      attr.Code,
			DeletedAt: fromPgValidToPtr(attr.DeletedAt.Time, attr.DeletedAt.Valid),
		})
	}
	attributeMap := make(map[uuid.UUID]domain.Attribute)
	for i, attr := range attributes {
		attributeMap[attr.ID] = attributes[i]
	}
	attributeValues := make([]domain.AttributeValue, 0, len(attributeValueEntities))
	for _, av := range attributeValueEntities {
		attributeValues = append(attributeValues, domain.AttributeValue{
			ID:        av.ID,
			DeletedAt: fromPgValidToPtr(av.DeletedAt.Time, av.DeletedAt.Valid),
			Value:     av.Value,
		})
	}
	product.AttributeValues = attributeValues
	return nil
}

func getOptionsAndValues(ctx context.Context,
	queries sqlc.Queries,
	product *domain.Product,
) error {
	optionEntities, err := queries.ListOptions(ctx, sqlc.ListOptionsParams{
		ProductID: pgtype.UUID{
			Bytes: product.ID,
			Valid: true,
		},
	})
	if err != nil {
		return err
	}
	optionIDs := make([]uuid.UUID, 0, len(optionEntities))
	for _, option := range optionEntities {
		optionIDs = append(optionIDs, option.ID)
	}
	optionValueEntities, err := queries.ListOptionValues(ctx, sqlc.ListOptionValuesParams{
		OptionIds: optionIDs,
	})
	if err != nil {
		return err
	}
	optionValueMap := make(map[uuid.UUID][]domain.OptionValue)
	for _, ov := range optionValueEntities {
		optionValueMap[ov.OptionID] = append(optionValueMap[ov.OptionID], domain.OptionValue{
			ID:        ov.ID,
			Value:     ov.Value,
			DeletedAt: fromPgValidToPtr(ov.DeletedAt.Time, ov.DeletedAt.Valid),
		})
	}
	options := make([]domain.Option, 0, len(optionEntities))
	for _, option := range optionEntities {
		options = append(options, domain.Option{
			ID:        option.ID,
			Name:      option.Name,
			Values:    optionValueMap[option.ID],
			DeletedAt: fromPgValidToPtr(option.DeletedAt.Time, option.DeletedAt.Valid),
		})
	}
	product.Options = options
	return nil
}

func getVariants(
	ctx context.Context,
	queries sqlc.Queries,
	product *domain.Product,
) error {
	variantEntities, err := queries.ListProductVariants(ctx, sqlc.ListProductVariantsParams{
		ProductID: pgtype.UUID{
			Bytes: product.ID,
			Valid: true,
		},
	})
	if err != nil {
		return err
	}
	variants := make([]domain.ProductVariant, 0, len(variantEntities))
	for _, variant := range variantEntities {
		variants = append(variants, domain.ProductVariant{
			ID:            variant.ID,
			SKU:           variant.SKU,
			Price:         variant.Price.Int.Int64(),
			Quantity:      int(variant.Quantity),
			PurchaseCount: int(variant.PurchaseCount),
			CreatedAt:     variant.CreatedAt.Time,
			UpdatedAt:     variant.UpdatedAt.Time,
			DeletedAt:     fromPgValidToPtr(variant.DeletedAt.Time, variant.DeletedAt.Valid),
		})
	}
	product.Variants = variants
	productVariantIDs := make([]uuid.UUID, 0, len(variantEntities))
	for _, v := range variantEntities {
		productVariantIDs = append(productVariantIDs, v.ID)
	}
	OptionValuesProductVariantsEntities, err := queries.ListOptionValuesProductVariants(ctx, sqlc.ListOptionValuesProductVariantsParams{
		ProductVariantIDs: productVariantIDs,
	})
	if err != nil {
		return err
	}
	variantIDOptionValueIDsMap := make(map[uuid.UUID][]uuid.UUID)
	for _, ovpve := range OptionValuesProductVariantsEntities {
		variantIDOptionValueIDsMap[ovpve.ProductVariantID] = append(variantIDOptionValueIDsMap[ovpve.ProductVariantID], ovpve.OptionValueID)
	}
	optionValueLength := 0
	for _, option := range product.Options {
		optionValueLength += len(option.Values)
	}
	optionValueIDOptionValueMap := make(map[uuid.UUID]domain.OptionValue, optionValueLength)
	for _, option := range product.Options {
		for _, ov := range option.Values {
			optionValueIDOptionValueMap[ov.ID] = ov
		}
	}
	variantIDOptionValuesMap := make(map[uuid.UUID][]domain.OptionValue, len(variantEntities))
	for variantID, optionValueIDs := range variantIDOptionValueIDsMap {
		for _, ovID := range optionValueIDs {
			if ov, exists := optionValueIDOptionValueMap[ovID]; exists {
				variantIDOptionValuesMap[variantID] = append(variantIDOptionValuesMap[variantID], ov)
			}
		}
	}
	for i, variant := range variants {
		if ovs, exists := variantIDOptionValuesMap[variant.ID]; exists {
			variants[i].OptionValues = ovs
		}
	}
	product.Variants = variants
	return nil
}

func getImages(
	ctx context.Context,
	queries sqlc.Queries,
	product *domain.Product,
) error {
	variantLength := len(product.Variants)
	variantIDs := make([]uuid.UUID, 0, variantLength)
	for _, variant := range product.Variants {
		variantIDs = append(variantIDs, variant.ID)
	}
	imagesEntities, err := queries.ListProductImages(ctx, sqlc.ListProductImagesParams{
		ProductVariantIDs: variantIDs,
	})
	if err != nil {
		return err
	}
	variantIDvariantMap := make(map[uuid.UUID]*domain.ProductVariant, variantLength)
	for i, variant := range product.Variants {
		variantIDvariantMap[variant.ID] = &product.Variants[i]
	}
	for _, imgEntity := range imagesEntities {
		img := domain.ProductImage{
			ID:        imgEntity.ID,
			URL:       imgEntity.URL,
			Order:     int(imgEntity.Order),
			CreatedAt: imgEntity.CreatedAt.Time,
			DeletedAt: fromPgValidToPtr(
				imgEntity.DeletedAt.Time,
				imgEntity.DeletedAt.Valid,
			),
		}
		if !imgEntity.ProductVariantID.Valid {
			product.Images = append(product.Images, img)
		} else {
			if variant, exists := variantIDvariantMap[imgEntity.ProductVariantID.Bytes]; exists {
				variant.Images = append(variant.Images, img)
			}
		}
	}
	return nil
}

func (r *Product) Save(ctx context.Context, product domain.Product) error {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return ToDomainErrorFromPostgres(err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	qtx := r.queries.WithTx(tx)
	if err := upsertProduct(ctx, *qtx, product); err != nil {
		return ToDomainErrorFromPostgres(err)
	}
	if err := mergeAttributeValues(ctx, *qtx, product); err != nil {
		return ToDomainErrorFromPostgres(err)
	}
	if err := mergeOptions(ctx, *qtx, product); err != nil {
		return ToDomainErrorFromPostgres(err)
	}
	if err := mergeOptionValues(ctx, *qtx, product); err != nil {
		return ToDomainErrorFromPostgres(err)
	}
	if err := mergeVariants(ctx, *qtx, product); err != nil {
		return ToDomainErrorFromPostgres(err)
	}
	if err := mergeOptionValuesProductVariants(ctx, *qtx, product); err != nil {
		return ToDomainErrorFromPostgres(err)
	}
	if err := mergeImages(ctx, *qtx, product); err != nil {
		return ToDomainErrorFromPostgres(err)
	}
	err = tx.Commit(ctx)
	if err != nil {
		return ToDomainErrorFromPostgres(err)
	}
	return nil
}

func upsertProduct(
	ctx context.Context,
	qtx sqlc.Queries,
	product domain.Product,
) error {
	return qtx.UpsertProduct(ctx, sqlc.UpsertProductParams{
		ID:            product.ID,
		Name:          product.Name,
		Description:   product.Description,
		ViewsCount:    int32(product.ViewsCount),
		TotalPurchase: int32(product.TotalPurchase),
		TrendingScore: float32(product.TrendingScore),
		Price: pgtype.Numeric{
			Int:   big.NewInt(product.Price),
			Valid: true,
		},
		Rating: float32(product.Rating),
		CreatedAt: pgtype.Timestamptz{
			Time:  product.CreatedAt,
			Valid: true,
		},
		UpdatedAt: pgtype.Timestamptz{
			Time:  product.UpdatedAt,
			Valid: true,
		},
		DeletedAt: pgtype.Timestamptz{
			Time:  ptr.Deref(product.DeletedAt, time.Time{}),
			Valid: product.DeletedAt != nil,
		},
		CategoryID: product.Category.ID,
	})
}

func mergeAttributeValues(
	ctx context.Context,
	qtx sqlc.Queries,
	product domain.Product,
) error {
	if err := qtx.CreateTempTableProductsAttributeValues(ctx); err != nil {
		return err
	}
	param := make([]sqlc.InsertTempTableProductsAttributeValuesParams, 0, len(product.AttributeValues))
	for _, pav := range product.AttributeValues {
		param = append(param, sqlc.InsertTempTableProductsAttributeValuesParams{
			ProductID:        product.ID,
			AttributeValueID: pav.ID,
		})
	}
	_, err := qtx.InsertTempTableProductsAttributeValues(ctx, param)
	if err != nil {
		return err
	}
	return qtx.MergeProductsAttributeValuesFromTemp(ctx)
}

func mergeOptions(
	ctx context.Context,
	qtx sqlc.Queries,
	product domain.Product,
) error {
	if err := qtx.CreateTempTableOptions(ctx); err != nil {
		return err
	}
	param := make([]sqlc.InsertTempTableOptionsParams, 0, len(product.Options))
	for _, option := range product.Options {
		param = append(param, sqlc.InsertTempTableOptionsParams{
			ProductID: product.ID,
			ID:        option.ID,
			Name:      option.Name,
			DeletedAt: pgtype.Timestamptz{
				Time:  ptr.Deref(option.DeletedAt, time.Time{}),
				Valid: option.DeletedAt != nil,
			},
		})
	}
	_, err := qtx.InsertTempTableOptions(ctx, param)
	if err != nil {
		return err
	}
	return qtx.MergeOptionsFromTemp(ctx)
}

func mergeOptionValues(
	ctx context.Context,
	qtx sqlc.Queries,
	product domain.Product,
) error {
	if err := qtx.CreateTempTableOptionValues(ctx); err != nil {
		return err
	}
	param := make([]sqlc.InsertTempTableOptionValuesParams, 0)
	for _, option := range product.Options {
		for _, ov := range option.Values {
			param = append(param, sqlc.InsertTempTableOptionValuesParams{
				OptionID: option.ID,
				ID:       ov.ID,
				Value:    ov.Value,
				DeletedAt: pgtype.Timestamptz{
					Time:  ptr.Deref(ov.DeletedAt, time.Time{}),
					Valid: ov.DeletedAt != nil,
				},
			})
		}
	}
	_, err := qtx.InsertTempTableOptionValues(ctx, param)
	if err != nil {
		return err
	}
	return qtx.MergeOptionValuesFromTemp(ctx)
}

func mergeVariants(
	ctx context.Context,
	qtx sqlc.Queries,
	product domain.Product,
) error {
	if err := qtx.CreateTempTableProductVariants(ctx); err != nil {
		return err
	}
	param := make([]sqlc.InsertTempTableProductVariantsParams, 0, len(product.Variants))
	for _, variant := range product.Variants {
		param = append(param, sqlc.InsertTempTableProductVariantsParams{
			ProductID: product.ID,
			ID:        variant.ID,
			SKU:       variant.SKU,
			Price: pgtype.Numeric{
				Int:   big.NewInt(variant.Price),
				Valid: true,
			},
			Quantity: int32(variant.Quantity),
			CreatedAt: pgtype.Timestamptz{
				Time:  variant.CreatedAt,
				Valid: true,
			},
			UpdatedAt: pgtype.Timestamptz{
				Time:  variant.UpdatedAt,
				Valid: true,
			},
			PurchaseCount: int32(variant.PurchaseCount),
			DeletedAt: pgtype.Timestamptz{
				Time:  ptr.Deref(variant.DeletedAt, time.Time{}),
				Valid: variant.DeletedAt != nil,
			},
		})
	}
	_, err := qtx.InsertTempTableProductVariants(ctx, param)
	if err != nil {
		return err
	}
	return qtx.MergeProductVariantsFromTemp(ctx)
}

func mergeOptionValuesProductVariants(
	ctx context.Context,
	qtx sqlc.Queries,
	product domain.Product,
) error {
	if err := qtx.CreateTempTableOptionValuesProductVariants(ctx); err != nil {
		return err
	}
	length := 0
	for _, variant := range product.Variants {
		length += len(variant.OptionValues)
	}
	param := make([]sqlc.InsertTempTableOptionValuesProductVariantsParams, 0, length)
	for _, variant := range product.Variants {
		for _, ov := range variant.OptionValues {
			param = append(param, sqlc.InsertTempTableOptionValuesProductVariantsParams{
				ProductVariantID: variant.ID,
				OptionValueID:    ov.ID,
			})
		}
	}
	_, err := qtx.InsertTempTableOptionValuesProductVariants(ctx, param)
	if err != nil {
		return err
	}
	return qtx.MergeOptionValuesProductVariantsFromTemp(ctx)
}

func mergeImages(
	ctx context.Context,
	qtx sqlc.Queries,
	product domain.Product,
) error {
	if err := qtx.CreateTempTableProductImages(ctx); err != nil {
		return err
	}
	length := len(product.Images)
	for _, variant := range product.Variants {
		length += len(variant.Images)
	}
	param := make([]sqlc.InsertTempTableProductImagesParams, 0, length)
	for _, img := range product.Images {
		param = append(param, sqlc.InsertTempTableProductImagesParams{
			ProductID: product.ID,
			ID:        img.ID,
			URL:       img.URL,
			Order:     int32(img.Order),
			CreatedAt: pgtype.Timestamptz{
				Time:  product.CreatedAt,
				Valid: true,
			},
			DeletedAt: pgtype.Timestamptz{
				Time:  ptr.Deref(img.DeletedAt, time.Time{}),
				Valid: img.DeletedAt != nil,
			},
		})
	}
	for _, variant := range product.Variants {
		for _, img := range variant.Images {
			param = append(param, sqlc.InsertTempTableProductImagesParams{
				ProductID: product.ID,
				ProductVariantID: pgtype.UUID{
					Bytes: variant.ID,
					Valid: true,
				},
				ID:    img.ID,
				URL:   img.URL,
				Order: int32(img.Order),
				CreatedAt: pgtype.Timestamptz{
					Time:  variant.CreatedAt,
					Valid: true,
				},
				DeletedAt: pgtype.Timestamptz{
					Time:  ptr.Deref(img.DeletedAt, time.Time{}),
					Valid: img.DeletedAt != nil,
				},
			})
		}
	}
	_, err := qtx.InsertTempTableProductImages(ctx, param)
	if err != nil {
		return err
	}
	return qtx.MergeProductImagesFromTemp(ctx)
}
