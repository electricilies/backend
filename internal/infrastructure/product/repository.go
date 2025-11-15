package product

import (
	"context"
	"fmt"
	"strings"
	"time"

	"backend/config"
	"backend/internal/domain/product"
	"backend/internal/infrastructure/errors"
	"backend/internal/infrastructure/persistence/postgres"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type RepositoryImpl struct {
	db              *postgres.Queries
	s3Client        *s3.Client
	s3PresignClient *s3.PresignClient
	redisClient     *redis.Client
	cfg             *config.Config
}

func NewRepository(
	query *postgres.Queries,
	s3Client *s3.Client,
	s3PresignClient *s3.PresignClient,
	redisClient *redis.Client,
	cfg *config.Config,
) product.Repository {
	return &RepositoryImpl{
		db:              query,
		s3Client:        s3Client,
		s3PresignClient: s3PresignClient,
		redisClient:     redisClient,
		cfg:             cfg,
	}
}

func ProvideRepository(
	query *postgres.Queries,
	s3Client *s3.Client,
	s3PresignClient *s3.PresignClient,
	redisClient *redis.Client,
	cfg *config.Config,
) *RepositoryImpl {
	return &RepositoryImpl{
		db:              query,
		s3Client:        s3Client,
		s3PresignClient: s3PresignClient,
		redisClient:     redisClient,
		cfg:             cfg,
	}
}

func (r *RepositoryImpl) GetUploadImageURL(ctx context.Context) (*product.UploadImageURLModel, error) {
	randomUUID, _ := uuid.NewV7()
	key := fmt.Sprintf("temp/products/image/%s", randomUUID.String()) // TODO: use env variable for folder
	url, err := r.s3PresignClient.PresignPutObject(
		ctx,
		&s3.PutObjectInput{
			Bucket: aws.String(r.cfg.S3Bucket),
			Key:    aws.String(key),
		},
		s3.WithPresignExpires(10*time.Minute),
	)
	if err != nil {
		return nil, errors.ToDomainErrorFromS3(err)
	}

	key = strings.Replace(key, "temp/", "", 1)
	model := &UploadURLImage{
		URL: &url.URL,
		Key: &key,
	}
	return model.ToDomain(), nil
}

// FIXME: param id? what id? product id? variant id?
func (r *RepositoryImpl) GetDeleteImageURL(ctx context.Context, imageID int) (string, error) {
	// TODO: get image URL from DB using id
	// TODO: from Kev, how about the tiki data? it's not from our S3 bucket
	imageURL := &struct {
		URL string
	}{
		URL: "products/image/example-image.jpg",
	}
	url, err := r.s3PresignClient.PresignDeleteObject(
		ctx,
		&s3.DeleteObjectInput{
			Bucket: aws.String(r.cfg.S3Bucket),
			Key:    aws.String(imageURL.URL),
		},
		s3.WithPresignExpires(10*time.Minute),
	)
	if err != nil {
		return "", errors.ToDomainErrorFromS3(err)
	}
	return url.URL, nil
}

func (r *RepositoryImpl) MoveImage(ctx context.Context, key string) error {
	_, err := r.s3Client.CopyObject(ctx, &s3.CopyObjectInput{
		Bucket:     aws.String(r.cfg.S3Bucket),
		CopySource: aws.String(fmt.Sprintf("%s/%s", r.cfg.S3Bucket, "temp/"+key)),
		Key:        aws.String(key),
	})
	if err != nil {
		return errors.ToDomainErrorFromS3(err)
	}
	_, err = r.s3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(r.cfg.S3Bucket),
		Key:    aws.String("temp/" + key),
	})
	if err != nil {
		return errors.ToDomainErrorFromS3(err)
	}
	return nil
}

// TODO: implement
func (r *RepositoryImpl) List(ctx context.Context, queryParams *product.QueryParams) (*product.PaginationModel, error) {
	result := &product.PaginationModel{}
	return result, nil
}

func (r *RepositoryImpl) Create(ctx context.Context, model *product.Model) (*product.Model, error) {
	return &product.Model{}, nil
}

func (r *RepositoryImpl) Update(ctx context.Context, model *product.Model, id int) (*product.Model, error) {
	return &product.Model{}, nil
}

func (r *RepositoryImpl) Delete(ctx context.Context, id int) error {
	return nil
}

func (r *RepositoryImpl) AddOption(
	ctx context.Context,
	optionModel *product.OptionModel,
	id int,
) (*product.OptionModel, error) {
	return &product.OptionModel{}, nil
}

func (r *RepositoryImpl) UpdateOption(
	ctx context.Context,
	optionModel *product.OptionModel,
	optionId int,
) (*product.OptionModel, error) {
	return &product.OptionModel{}, nil
}

func (r *RepositoryImpl) AddVariant(
	ctx context.Context,
	variantModel *product.VariantModel,
	id int,
) (*product.VariantModel, error) {
	return &product.VariantModel{}, nil
}

func (r *RepositoryImpl) UpdateVariant(
	ctx context.Context,
	variantModel *product.VariantModel,
	variantId int,
) (*product.VariantModel, error) {
	return &product.VariantModel{}, nil
}

func (r *RepositoryImpl) AddImages(
	ctx context.Context,
	imageModels []*product.ImageModel,
	id int,
) error {
	return nil
}
