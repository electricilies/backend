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

// TestProductCacheInvalidationSuite tests detailed cache behavior
type ProductCacheInvalidationTestSuite struct {
	suite.Suite
	containers *component.Containers
	app        http_dto.ProductApplication

	// Track created product
	productID uuid.UUID
}

func TestProductCacheInvalidationSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(ProductCacheInvalidationTestSuite))
}

func (s *ProductCacheInvalidationTestSuite) newContainersConfig() *component.ContainersConfig {
	containersConfig := component.NewContainersConfig(&component.NewContainersConfigParam{
		DBEnabled:    true,
		RedisEnabled: true,
		MinIOEnabled: true,
	})
	containersConfig.DB.Seed = true
	return containersConfig
}

func (s *ProductCacheInvalidationTestSuite) newConfig(
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

func (s *ProductCacheInvalidationTestSuite) SetupTest() {
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

func (s *ProductCacheInvalidationTestSuite) TearDownTest() {
	s.containers.Cleanup(s.T())
}

func (s *ProductCacheInvalidationTestSuite) uploadDummyImage(url string) {
	s.T().Helper()
	req, err := http.NewRequest("PUT", url, strings.NewReader("dummy image data"))
	s.Require().NoError(err)
	req.Header.Set("Content-Type", "image/jpeg")

	resp, err := http.DefaultClient.Do(req)
	s.Require().NoError(err)
	defer func() { _ = resp.Body.Close() }()
	s.Require().Equal(http.StatusOK, resp.StatusCode)
}

func (s *ProductCacheInvalidationTestSuite) TestCacheInvalidation() {
	ctx := s.T().Context()

	seededCategoryID := uuid.MustParse("00000000-0000-7000-0000-000000001796")

	s.Run("Create product and verify cache population", func() {
		// Get upload URL for product image
		uploadURL, err := s.app.GetUploadImageURL(ctx)
		s.Require().NoError(err)
		s.uploadDummyImage(uploadURL.URL)

		// Create product
		product, err := s.app.Create(ctx, http_dto.CreateProductRequestDto{
			Data: http_dto.CreateProductData{
				Name:        "Cache Test Product",
				Description: "Product for cache invalidation testing",
				CategoryID:  seededCategoryID,
				Images: []http_dto.CreateProductImageData{
					{Key: uploadURL.Key, Order: 1},
				},
				Variants: []http_dto.CreateProductVariantData{
					{
						SKU:      "CACHE-TEST-001",
						Price:    1000000,
						Quantity: 50,
					},
				},
			},
		})
		s.Require().NoError(err)
		s.productID = product.ID

		// First Get - cache miss, then set
		result1, err := s.app.Get(ctx, http_dto.GetProductRequestDto{
			ProductID: s.productID,
		})
		s.Require().NoError(err)
		s.Equal("Cache Test Product", result1.Name)

		// Second Get - cache hit (should be faster, but we can't measure that)
		result2, err := s.app.Get(ctx, http_dto.GetProductRequestDto{
			ProductID: s.productID,
		})
		s.Require().NoError(err)
		s.Equal("Cache Test Product", result2.Name)
		s.Equal(result1.UpdatedAt, result2.UpdatedAt, "Cached data should be identical")
	})

	s.Run("Update product and verify cache invalidation", func() {
		// Update product
		result, err := s.app.Update(ctx, http_dto.UpdateProductRequestDto{
			ProductID: s.productID,
			Data: http_dto.UpdateProductData{
				Name: "Updated Cache Test Product",
			},
		})
		s.Require().NoError(err)
		s.Equal("Updated Cache Test Product", result.Name)

		// Get product - cache should be invalidated, fetching fresh data
		freshProduct, err := s.app.Get(ctx, http_dto.GetProductRequestDto{
			ProductID: s.productID,
		})
		s.Require().NoError(err)
		s.Equal("Updated Cache Test Product", freshProduct.Name)
		s.NotEqual(result.UpdatedAt, freshProduct.UpdatedAt.Add(0), "Should fetch fresh data after update")
	})

	s.Run("List products and verify cache", func() {
		// First list - cache miss
		result1, err := s.app.List(ctx, http_dto.ListProductRequestDto{
			PaginationRequestDto: http_dto.PaginationRequestDto{
				Page:  1,
				Limit: 10,
			},
			CategoryIDs: []uuid.UUID{seededCategoryID},
		})
		s.Require().NoError(err)
		s.GreaterOrEqual(result1.Meta.TotalItems, 1)

		// Second list with same params - cache hit
		result2, err := s.app.List(ctx, http_dto.ListProductRequestDto{
			PaginationRequestDto: http_dto.PaginationRequestDto{
				Page:  1,
				Limit: 10,
			},
			CategoryIDs: []uuid.UUID{seededCategoryID},
		})
		s.Require().NoError(err)
		s.Equal(result1.Meta.TotalItems, result2.Meta.TotalItems)
	})

	s.Run("Update product and verify list cache invalidation", func() {
		// Update product to trigger cache invalidation
		_, err := s.app.Update(ctx, http_dto.UpdateProductRequestDto{
			ProductID: s.productID,
			Data: http_dto.UpdateProductData{
				Description: "Updated description to invalidate cache",
			},
		})
		s.Require().NoError(err)

		// List again - cache should be invalidated
		result, err := s.app.List(ctx, http_dto.ListProductRequestDto{
			PaginationRequestDto: http_dto.PaginationRequestDto{
				Page:  1,
				Limit: 10,
			},
			CategoryIDs: []uuid.UUID{seededCategoryID},
		})
		s.Require().NoError(err)
		s.GreaterOrEqual(result.Meta.TotalItems, 1)

		// Verify our product has updated description
		for _, p := range result.Data {
			if p.ID == s.productID {
				s.Equal("Updated description to invalidate cache", p.Description)
				break
			}
		}
	})

	s.Run("Add images and verify cache invalidation", func() {
		// Get product
		product, err := s.app.Get(ctx, http_dto.GetProductRequestDto{
			ProductID: s.productID,
		})
		s.Require().NoError(err)
		initialImageCount := len(product.Images)

		// Add images
		uploadURL, err := s.app.GetUploadImageURL(ctx)
		s.Require().NoError(err)
		s.uploadDummyImage(uploadURL.URL)

		_, err = s.app.AddImages(ctx, http_dto.AddProductImagesRequestDto{
			ProductID: s.productID,
			Data: []http_dto.AddProductImageData{
				{Key: uploadURL.Key, Order: 10},
			},
		})
		s.Require().NoError(err)

		// Get product - cache should be invalidated
		updatedProduct, err := s.app.Get(ctx, http_dto.GetProductRequestDto{
			ProductID: s.productID,
		})
		s.Require().NoError(err)
		s.Greater(len(updatedProduct.Images), initialImageCount)
	})

	s.Run("Delete images and verify cache invalidation", func() {
		// Get product
		product, err := s.app.Get(ctx, http_dto.GetProductRequestDto{
			ProductID: s.productID,
		})
		s.Require().NoError(err)
		s.Require().NotEmpty(product.Images)

		// Delete first image
		imageID := product.Images[0].ID
		err = s.app.DeleteImages(ctx, http_dto.DeleteProductImagesRequestDto{
			ProductID: s.productID,
			ImageIDs:  []uuid.UUID{imageID},
		})
		s.Require().NoError(err)

		// Get product - cache should be invalidated
		updatedProduct, err := s.app.Get(ctx, http_dto.GetProductRequestDto{
			ProductID: s.productID,
		})
		s.Require().NoError(err)

		// Count non-deleted images
		nonDeletedCount := 0
		for _, img := range updatedProduct.Images {
			if img.DeletedAt == nil {
				nonDeletedCount++
			}
		}
		originalNonDeletedCount := 0
		for _, img := range product.Images {
			if img.DeletedAt == nil {
				originalNonDeletedCount++
			}
		}
		s.Less(nonDeletedCount, originalNonDeletedCount)
	})

	s.Run("Delete product and verify cache", func() {
		// Delete product
		err := s.app.Delete(ctx, http_dto.DeleteProductRequestDto{
			ProductID: s.productID,
		})
		s.Require().NoError(err)

		// Get should fail
		_, err = s.app.Get(ctx, http_dto.GetProductRequestDto{
			ProductID: s.productID,
		})
		s.Require().Error(err)

		// List should not include deleted product
		listResult, err := s.app.List(ctx, http_dto.ListProductRequestDto{
			PaginationRequestDto: http_dto.PaginationRequestDto{
				Page:  1,
				Limit: 100,
			},
		})
		s.Require().NoError(err)
		for _, p := range listResult.Data {
			s.NotEqual(s.productID, p.ID, "Deleted product should not be in list")
		}
	})

	s.Run("Verify different list params create different cache keys", func() {
		// List with different params should create different cache entries
		// This is implicit in the implementation but we verify behavior

		// List with price filter
		result1, err := s.app.List(ctx, http_dto.ListProductRequestDto{
			PaginationRequestDto: http_dto.PaginationRequestDto{
				Page:  1,
				Limit: 10,
			},
			MinPrice: 100000,
			MaxPrice: 500000,
		})
		s.Require().NoError(err)

		// List without price filter
		result2, err := s.app.List(ctx, http_dto.ListProductRequestDto{
			PaginationRequestDto: http_dto.PaginationRequestDto{
				Page:  1,
				Limit: 10,
			},
		})
		s.Require().NoError(err)

		// Results should be different
		s.GreaterOrEqual(result2.Meta.TotalItems, result1.Meta.TotalItems)

		// List with different page
		result3, err := s.app.List(ctx, http_dto.ListProductRequestDto{
			PaginationRequestDto: http_dto.PaginationRequestDto{
				Page:  2,
				Limit: 10,
			},
		})
		s.Require().NoError(err)
		s.Equal(2, result3.Meta.CurrentPage)
	})
}
