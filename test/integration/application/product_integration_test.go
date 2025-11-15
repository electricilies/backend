//go:build integration

package application_test

import (
	"context"
	"log"
	"strings"
	"testing"

	"backend/config"
	"backend/internal/application"
	"backend/internal/client"
	domainproduct "backend/internal/domain/product"
	"backend/internal/infrastructure/persistence/postgres"
	infraproduct "backend/internal/infrastructure/product"
	"backend/test/integration/component"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
)

type ProductTestSuite struct {
	suite.Suite
	containers      *component.Containers
	app             application.Product
	productRepo     domainproduct.Repository
	queries         *postgres.Queries
	s3Client        *s3.Client
	s3PresignClient *s3.PresignClient
	redisClient     *redis.Client
	config          *config.Config
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
	s.T().Helper()

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
		S3Endpoint:   "http://" + s3ConnStr,
		S3Bucket:     containersConfig.MinIO.Bucket,
		RedisAddr:    strings.TrimPrefix(redisConnStr, "redis://"),
	}
}

func (s *ProductTestSuite) SetupSuite() {
	ctx := s.T().Context()
	containersConfig := s.newContainersConfig()

	var err error
	s.containers, err = component.NewContainers(ctx, containersConfig)
	s.Require().NoError(err, "failed to start containers")

	s.config = s.newConfig(ctx, containersConfig)

	dbPool := client.NewDBConnection(ctx, s.config)
	s.Require().NotNil(dbPool, "db pool should not be nil")
	s.queries = client.NewDBQueries(dbPool)
	s.s3Client = client.NewS3(ctx, s.config)
	s.Require().NotNil(s.s3Client, "s3 client should not be nil")
	s.s3PresignClient = client.NewS3Presign(s.s3Client)
	s.redisClient = client.NewRedis(ctx, s.config)

	err = component.CreateBucket(ctx, s.s3Client, s.config.S3Bucket)
	s.Require().NoError(err, "failed to create s3 bucket")

	s.productRepo = infraproduct.NewRepository(
		s.queries,
		s.s3Client,
		s.s3PresignClient,
		s.redisClient,
		s.config,
	)
	s.app = application.NewProduct(s.productRepo)
}

func (s *ProductTestSuite) TearDownSuite() {
	s.containers.Cleanup(s.T())
}

func (s *ProductTestSuite) TestGetProductImageUploadURL() {
	s.T().Parallel()
	ctx := s.T().Context()

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
			log.Println("Generated upload URL:", result.URL)

			s.NoError(err)
			s.NotNil(result)
			s.NotEmpty(result.URL, "URL should not be empty")
			s.NotEmpty(result.Key, "Key should not be empty")
			s.Contains(result.URL, minioConnStr, "URL should contain MinIO connection string")
			s.Contains(result.URL, result.Key, "URL should contain the object key")
		})
	}
}

func (s *ProductTestSuite) TestGetProductImageDeleteURL() {
	s.T().Parallel()
	ctx := s.T().Context()

	minioConnStr, err := s.containers.MinIO.ConnectionString(ctx)
	s.Require().NoError(err, "failed to get MinIO connection string")
	s.Require().NotEmpty(minioConnStr, "MinIO connection string should not be empty")

	_, err = s.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &s.config.S3Bucket,
		Key:    aws.String("test-empty-object"),
		Body:   strings.NewReader(""),
	})
	s.Require().NoError(err, "failed to upload empty object")

	tests := []struct {
		name        string
		imageID     int
		expectError bool
	}{
		{
			name:        "should fail on generating delete URL for non-existing image",
			imageID:     1,
			expectError: true,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			result, err := s.app.GetDeleteImageURL(ctx, tt.imageID)

			if tt.expectError {
				s.Error(err)
				return
			}

			s.NoError(err)
			s.NotEmpty(result, "Delete URL should not be empty")
			s.Contains(result, minioConnStr, "Delete URL should contain MinIO connection string")
			s.Contains(result, "test-empty-object", "Delete URL should contain the correct object key")
		})
	}
}

func TestProductTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(ProductTestSuite))
}
