//go:build integration

package application_test

import (
	"context"
	"testing"

	"backend/config"
	"backend/internal/application"
	"backend/internal/di/client"
	"backend/internal/di/db"
	"backend/internal/infrastructure/product"
	"backend/test/integration/component"

	"github.com/stretchr/testify/suite"
)

type ProductTestSuite struct {
	suite.Suite
	containers *component.Containers
	app        application.Product
}

func (s *ProductTestSuite) newContainersConfig() *component.ContainersConfig {
	containersConfig := component.NewContainersConfig()
	containersConfig.DB.Enabled = true
	containersConfig.MinIO.Enabled = true
	containersConfig.Redis.Enabled = true
	return containersConfig
}

func (s *ProductTestSuite) newConfig(
	ctx context.Context,
	containersConfig *component.ContainersConfig,
) *config.Config {
	dbConnStr, err := s.containers.DB.ConnectionString(ctx, "sslmode=disable")
	s.Require().NoError(err, "failed to get db connection string")

	s3ConnStr, err := s.containers.MinIO.ConnectionString(ctx)
	s.Require().NoError(err, "failed to get s3 connection string")

	redisConnStr, err := s.containers.Redis.ConnectionString(ctx)
	s.Require().NoError(err, "failed to get redis connection string")
	return &config.Config{
		DBURL:        dbConnStr,
		S3AccessKey:  containersConfig.MinIO.User,
		S3SecretKey:  containersConfig.MinIO.Password,
		S3RegionName: "us-east-1",
		S3Endpoint:   s3ConnStr,
		S3Bucket:     containersConfig.MinIO.Bucket,
		RedisAddr:    redisConnStr,
	}
}

func (s *ProductTestSuite) SetupSuite() {
	ctx := s.T().Context()
	containersConfig := s.newContainersConfig()

	var err error
	s.containers, err = component.NewContainers(ctx, containersConfig)
	s.Require().NoError(err, "failed to start containers")

	config := s.newConfig(ctx, containersConfig)

	dbPool := db.NewConnection(config)
	queries := db.New(dbPool)
	s3Client := client.NewS3(config)
	s3PresignClient := client.NewS3Presign(s3Client)
	redisClient := client.NewRedis(config)

	err = component.CreateBucket(ctx, s3Client, config.S3Bucket)
	s.Require().NoError(err, "failed to create s3 bucket")

	productRepo := product.NewRepository(
		queries,
		s3Client,
		s3PresignClient,
		redisClient,
		config,
	)
	s.app = application.NewProduct(productRepo)
}

func (s *ProductTestSuite) TearDownSuite() {
	s.containers.Cleanup(s.T())
}

func (s *ProductTestSuite) TestGetProductImageUploadURL() {
	ctx := s.T().Context()

	// Get MinIO connection string for regex validation
	minioConnStr, err := s.containers.MinIO.ConnectionString(ctx)
	s.Require().NoError(err, "failed to get MinIO connection string")
	s.Require().NotEmpty(minioConnStr, "MinIO connection string should not be empty")

	tests := []struct {
		name        string
		expectError bool
	}{
		{
			name:        "should successfully generate upload URL",
			expectError: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			result, err := s.app.GetUploadImageURL(ctx)

			if tt.expectError {
				s.Error(err)
				return
			}

			s.NoError(err)
			s.NotNil(result)
			s.NotEmpty(result.URL, "URL should not be empty")
			s.NotEmpty(result.Key, "Key should not be empty")
			s.Contains(result.URL, minioConnStr, "URL should contain MinIO connection string")
		})
	}
}

func TestProductTestSuite(t *testing.T) {
	suite.Run(t, new(ProductTestSuite))
}
