package product

import (
	"context"
	"fmt"
	"strings"
	"time"

	"backend/config"
	"backend/internal/domain"
	"backend/internal/domain/product"
	"backend/internal/infrastructure/mapper"
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
	srcCfg          *config.Server
}

func NewRepository(
	db *postgres.Queries,
	s3Client *s3.Client,
	s3PresignClient *s3.PresignClient,
	redisClient *redis.Client,
	srvCfg *config.Server,
) product.Repository {
	return &RepositoryImpl{
		db:              db,
		s3Client:        s3Client,
		s3PresignClient: s3PresignClient,
		redisClient:     redisClient,
		srcCfg:          srvCfg,
	}
}

func ProvideRepository(
	db *postgres.Queries,
	s3Client *s3.Client,
	s3PresignClient *s3.PresignClient,
	redisClient *redis.Client,
	srvCfg *config.Server,
) *RepositoryImpl {
	return &RepositoryImpl{
		db:              db,
		s3Client:        s3Client,
		s3PresignClient: s3PresignClient,
		redisClient:     redisClient,
		srcCfg:          srvCfg,
	}
}

func (r *RepositoryImpl) GetUploadImageURL(ctx context.Context) (*product.UploadImageURLModel, error) {
	randomUUID, _ := uuid.NewV7()
	key := fmt.Sprintf("temp/products/image/%s", randomUUID.String()) // TODO: use env variable for folder
	url, err := r.s3PresignClient.PresignPutObject(
		ctx,
		&s3.PutObjectInput{
			Bucket: aws.String(r.srcCfg.S3Bucket),
			Key:    aws.String(key),
		},
		s3.WithPresignExpires(10*time.Minute),
	)
	if err != nil {
		return nil, mapper.ToDomainErrorFromS3(err)
	}

	key = strings.Replace(key, "temp/", "", 1)
	model := &UploadURLImage{
		URL: url.URL,
		Key: key,
	}
	return model.ToDomain(), nil
}

// FIXME: param id? what id? product id? variant id?
func (r *RepositoryImpl) GetDeleteImageURL(ctx context.Context, imageID int) (string, error) {
	// TODO: get image URL from DB using id
	// TODO: from Kev, how about the tiki data? it's not from our S3 bucket
	imageURL, err := r.db.GetProductImage(ctx, *ToGetProductImageParam(imageID))
	if err != nil {
		return "", mapper.ToDomainErrorFromPostgres(err)
	}
	url, err := r.s3PresignClient.PresignDeleteObject(
		ctx,
		&s3.DeleteObjectInput{
			Bucket: aws.String(r.srcCfg.S3Bucket),
			Key:    aws.String(imageURL.URL),
		},
		s3.WithPresignExpires(10*time.Minute),
	)
	if err != nil {
		return "", mapper.ToDomainErrorFromS3(err)
	}
	return url.URL, nil
}

func (r *RepositoryImpl) MoveImage(ctx context.Context, key string) error {
	_, err := r.s3Client.CopyObject(ctx, &s3.CopyObjectInput{
		Bucket:     aws.String(r.srcCfg.S3Bucket),
		CopySource: aws.String(fmt.Sprintf("%s/%s", r.srcCfg.S3Bucket, "temp/"+key)),
		Key:        aws.String(key),
	})
	if err != nil {
		return mapper.ToDomainErrorFromS3(err)
	}
	_, err = r.s3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(r.srcCfg.S3Bucket),
		Key:    aws.String("temp/" + key),
	})
	if err != nil {
		return mapper.ToDomainErrorFromS3(err)
	}
	return nil
}

// TODO: implement cache later
func (r *RepositoryImpl) List(ctx context.Context, queryParams product.QueryParams) (*product.PaginationModel, error) {
	productRows, err := r.db.ListProducts(ctx, *ToListProductsParam(queryParams))
	if err != nil {
		return nil, mapper.ToDomainErrorFromPostgres(err)
	}
	return ListProductRowsToDomain(productRows, queryParams.PaginationParams), nil
}

func (r *RepositoryImpl) Create(ctx context.Context, model product.Model) (*product.Model, error) {
	productEntity, err := r.db.CreateProduct(ctx, *ToCreateProductParams(model))
	if err != nil {
		return nil, mapper.ToDomainErrorFromPostgres(err)
	}
	return ToDomain(productEntity), nil
}

func (r *RepositoryImpl) Update(ctx context.Context, model product.Model, id int) (*product.Model, error) {
	productEntity, err := r.db.UpdateProduct(ctx, *ToUpdateProductParams(model, id))
	if err != nil {
		return nil, mapper.ToDomainErrorFromPostgres(err)
	}
	return ToDomain(productEntity), nil
}

func (r *RepositoryImpl) Deletes(ctx context.Context, id []int) error {
	rowAffected, err := r.db.DeleteProducts(ctx, *ToDeleteProductsParam(id))
	if err != nil {
		return mapper.ToDomainErrorFromPostgres(err)
	}
	if rowAffected == 0 {
		return domain.NewNotFoundError("no products deleted", nil)
	}
	return nil
}

func (r *RepositoryImpl) AddOption(
	ctx context.Context,
	optionModel product.OptionModel,
	id int,
) (*product.OptionModel, error) {
	optionEntity, err := r.db.CreateOption(ctx, *ToCreateOptionParams(optionModel, id))
	if err != nil {
		return nil, mapper.ToDomainErrorFromPostgres(err)
	}
	return OptionToDomain(optionEntity), nil
}

func (r *RepositoryImpl) UpdateOption(
	ctx context.Context,
	optionModel product.OptionModel,
	optionId int,
) (*product.OptionModel, error) {
	return &product.OptionModel{}, nil
}

func (r *RepositoryImpl) AddVariants(
	ctx context.Context,
	variantModel []product.VariantModel,
	id int,
) (*[]product.VariantModel, error) {
	productVariantEntities, err := r.db.CreateProductVariants(ctx, *ToCreateProductVariantParams(variantModel, id))
	if err != nil {
		return nil, mapper.ToDomainErrorFromPostgres(err)
	}
	var productVariantModels []product.VariantModel
	for _, v := range productVariantEntities {
		variant := VariantToDomain(v)
		productVariantModels = append(productVariantModels, *variant)
	}
	return &productVariantModels, nil
}

func (r *RepositoryImpl) UpdateVariants(
	ctx context.Context,
	variantModel []product.VariantModel,
	variantId int,
) (*[]product.VariantModel, error) {
	return &[]product.VariantModel{}, nil
}

func (r *RepositoryImpl) AddImages(
	ctx context.Context,
	imageModels []product.ImageModel,
	id int,
) (*[]product.ImageModel, error) {
	productImageEntities, err := r.db.CreateProductImages(ctx, *ToCreateProductImagesParams(imageModels, id))
	if err != nil {
		return nil, mapper.ToDomainErrorFromPostgres(err)
	}
	var productImageModels []product.ImageModel
	for _, entity := range productImageEntities {
		image := ImageToDomain(entity)
		productImageModels = append(productImageModels, *image)
	}
	return &productImageModels, nil
}
