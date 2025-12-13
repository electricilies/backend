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

type ProductAbnormalCasesTestSuite struct {
	suite.Suite
	containers *component.Containers
	app        http_dto.ProductApplication
}

func TestProductAbnormalCasesSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(ProductAbnormalCasesTestSuite))
}

func (s *ProductAbnormalCasesTestSuite) newContainersConfig() *component.ContainersConfig {
	containersConfig := component.NewContainersConfig(&component.NewContainersConfigParam{
		DBEnabled:    true,
		RedisEnabled: true,
		MinIOEnabled: true,
	})
	containersConfig.DB.Seed = true
	return containersConfig
}

func (s *ProductAbnormalCasesTestSuite) newConfig(
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

func (s *ProductAbnormalCasesTestSuite) SetupTest() {
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

func (s *ProductAbnormalCasesTestSuite) TearDownTest() {
	s.containers.Cleanup(s.T())
}

func (s *ProductAbnormalCasesTestSuite) uploadDummyImage(url string) {
	s.T().Helper()
	req, err := http.NewRequest("PUT", url, strings.NewReader("dummy image data"))
	s.Require().NoError(err)
	req.Header.Set("Content-Type", "image/jpeg")

	resp, err := http.DefaultClient.Do(req)
	s.Require().NoError(err)
	defer func() { _ = resp.Body.Close() }()
	s.Require().Equal(http.StatusOK, resp.StatusCode)
}

func (s *ProductAbnormalCasesTestSuite) TestProductAbnormalCases() {
	ctx := s.T().Context()

	// Seeded category ID from seed data
	seededCategoryID := uuid.MustParse("00000000-0000-7000-0000-000000001796")
	nonExistentID := uuid.MustParse("00000000-0000-0000-0000-000000000001")

	s.Run("Create product with non-existent category", func() {
		uploadURL, err := s.app.GetUploadImageURL(ctx)
		s.Require().NoError(err)
		s.uploadDummyImage(uploadURL.URL)

		_, err = s.app.Create(ctx, http_dto.CreateProductRequestDto{
			Data: http_dto.CreateProductData{
				Name:        "Test Product",
				Description: "Test Description",
				CategoryID:  nonExistentID,
				Images: []http_dto.CreateProductImageData{
					{Key: uploadURL.Key, Order: 1},
				},
				Variants: []http_dto.CreateProductVariantData{
					{
						SKU:      "TEST-SKU",
						Price:    100000,
						Quantity: 10,
					},
				},
			},
		})
		s.Require().Error(err, "Should fail with non-existent category")
	})

	s.Run("Create product with non-existent attribute", func() {
		uploadURL, err := s.app.GetUploadImageURL(ctx)
		s.Require().NoError(err)
		s.uploadDummyImage(uploadURL.URL)

		_, err = s.app.Create(ctx, http_dto.CreateProductRequestDto{
			Data: http_dto.CreateProductData{
				Name:        "Test Product",
				Description: "Test Description",
				CategoryID:  seededCategoryID,
				AttributeValueIDs: []http_dto.CreateProductAttributesData{
					{
						AttributeID: nonExistentID,
						ValueID:     nonExistentID,
					},
				},
				Images: []http_dto.CreateProductImageData{
					{Key: uploadURL.Key, Order: 1},
				},
				Variants: []http_dto.CreateProductVariantData{
					{
						SKU:      "TEST-SKU-2",
						Price:    100000,
						Quantity: 10,
					},
				},
			},
		})
		s.Require().Error(err, "Should fail with non-existent attribute")
	})

	s.Run("Create product without images", func() {
		_, err := s.app.Create(ctx, http_dto.CreateProductRequestDto{
			Data: http_dto.CreateProductData{
				Name:        "Test Product",
				Description: "Test Description",
				CategoryID:  seededCategoryID,
				Images:      []http_dto.CreateProductImageData{}, // Empty images
				Variants: []http_dto.CreateProductVariantData{
					{
						SKU:      "TEST-SKU-3",
						Price:    100000,
						Quantity: 10,
					},
				},
			},
		})
		s.Require().Error(err, "Should fail without images")
	})

	s.Run("Create product without variants", func() {
		uploadURL, err := s.app.GetUploadImageURL(ctx)
		s.Require().NoError(err)
		s.uploadDummyImage(uploadURL.URL)

		_, err = s.app.Create(ctx, http_dto.CreateProductRequestDto{
			Data: http_dto.CreateProductData{
				Name:        "Test Product",
				Description: "Test Description",
				CategoryID:  seededCategoryID,
				Images: []http_dto.CreateProductImageData{
					{Key: uploadURL.Key, Order: 1},
				},
				Variants: []http_dto.CreateProductVariantData{}, // Empty variants
			},
		})
		s.Require().Error(err, "Should fail without variants")
	})

	s.Run("Create product with invalid name (too short)", func() {
		uploadURL, err := s.app.GetUploadImageURL(ctx)
		s.Require().NoError(err)
		s.uploadDummyImage(uploadURL.URL)

		_, err = s.app.Create(ctx, http_dto.CreateProductRequestDto{
			Data: http_dto.CreateProductData{
				Name:        "AB", // Only 2 characters
				Description: "Test Description",
				CategoryID:  seededCategoryID,
				Images: []http_dto.CreateProductImageData{
					{Key: uploadURL.Key, Order: 1},
				},
				Variants: []http_dto.CreateProductVariantData{
					{
						SKU:      "TEST-SKU-4",
						Price:    100000,
						Quantity: 10,
					},
				},
			},
		})
		s.Require().Error(err, "Should fail with name too short")
	})

	s.Run("Create product with invalid description (too short)", func() {
		uploadURL, err := s.app.GetUploadImageURL(ctx)
		s.Require().NoError(err)
		s.uploadDummyImage(uploadURL.URL)

		_, err = s.app.Create(ctx, http_dto.CreateProductRequestDto{
			Data: http_dto.CreateProductData{
				Name:        "Test Product",
				Description: "Short", // Less than 10 characters
				CategoryID:  seededCategoryID,
				Images: []http_dto.CreateProductImageData{
					{Key: uploadURL.Key, Order: 1},
				},
				Variants: []http_dto.CreateProductVariantData{
					{
						SKU:      "TEST-SKU-5",
						Price:    100000,
						Quantity: 10,
					},
				},
			},
		})
		s.Require().Error(err, "Should fail with description too short")
	})

	s.Run("Create product with invalid price (zero)", func() {
		uploadURL, err := s.app.GetUploadImageURL(ctx)
		s.Require().NoError(err)
		s.uploadDummyImage(uploadURL.URL)

		_, err = s.app.Create(ctx, http_dto.CreateProductRequestDto{
			Data: http_dto.CreateProductData{
				Name:        "Test Product",
				Description: "Test Description",
				CategoryID:  seededCategoryID,
				Images: []http_dto.CreateProductImageData{
					{Key: uploadURL.Key, Order: 1},
				},
				Variants: []http_dto.CreateProductVariantData{
					{
						SKU:      "TEST-SKU-6",
						Price:    0, // Invalid price
						Quantity: 10,
					},
				},
			},
		})
		s.Require().Error(err, "Should fail with zero price")
	})

	s.Run("Get non-existent product", func() {
		_, err := s.app.Get(ctx, http_dto.GetProductRequestDto{
			ProductID: nonExistentID,
		})
		s.Require().Error(err, "Should fail to get non-existent product")
	})

	s.Run("Update non-existent product", func() {
		_, err := s.app.Update(ctx, http_dto.UpdateProductRequestDto{
			ProductID: nonExistentID,
			Data: http_dto.UpdateProductData{
				Name: "New Name",
			},
		})
		s.Require().Error(err, "Should fail to update non-existent product")
	})

	s.Run("Delete non-existent product", func() {
		err := s.app.Delete(ctx, http_dto.DeleteProductRequestDto{
			ProductID: nonExistentID,
		})
		s.Require().Error(err, "Should fail to delete non-existent product")
	})

	s.Run("Add variant to non-existent product", func() {
		_, err := s.app.AddVariants(ctx, http_dto.AddProductVariantsRequestDto{
			ProductID: nonExistentID,
			Data: []http_dto.AddProductVariantsData{
				{
					SKU:      "TEST-SKU-7",
					Price:    100000,
					Quantity: 10,
				},
			},
		})
		s.Require().Error(err, "Should fail to add variant to non-existent product")
	})

	s.Run("Update non-existent variant", func() {
		// Get a real product first
		seededProductID := uuid.MustParse("00000000-0000-7000-0000-000278469304")

		_, err := s.app.UpdateVariant(ctx, http_dto.UpdateProductVariantRequestDto{
			ProductID:        seededProductID,
			ProductVariantID: nonExistentID,
			Data: http_dto.UpdateProductVariantData{
				Price: 200000,
			},
		})
		s.Require().Error(err, "Should fail to update non-existent variant")
	})

	s.Run("Add images with invalid product ID", func() {
		uploadURL, err := s.app.GetUploadImageURL(ctx)
		s.Require().NoError(err)
		s.uploadDummyImage(uploadURL.URL)

		_, err = s.app.AddImages(ctx, http_dto.AddProductImagesRequestDto{
			ProductID: nonExistentID,
			Data: []http_dto.AddProductImageData{
				{Key: uploadURL.Key, Order: 1},
			},
		})
		s.Require().Error(err, "Should fail to add images to non-existent product")
	})

	s.Run("Add images to non-existent variant", func() {
		seededProductID := uuid.MustParse("00000000-0000-7000-0000-000278469304")
		uploadURL, err := s.app.GetUploadImageURL(ctx)
		s.Require().NoError(err)
		s.uploadDummyImage(uploadURL.URL)

		_, err = s.app.AddImages(ctx, http_dto.AddProductImagesRequestDto{
			ProductID: seededProductID,
			Data: []http_dto.AddProductImageData{
				{
					Key:              uploadURL.Key,
					Order:            1,
					ProductVariantID: nonExistentID,
				},
			},
		})
		s.Require().Error(err, "Should fail to add images to non-existent variant")
	})

	s.Run("Delete non-existent images", func() {
		seededProductID := uuid.MustParse("00000000-0000-7000-0000-000278469304")

		err := s.app.DeleteImages(ctx, http_dto.DeleteProductImagesRequestDto{
			ProductID: seededProductID,
			ImageIDs:  []uuid.UUID{nonExistentID},
		})
		// This should not error, it just doesn't delete anything
		s.Require().NoError(err, "Deleting non-existent images should not error")
	})

	s.Run("Update non-existent option", func() {
		seededProductID := uuid.MustParse("00000000-0000-7000-0000-000278469304")

		_, err := s.app.UpdateOptions(ctx, http_dto.UpdateProductOptionsRequestDto{
			ProductID: seededProductID,
			Data: []http_dto.UpdateProductOptionsData{
				{
					ID:   nonExistentID,
					Name: "New Option",
				},
			},
		})
		s.Require().Error(err, "Should fail to update non-existent option")
	})

	s.Run("Update non-existent option value", func() {
		seededProductID := uuid.MustParse("00000000-0000-7000-0000-000278469304")
		seededOptionID := uuid.MustParse("00000000-0000-7000-0000-000000000001")

		_, err := s.app.UpdateOptionValues(ctx, http_dto.UpdateProductOptionValuesRequestDto{
			ProductID: seededProductID,
			OptionID:  seededOptionID,
			Data: []http_dto.UpdateProductOptionValuesData{
				{
					ID:    nonExistentID,
					Value: "New Value",
				},
			},
		})
		s.Require().Error(err, "Should fail to update non-existent option value")
	})
}
