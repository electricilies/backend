package product

import (
	"context"
	"fmt"
	"strings"
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

func (r *repositoryImpl) GetUploadImageURL(ctx context.Context) (*product.UploadImageURLModel, error) {
	randomUUID, _ := uuid.NewV7()
	key := fmt.Sprintf("temp/products/image/%s", randomUUID.String()) // TODO: use env variable for folder
	url, err := r.s3PresignClient.PresignPutObject(
		ctx,
		&s3.PutObjectInput{
			Bucket: aws.String(config.Cfg.S3Bucket),
			Key:    aws.String(key),
		},
		s3.WithPresignExpires(10*time.Minute),
	)
	if err != nil {
		return nil, errors.ToDomainErrorFromS3(err)
	}
	model := &UploadURLImage{
		URL: url.URL,
		Key: strings.Replace(key, "temp/", "", 1),
	}
	return model.ToDomain(), nil
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

func (r *repositoryImpl) MoveImage(ctx context.Context, key string) error {
	_, err := r.s3Client.CopyObject(ctx, &s3.CopyObjectInput{
		Bucket:     aws.String(config.Cfg.S3Bucket),
		CopySource: aws.String(fmt.Sprintf("%s/%s", config.Cfg.S3Bucket, "temp/"+key)),
		Key:        aws.String(key),
	})
	if err != nil {
		return errors.ToDomainErrorFromS3(err)
	}
	_, err = r.s3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(config.Cfg.S3Bucket),
		Key:    aws.String("temp/" + key),
	})
	if err != nil {
		return errors.ToDomainErrorFromS3(err)
	}
	return nil
}
