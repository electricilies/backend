package product

import (
	"backend/internal/domain/category"
	"backend/internal/domain/param"
	"backend/internal/domain/product"
	"backend/internal/helper"
	"backend/internal/infrastructure/mapper"
	"backend/internal/infrastructure/persistence/postgres"

	"github.com/jackc/pgx/v5/pgtype"
)

func ToGetProductImageParam(id int) *postgres.GetProductImageParams {
	return &postgres.GetProductImageParams{
		ID: int32(id),
	}
}

func ToListProductsParam(productQueryParams product.QueryParams) *postgres.ListProductsParams {
	return &postgres.ListProductsParams{
		Search:      mapper.ToPostgresTextType(productQueryParams.Search),
		IDs:         mapper.ToInt32Slice(productQueryParams.IDs),
		MinPrice:    mapper.ToPostgresTypeNumeric(productQueryParams.MinPrice),
		MaxPrice:    mapper.ToPostgresTypeNumeric(productQueryParams.MaxPrice),
		Rating:      mapper.ToPostgresTypeFloat(productQueryParams.Rating),
		CategoryIDs: mapper.ToInt32Slice(productQueryParams.CategoryIDs),
		Deleted:     string(productQueryParams.Deleted),
		SortRating:  mapper.ToPostgresTextType(helper.ToPtr(string(productQueryParams.SortRating))),
		SortPrice:   mapper.ToPostgresTextType(helper.ToPtr(string(productQueryParams.SortPrice))),
		Limit: mapper.ToPostgresTypeInt(
			&productQueryParams.PaginationParams.Limit,
		),
		Offset: mapper.ToPostgresTypeInt(
			&productQueryParams.PaginationParams.Offset,
		),
	}
}

func ToCreateProductParams(model product.Model) *postgres.CreateProductParams {
	return &postgres.CreateProductParams{
		Name:        *model.Name,
		Description: *model.Description,
	}
}

func ToUpdateProductParams(model product.Model, id int) *postgres.UpdateProductParams {
	return &postgres.UpdateProductParams{
		ID:          int32(id),
		Name:        mapper.ToPostgresTextType(model.Name),
		Description: mapper.ToPostgresTextType(model.Description),
		CategoryID:  mapper.ToPostgresTypeInt(&model.Category.ID),
	}
}

func ToDeleteProductsParam(ids []int) *postgres.DeleteProductsParams {
	return &postgres.DeleteProductsParams{
		IDs: mapper.ToInt32Slice(&ids),
	}
}

func ToCreateOptionParams(
	optionModel product.OptionModel,
	id int,
) *postgres.CreateOptionParams {
	return &postgres.CreateOptionParams{
		Name:      optionModel.Name,
		ProductID: int32(id),
	}
}

func ToCreateProductImagesParams(
	imageModel []product.ImageModel,
	productID int,
) *postgres.CreateProductImagesParams {
	length := len(imageModel)
	URLs := make([]string, 0, length)
	Orders := make([]int32, 0, length)
	VariantIDs := make([]int32, 0, length)
	for _, image := range imageModel {
		URLs = append(URLs, *image.URL)
		Orders = append(Orders, int32(*image.Order))
		VariantIDs = append(
			VariantIDs,
			int32(*image.ProductVariantID),
		)
	}
	return &postgres.CreateProductImagesParams{
		URLs:              URLs,
		Orders:            Orders,
		ProductVariantIDs: VariantIDs,
		ProductID:         mapper.ToPostgresTypeInt(&productID),
	}
}

func ToCreateProductVariantParams(
	variantModel []product.VariantModel,
	productId int,
) *postgres.CreateProductVariantsParams {
	length := len(variantModel)
	SKUs := make([]string, 0, length)
	Prices := make([]pgtype.Numeric, 0, length)
	Quantities := make([]int32, 0, length)
	for _, variant := range variantModel {
		SKUs = append(SKUs, variant.SKU)
		Prices = append(
			Prices,
			mapper.ToPostgresTypeNumeric(variant.Price),
		)
		Quantities = append(
			Quantities,
			int32(*variant.Quantity),
		)
	}
	return &postgres.CreateProductVariantsParams{
		SKUs:       SKUs,
		Prices:     Prices,
		Quantities: Quantities,
		ProductID:  int32(productId),
	}
}

func ImageToDomain(productImageEntity postgres.ProductImage) *product.ImageModel {
	return &product.ImageModel{
		ID:               helper.ToPtr(int(productImageEntity.ID)),
		URL:              helper.ToPtr(productImageEntity.URL),
		Order:            helper.ToPtr(int(productImageEntity.Order)),
		CreatedAt:        &productImageEntity.CreatedAt.Time,
		ProductVariantID: helper.ToPtr(int(productImageEntity.ProductVariantID.Int32)),
	}
}

func VariantToDomain(productVariantEntity postgres.ProductVariant) *product.VariantModel {
	return &product.VariantModel{
		ID:            int(productVariantEntity.ID),
		SKU:           productVariantEntity.SKU,
		Price:         helper.ToPtr(productVariantEntity.Price.Int.Int64()),
		Quantity:      helper.ToPtr(int(productVariantEntity.Quantity)),
		PurchaseCount: int(productVariantEntity.PurchaseCount),
		CreatedAt:     productVariantEntity.CreatedAt.Time,
		UpdatedAt:     productVariantEntity.UpdatedAt.Time,
		DeletedAt:     &productVariantEntity.DeletedAt.Time,
	}
}

func ToDomain(productEntity postgres.Product) *product.Model {
	return &product.Model{
		ID:            int(productEntity.ID),
		Name:          &productEntity.Name,
		Description:   &productEntity.Description,
		Price:         productEntity.Price.Int.Int64(),
		ViewsCount:    int(productEntity.ViewsCount),
		TotalPurchase: int(productEntity.TotalPurchase),
		Rating:        productEntity.Rating,
		TrendingScore: productEntity.TrendingScore,
		Category: &category.Model{
			ID: int(productEntity.CategoryID), // TODO: add mapper later
		},
		CreatedAt: productEntity.CreatedAt.Time,
		UpdatedAt: productEntity.UpdatedAt.Time,
		DeletedAt: &productEntity.DeletedAt.Time,
	}
}

func OptionToDomain(productEntity postgres.Option) *product.OptionModel {
	return &product.OptionModel{
		ID:   int(productEntity.ID),
		Name: productEntity.Name,
	}
}

func PaginationMetadataToDomain(
	paginationParams param.Pagination,
	currentCount int,
	totalCount int,
) *param.PaginationMetadata {
	currentPage := int(paginationParams.Offset/paginationParams.Limit) + 1
	return &param.PaginationMetadata{
		TotalRecords: totalCount,
		CurrentPage:  currentPage,
		ItemsPerPage: paginationParams.Limit,
		PageItems:    currentCount,
	}
}

func ListProductRowsToDomain(
	listProductRow []postgres.ListProductsRow,
	paginationParams param.Pagination,
) *product.PaginationModel {
	length := len(listProductRow)
	if length == 0 {
		emptyProducts := make([]product.Model, 0)
		return &product.PaginationModel{
			Products: emptyProducts,
			Metadata: *PaginationMetadataToDomain(
				paginationParams,
				0,
				0,
			),
		}
	}
	domainProducts := make([]product.Model, 0, len(listProductRow))
	for _, row := range listProductRow {
		domainProducts = append(domainProducts, *ToDomain(row.Product))
	}
	return &product.PaginationModel{
		Products: domainProducts,
		Metadata: *PaginationMetadataToDomain(
			paginationParams,
			int((listProductRow)[0].CurrentCount),
			int((listProductRow)[0].TotalCount),
		),
	}
}

type UploadURLImage struct {
	URL string
	Key string
}

func (p UploadURLImage) ToDomain() *product.UploadImageURLModel {
	return &product.UploadImageURLModel{
		URL: p.URL,
		Key: p.Key,
	}
}
