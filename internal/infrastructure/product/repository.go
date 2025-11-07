package product

import (
	"context"
	"strconv"
	"time"

	"backend/config"
	"backend/internal/domain/product"
	"backend/internal/infrastructure/presistence/postgres"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
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
	url, err := r.s3PresignClient.PresignPutObject(
		ctx,
		&s3.PutObjectInput{
			Bucket: aws.String(config.Cfg.S3Bucket),
			Key:    aws.String(strconv.FormatInt(time.Now().Unix(), 10)),
		},
		s3.WithPresignExpires(10*time.Minute),
	)
	if err != nil {
		return "", err
	}
	return url.URL, nil
}
