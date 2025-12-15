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

type ProductLifecycleTestSuite struct {
	suite.Suite
	containers *component.Containers
	app        http_dto.ProductApplication

	// For tracking created resources
	firstProductID uuid.UUID
}

// uploadDummyImage uploads a dummy image to S3 for testing
func (s *ProductLifecycleTestSuite) uploadDummyImage(url string) {
	s.T().Helper()
	req, err := http.NewRequest("PUT", url, strings.NewReader("dummy image data"))
	s.Require().NoError(err)
	req.Header.Set("Content-Type", "image/jpeg")

	resp, err := http.DefaultClient.Do(req)
	s.Require().NoError(err)
	defer func() { _ = resp.Body.Close() }()
	s.Require().Equal(http.StatusOK, resp.StatusCode)
}

func TestProductLifecycleSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(ProductLifecycleTestSuite))
}

func (s *ProductLifecycleTestSuite) newContainersConfig() *component.ContainersConfig {
	containersConfig := component.NewContainersConfig(&component.NewContainersConfigParam{
		DBEnabled:    true,
		RedisEnabled: true,
		MinIOEnabled: true,
	})
	containersConfig.DB.Seed = true
	return containersConfig
}

func (s *ProductLifecycleTestSuite) newConfig(
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

func (s *ProductLifecycleTestSuite) SetupTest() {
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

func (s *ProductLifecycleTestSuite) TearDownTest() {
	s.containers.Cleanup(s.T())
}

func (s *ProductLifecycleTestSuite) TestSimpleProductLifecycle() {
	ctx := s.T().Context()

	// Seeded category ID from seed data
	seededCategoryID := uuid.MustParse("00000000-0000-7000-0000-000000001796")
	seededAttributeID := uuid.MustParse("00000000-0000-7000-0000-000000000003")
	seededAttributeValueID := uuid.MustParse("00000000-0000-7000-0000-000000000104")

	// FIXME: DefectID: DF-P-AR-01
	s.Run("Create simple product with single variant and no options", func() {
		// Get upload URLs for 2 product images
		uploadURL1, err := s.app.GetUploadImageURL(ctx)
		s.Require().NoError(err)
		s.Require().NotEmpty(uploadURL1.URL)
		s.Require().NotEmpty(uploadURL1.Key)
		s.uploadDummyImage(uploadURL1.URL)

		uploadURL2, err := s.app.GetUploadImageURL(ctx)
		s.Require().NoError(err)
		s.Require().NotEmpty(uploadURL2.URL)
		s.Require().NotEmpty(uploadURL2.Key)
		s.uploadDummyImage(uploadURL2.URL)

		// Create product
		result, err := s.app.Create(ctx, http_dto.CreateProductRequestDto{
			Data: http_dto.CreateProductData{
				Name:        "Simple Test Product",
				Description: "This is a simple test product with single variant",
				CategoryID:  seededCategoryID,
				AttributeValueIDs: []http_dto.CreateProductAttributesData{
					{
						AttributeID: seededAttributeID,
						ValueID:     seededAttributeValueID,
					},
				},
				Images: []http_dto.CreateProductImageData{
					{
						Key:   uploadURL1.Key,
						Order: 1,
					},
					{
						Key:   uploadURL2.Key,
						Order: 2,
					},
				},
				Variants: []http_dto.CreateProductVariantData{
					{
						SKU:      "SIMPLE-TEST-001",
						Price:    500000,
						Quantity: 100,
					},
				},
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal("Simple Test Product", result.Name)
		s.Equal("This is a simple test product with single variant", result.Description)
		s.InDelta(float64(500000), result.Price, 0.001)
		s.Len(result.Variants, 1)
		s.Equal("SIMPLE-TEST-001", result.Variants[0].SKU)
		s.Equal(int64(500000), result.Variants[0].Price)
		s.Equal(100, result.Variants[0].Quantity)
		s.Len(result.Images, 2)
		s.NotNil(result.Category)
		s.Equal(seededCategoryID, result.Category.ID)
		s.Len(result.Attributes, 1, "Attributes should be returned in response")
		s.firstProductID = result.ID
	})

	s.Run("Get created product", func() {
		result, err := s.app.Get(ctx, http_dto.GetProductRequestDto{
			ProductID: s.firstProductID,
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal(s.firstProductID, result.ID)
		s.Equal("Simple Test Product", result.Name)
		s.NotNil(result.Category)
		s.Len(result.Attributes, 1, "Attributes should be returned in Get response")
		s.Len(result.Images, 2)
	})

	s.Run("Get product again to test Redis cache hit", func() {
		result, err := s.app.Get(ctx, http_dto.GetProductRequestDto{
			ProductID: s.firstProductID,
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal(s.firstProductID, result.ID)
		s.Equal("Simple Test Product", result.Name)
	})

	s.Run("List products", func() {
		result, err := s.app.List(ctx, http_dto.ListProductRequestDto{
			PaginationRequestDto: http_dto.PaginationRequestDto{
				Page:  1,
				Limit: 100,
			},
			ProductIDs: []uuid.UUID{s.firstProductID},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.GreaterOrEqual(result.Meta.TotalItems, 1)

		// Find our product in the list
		found := false
		for _, p := range result.Data {
			if p.ID == s.firstProductID {
				found = true
				break
			}
		}
		s.True(found, "Created product should be in the list")
	})

	s.Run("List products again to test Redis cache hit", func() {
		result, err := s.app.List(ctx, http_dto.ListProductRequestDto{
			PaginationRequestDto: http_dto.PaginationRequestDto{
				Page:  1,
				Limit: 10,
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.GreaterOrEqual(result.Meta.TotalItems, 1)
	})

	s.Run("Update product basic info", func() {
		result, err := s.app.Update(ctx, http_dto.UpdateProductRequestDto{
			ProductID: s.firstProductID,
			Data: http_dto.UpdateProductData{
				Name:        "Updated Simple Product",
				Description: "Updated description for test product",
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal("Updated Simple Product", result.Name)
		s.Equal("Updated description for test product", result.Description)
	})

	s.Run("Get product after update to verify cache invalidation", func() {
		result, err := s.app.Get(ctx, http_dto.GetProductRequestDto{
			ProductID: s.firstProductID,
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal("Updated Simple Product", result.Name)
		s.Equal("Updated description for test product", result.Description)
	})

	s.Run("Update product variant", func() {
		// First get product to get variant ID
		product, err := s.app.Get(ctx, http_dto.GetProductRequestDto{
			ProductID: s.firstProductID,
		})
		s.Require().NoError(err)
		variantID := product.Variants[0].ID

		result, err := s.app.UpdateVariant(ctx, http_dto.UpdateProductVariantRequestDto{
			ProductID:        s.firstProductID,
			ProductVariantID: variantID,
			Data: http_dto.UpdateProductVariantData{
				Price:    600000,
				Quantity: 150,
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal(int64(600000), result.Price)
		s.Equal(150, result.Quantity)
	})

	s.Run("Add new images to product", func() {
		uploadURL3, err := s.app.GetUploadImageURL(ctx)
		s.Require().NoError(err)
		s.uploadDummyImage(uploadURL3.URL)

		uploadURL4, err := s.app.GetUploadImageURL(ctx)
		s.Require().NoError(err)
		s.uploadDummyImage(uploadURL4.URL)

		result, err := s.app.AddImages(ctx, http_dto.AddProductImagesRequestDto{
			ProductID: s.firstProductID,
			Data: []http_dto.AddProductImageData{
				{
					Key:   uploadURL3.Key,
					Order: 2,
				},
				{
					Key:   uploadURL4.Key,
					Order: 3,
				},
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Len(*result, 2)
	})

	s.Run("Add new images to variant", func() {
		// Get product to get variant ID
		product, err := s.app.Get(ctx, http_dto.GetProductRequestDto{
			ProductID: s.firstProductID,
		})
		s.Require().NoError(err)
		variantID := product.Variants[0].ID

		uploadURL, err := s.app.GetUploadImageURL(ctx)
		s.Require().NoError(err)
		s.uploadDummyImage(uploadURL.URL)

		result, err := s.app.AddImages(ctx, http_dto.AddProductImagesRequestDto{
			ProductID: s.firstProductID,
			Data: []http_dto.AddProductImageData{
				{
					Key:              uploadURL.Key,
					Order:            1,
					ProductVariantID: variantID,
				},
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
	})

	s.Run("Delete specific images", func() {
		// Get product to get image IDs
		product, err := s.app.Get(ctx, http_dto.GetProductRequestDto{
			ProductID: s.firstProductID,
		})
		s.Require().NoError(err)
		s.Require().GreaterOrEqual(len(product.Images), 1)

		imageID := product.Images[0].ID

		err = s.app.DeleteImages(ctx, http_dto.DeleteProductImagesRequestDto{
			ProductID: s.firstProductID,
			ImageIDs:  []uuid.UUID{imageID},
		})
		s.Require().NoError(err)

		// Verify image is soft-deleted
		updatedProduct, err := s.app.Get(ctx, http_dto.GetProductRequestDto{
			ProductID: s.firstProductID,
		})
		s.Require().NoError(err)

		// Count non-deleted images
		nonDeletedCount := 0
		for _, img := range updatedProduct.Images {
			if img.DeletedAt == nil {
				nonDeletedCount++
			}
		}
		s.Less(nonDeletedCount, len(product.Images))
	})

	s.Run("Soft delete product", func() {
		err := s.app.Delete(ctx, http_dto.DeleteProductRequestDto{
			ProductID: s.firstProductID,
		})
		s.Require().NoError(err)
	})

	s.Run("Get deleted product fails", func() {
		_, err := s.app.Get(ctx, http_dto.GetProductRequestDto{
			ProductID: s.firstProductID,
		})
		s.Require().Error(err)
	})

	s.Run("List products excludes deleted", func() {
		result, err := s.app.List(ctx, http_dto.ListProductRequestDto{
			PaginationRequestDto: http_dto.PaginationRequestDto{
				Page:  1,
				Limit: 100,
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)

		// Verify deleted product not in list
		for _, p := range result.Data {
			s.NotEqual(s.firstProductID, p.ID, "Deleted product should not be in list")
		}
	})
}
