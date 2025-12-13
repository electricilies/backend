// vim: tabstop=4 shiftwidth=4:
//go:build integration

package application_test

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"backend/config"
	"backend/internal/application"
	"backend/internal/client"
	http_dto "backend/internal/delivery/http"
	"backend/internal/domain"
	"backend/internal/infrastructure/cacheredis"
	"backend/internal/infrastructure/objectstorages3"
	"backend/internal/infrastructure/repositorypostgres"
	"backend/internal/service"
	"backend/test/integration/component"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type ProductWithOptionsLifecycleTestSuite struct {
	suite.Suite
	containers *component.Containers
	app        http_dto.ProductApplication

	// For tracking created resources
	firstProductID uuid.UUID
}

func TestProductWithOptionsLifecycleSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(ProductWithOptionsLifecycleTestSuite))
}

func (s *ProductWithOptionsLifecycleTestSuite) newContainersConfig() *component.ContainersConfig {
	containersConfig := component.NewContainersConfig(&component.NewContainersConfigParam{
		DBEnabled:    true,
		RedisEnabled: true,
		MinIOEnabled: true,
	})
	containersConfig.DB.Seed = true
	return containersConfig
}

func (s *ProductWithOptionsLifecycleTestSuite) newConfig(
	ctx context.Context,
) *config.Server {
	s.T().Helper()

	dbConnStr, err := s.containers.DB.ConnectionString(ctx, "sslmode=disable")
	s.Require().NoError(err, "failed to get db connection string")

	redisConnStr, err := s.containers.Redis.ConnectionString(ctx)
	s.Require().NoError(err, "failed to get redis connection string")

	minioConnStr, err := s.containers.MinIO.ConnectionString(ctx)
	s.Require().NoError(err, "failed to get minio connection string")

	return &config.Server{
		DBURL:        dbConnStr,
		RedisAddr:    strings.TrimPrefix(redisConnStr, "redis://"),
		S3Endpoint:   "http://" + minioConnStr,
		S3Bucket:     "electricilies",
		S3AccessKey:  "electricilies",
		S3SecretKey:  "electricilies",
		S3RegionName: "us-east-1",
	}
}

func (s *ProductWithOptionsLifecycleTestSuite) SetupTest() {
	ctx := s.T().Context()
	containersConfig := s.newContainersConfig()

	var err error
	s.containers, err = component.NewContainers(ctx, containersConfig)
	s.Require().NoError(err, "failed to start containers")

	cfg := s.newConfig(ctx)

	validate := validator.New(
		validator.WithRequiredStructEnabled(),
	)
	err = domain.RegisterProductValidates(validate)
	s.Require().NoError(err)

	conn := client.NewDBConnection(ctx, cfg)
	queries := client.NewDBQueries(conn)

	productRepo := repositorypostgres.ProvideProduct(queries, conn)
	categoryRepo := repositorypostgres.ProvideCategory(queries)
	attributeRepo := repositorypostgres.ProvideAttribute(queries, conn)

	productService := service.ProvideProduct(validate)
	attributeService := service.ProvideAttribute(validate)

	redisClient := client.NewRedis(ctx, cfg)
	productCache := cacheredis.ProvideProduct(redisClient)

	s3Client := client.NewS3(ctx, cfg)
	s3PresignClient := client.NewS3Presign(s3Client)
	s3ClientWrapper := client.ProvideS3(s3Client, s3PresignClient)

	// Create bucket if it doesn't exist
	_, err = s3Client.CreateBucket(ctx, &s3.CreateBucketInput{
		Bucket: &cfg.S3Bucket,
	})
	if err != nil {
		// Ignore error if bucket already exists
		_ = err
	}

	productObjectStorage := objectstorages3.ProvideProduct(s3ClientWrapper, cfg)

	s.app = application.ProvideProduct(
		attributeRepo,
		attributeService,
		categoryRepo,
		productCache,
		productObjectStorage,
		productRepo,
		productService,
		cfg,
	)
}

func (s *ProductWithOptionsLifecycleTestSuite) TearDownTest() {
	s.containers.Cleanup(s.T())
}

func (s *ProductWithOptionsLifecycleTestSuite) uploadDummyImage(url string) {
	s.T().Helper()
	req, err := http.NewRequest("PUT", url, strings.NewReader("dummy image data"))
	s.Require().NoError(err)
	req.Header.Set("Content-Type", "image/jpeg")

	resp, err := http.DefaultClient.Do(req)
	s.Require().NoError(err)
	defer func() { _ = resp.Body.Close() }()
	s.Require().Equal(http.StatusOK, resp.StatusCode)
}

func (s *ProductWithOptionsLifecycleTestSuite) TestProductWithOptionsLifecycle() {
	ctx := s.T().Context()

	// Seeded category ID from seed data
	seededCategoryID := uuid.MustParse("00000000-0000-7000-0000-000000001796")
	seededAttributeID := uuid.MustParse("00000000-0000-7000-0000-000000000003")
	seededAttributeValueID := uuid.MustParse("00000000-0000-7000-0000-000000000104")

	s.Run("Create configurable product with options (partial matrix)", func() {
		// Get upload URLs for 3 product images
		uploadURL1, err := s.app.GetUploadImageURL(ctx)
		s.Require().NoError(err)
		s.uploadDummyImage(uploadURL1.URL)

		uploadURL2, err := s.app.GetUploadImageURL(ctx)
		s.Require().NoError(err)
		s.uploadDummyImage(uploadURL2.URL)

		uploadURL3, err := s.app.GetUploadImageURL(ctx)
		s.Require().NoError(err)
		s.uploadDummyImage(uploadURL3.URL)

		// Get upload URLs for variant images
		variantImgURL1, err := s.app.GetUploadImageURL(ctx)
		s.Require().NoError(err)
		s.uploadDummyImage(variantImgURL1.URL)

		variantImgURL2, err := s.app.GetUploadImageURL(ctx)
		s.Require().NoError(err)
		s.uploadDummyImage(variantImgURL2.URL)

		// Create product with 2 options: Size (S, M, L) and Color (Red, Blue)
		// Create 4 variants (partial matrix): S+Red, M+Blue, L+Red, L+Blue
		result, err := s.app.Create(ctx, http_dto.CreateProductRequestDto{
			Data: http_dto.CreateProductData{
				Name:        "Configurable Test Product",
				Description: "This is a configurable test product with options",
				CategoryID:  seededCategoryID,
				AttributeValueIDs: []http_dto.CreateProductAttributesData{
					{
						AttributeID: seededAttributeID,
						ValueID:     seededAttributeValueID,
					},
				},
				Options: []http_dto.CreateProductOptionData{
					{
						Name:   "Size",
						Values: []string{"S", "M", "L"},
					},
					{
						Name:   "Color",
						Values: []string{"Red", "Blue"},
					},
				},
				Images: []http_dto.CreateProductImageData{
					{Key: uploadURL1.Key, Order: 1},
					{Key: uploadURL2.Key, Order: 2},
					{Key: uploadURL3.Key, Order: 3},
				},
				Variants: []http_dto.CreateProductVariantData{
					{
						SKU:      "CONFIG-TEST-S-RED",
						Price:    100000,
						Quantity: 50,
						Options: []http_dto.CreateProductVariantOption{
							{Name: "Size", Value: "S"},
							{Name: "Color", Value: "Red"},
						},
						Images: []http_dto.CreateProductVariantImage{
							{Key: variantImgURL1.Key, Order: 1},
						},
					},
					{
						SKU:      "CONFIG-TEST-M-BLUE",
						Price:    120000,
						Quantity: 75,
						Options: []http_dto.CreateProductVariantOption{
							{Name: "Size", Value: "M"},
							{Name: "Color", Value: "Blue"},
						},
					},
					{
						SKU:      "CONFIG-TEST-L-RED",
						Price:    110000,
						Quantity: 60,
						Options: []http_dto.CreateProductVariantOption{
							{Name: "Size", Value: "L"},
							{Name: "Color", Value: "Red"},
						},
						Images: []http_dto.CreateProductVariantImage{
							{Key: variantImgURL2.Key, Order: 1},
						},
					},
					{
						SKU:      "CONFIG-TEST-L-BLUE",
						Price:    115000,
						Quantity: 80,
						Options: []http_dto.CreateProductVariantOption{
							{Name: "Size", Value: "L"},
							{Name: "Color", Value: "Blue"},
						},
					},
				},
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal("Configurable Test Product", result.Name)
		s.Len(result.Options, 2, "Should have 2 options")
		s.Len(result.Variants, 4, "Should have 4 variants (partial matrix)")
		s.InDelta(float64(100000), result.Price, 0.001, "Min price should be 100000")

		// Verify options and option values
		var sizeOption, colorOption *http_dto.ProductOptionResponseDto
		for i := range result.Options {
			switch result.Options[i].Name {
			case "Size":
				sizeOption = &result.Options[i]
			case "Color":
				colorOption = &result.Options[i]
			}
		}
		s.Require().NotNil(sizeOption, "Size option should exist")
		s.Require().NotNil(colorOption, "Color option should exist")
		s.Len(sizeOption.Values, 3, "Size should have 3 values")
		s.Len(colorOption.Values, 2, "Color should have 2 values")

		// Verify variants have correct option values
		for _, variant := range result.Variants {
			s.Len(variant.OptionValues, 2, "Each variant should have 2 option values")
		}

		s.firstProductID = result.ID
	})

	s.Run("Get product with options", func() {
		result, err := s.app.Get(ctx, http_dto.GetProductRequestDto{
			ProductID: s.firstProductID,
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal("Configurable Test Product", result.Name)
		s.Len(result.Options, 2)
		s.Len(result.Variants, 4)
	})

	s.Run("Update option name", func() {
		// Get product to get option IDs
		product, err := s.app.Get(ctx, http_dto.GetProductRequestDto{
			ProductID: s.firstProductID,
		})
		s.Require().NoError(err)

		var sizeOptionID uuid.UUID
		for _, opt := range product.Options {
			if opt.Name == "Size" {
				sizeOptionID = opt.ID
				break
			}
		}
		s.Require().NotEqual(uuid.Nil, sizeOptionID)

		result, err := s.app.UpdateOptions(ctx, http_dto.UpdateProductOptionsRequestDto{
			ProductID: s.firstProductID,
			Data: []http_dto.UpdateProductOptionsData{
				{
					ID:   sizeOptionID,
					Name: "Kích thước",
				},
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Len(*result, 1)
		s.Equal("Kích thước", (*result)[0].Name)
	})

	s.Run("Update option values", func() {
		// Get product to get option IDs and value IDs
		product, err := s.app.Get(ctx, http_dto.GetProductRequestDto{
			ProductID: s.firstProductID,
		})
		s.Require().NoError(err)

		var sizeOptionID, sValueID uuid.UUID
		for _, opt := range product.Options {
			if opt.Name == "Kích thước" {
				sizeOptionID = opt.ID
				for _, val := range opt.Values {
					if val.Value == "S" {
						sValueID = val.ID
						break
					}
				}
				break
			}
		}
		s.Require().NotEqual(uuid.Nil, sizeOptionID)
		s.Require().NotEqual(uuid.Nil, sValueID)

		result, err := s.app.UpdateOptionValues(ctx, http_dto.UpdateProductOptionValuesRequestDto{
			ProductID: s.firstProductID,
			OptionID:  sizeOptionID,
			Data: []http_dto.UpdateProductOptionValuesData{
				{
					ID:    sValueID,
					Value: "Small",
				},
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
	})

	s.Run("Add more variants to existing product", func() {
		// Get product to retrieve option value IDs
		product, err := s.app.Get(ctx, http_dto.GetProductRequestDto{
			ProductID: s.firstProductID,
		})
		s.Require().NoError(err)

		// Find option value IDs for "Small" (formerly "S") and "Blue"
		var smallValueID, blueValueID, mValueID, redValueID uuid.UUID
		for _, opt := range product.Options {
			switch opt.Name {
			case "Kích thước":
				for _, val := range opt.Values {
					switch val.Value {
					case "Small":
						smallValueID = val.ID
					case "M":
						mValueID = val.ID
					}
				}
			case "Color":
				for _, val := range opt.Values {
					switch val.Value {
					case "Blue":
						blueValueID = val.ID
					case "Red":
						redValueID = val.ID
					}
				}
			}
		}
		s.Require().NotEqual(uuid.Nil, smallValueID, "Small option value should exist")
		s.Require().NotEqual(uuid.Nil, blueValueID, "Blue option value should exist")
		s.Require().NotEqual(uuid.Nil, mValueID, "M option value should exist")
		s.Require().NotEqual(uuid.Nil, redValueID, "Red option value should exist")

		// Add 2 missing variants: Small+Blue and M+Red
		result, err := s.app.AddVariants(ctx, http_dto.AddProductVariantsRequestDto{
			ProductID: s.firstProductID,
			Data: []http_dto.AddProductVariantsData{
				{
					SKU:            "CONFIG-TEST-S-BLUE",
					Price:          105000,
					Quantity:       45,
					OptionValueIDs: []uuid.UUID{smallValueID, blueValueID},
				},
				{
					SKU:            "CONFIG-TEST-M-RED",
					Price:          125000,
					Quantity:       65,
					OptionValueIDs: []uuid.UUID{mValueID, redValueID},
				},
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Len(*result, 2, "Should have added 2 variants")

		// Verify product now has 6 variants
		updatedProduct, err := s.app.Get(ctx, http_dto.GetProductRequestDto{
			ProductID: s.firstProductID,
		})
		s.Require().NoError(err)
		s.Len(updatedProduct.Variants, 6, "Product should now have 6 variants")

		// Verify new variants have correct option values
		for _, variant := range updatedProduct.Variants {
			if variant.SKU == "CONFIG-TEST-S-BLUE" || variant.SKU == "CONFIG-TEST-M-RED" {
				s.Len(variant.OptionValues, 2, "New variants should have 2 option values")
			}
		}
	})

	s.Run("Update specific variant", func() {
		product, err := s.app.Get(ctx, http_dto.GetProductRequestDto{
			ProductID: s.firstProductID,
		})
		s.Require().NoError(err)

		// Find variant by SKU
		var variantID uuid.UUID
		for _, v := range product.Variants {
			if v.SKU == "CONFIG-TEST-M-BLUE" {
				variantID = v.ID
				break
			}
		}
		s.Require().NotEqual(uuid.Nil, variantID)

		result, err := s.app.UpdateVariant(ctx, http_dto.UpdateProductVariantRequestDto{
			ProductID:        s.firstProductID,
			ProductVariantID: variantID,
			Data: http_dto.UpdateProductVariantData{
				Price:    150000,
				Quantity: 100,
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal(int64(150000), result.Price)
		s.Equal(100, result.Quantity)
	})

	s.Run("List products with filters", func() {
		// Filter by category ID
		result, err := s.app.List(ctx, http_dto.ListProductRequestDto{
			PaginationRequestDto: http_dto.PaginationRequestDto{
				Page:  1,
				Limit: 100,
			},
			CategoryIDs: []uuid.UUID{seededCategoryID},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.GreaterOrEqual(result.Meta.TotalItems, 1)

		// Filter by price range
		result2, err := s.app.List(ctx, http_dto.ListProductRequestDto{
			PaginationRequestDto: http_dto.PaginationRequestDto{
				Page:  1,
				Limit: 100,
			},
			MinPrice: 100000,
			MaxPrice: 200000,
		})
		s.Require().NoError(err)
		s.Require().NotNil(result2)

		// Verify all products in price range
		for _, p := range result2.Data {
			s.GreaterOrEqual(p.Price, float64(100000))
			s.LessOrEqual(p.Price, float64(200000))
		}

		// Search by name
		result3, err := s.app.List(ctx, http_dto.ListProductRequestDto{
			PaginationRequestDto: http_dto.PaginationRequestDto{
				Page:  1,
				Limit: 100,
			},
			Search: "Configurable",
		})
		s.Require().NoError(err)
		s.Require().NotNil(result3)
		s.GreaterOrEqual(result3.Meta.TotalItems, 1)
	})

	s.Run("Delete product with options", func() {
		err := s.app.Delete(ctx, http_dto.DeleteProductRequestDto{
			ProductID: s.firstProductID,
		})
		s.Require().NoError(err)

		// Verify product is soft-deleted
		_, err = s.app.Get(ctx, http_dto.GetProductRequestDto{
			ProductID: s.firstProductID,
		})
		s.Require().Error(err)
	})
}
