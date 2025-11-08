package product

import (
	"context"
	"fmt"
	"time"

	"backend/config"
	"backend/internal/domain/product"
	"backend/internal/infrastructure/errors"
	"backend/internal/infrastructure/presistence/postgres"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type repositoryImpl struct {
	db              *postgres.Queries
	s3Client        *s3.Client
	s3PresignClient *s3.PresignClient
	redisClient     *redis.Client
}

func NewRepository(
	query *postgres.Queries,
	s3Client *s3.Client,
	s3PresignClient *s3.PresignClient,
	redisClient *redis.Client,
) product.Repository {
	return &repositoryImpl{
		db:              query,
		s3Client:        s3Client,
		s3PresignClient: s3PresignClient,
		redisClient:     redisClient,
	}
}

func (r *repositoryImpl) GetUploadImageURL(ctx context.Context) (string, error) {
	randomUUID, _ := uuid.NewV7()
	key := fmt.Sprintf("products/%s", randomUUID.String())
	url, err := r.s3PresignClient.PresignPutObject(
		ctx,
		&s3.PutObjectInput{
			Bucket: aws.String(config.Cfg.S3Bucket),
			Key:    aws.String(key),
		},
		s3.WithPresignExpires(10*time.Minute),
	)
	if err != nil {
		return "", errors.ToDomainErrorFromS3(err)
	}
	return url.URL, nil
}

func (r *repositoryImpl) GetDeleteImageURL(ctx context.Context, id int) (string, error) {
	imageURL, err := r.db.GetProductImageByID(ctx, *ToGetProductImageByIDParams(id))
	if err != nil {
		return "", errors.ToDomainErrorFromPostgres(err)
	}
	url, err := r.s3PresignClient.PresignDeleteObject(
		ctx,
		&s3.DeleteObjectInput{
			Bucket: aws.String(config.Cfg.S3Bucket),
			Key:    aws.String(imageURL.URL),
		},
		s3.WithPresignExpires(10*time.Minute),
	)
	if err != nil {
		return "", errors.ToDomainErrorFromS3(err)
	}
	return url.URL, nil
}
