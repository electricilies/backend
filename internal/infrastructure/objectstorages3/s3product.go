package objectstorages3

import (
	"context"
	"time"

	"backend/config"
	"backend/internal/application"
	"backend/internal/delivery/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

type Product struct {
	s3Client        *s3.Client
	s3PresignClient *s3.PresignClient
	cfgSrv          *config.Server
}

func ProvideProduct(
	s3Client *s3.Client,
	s3PresignClient *s3.PresignClient,
	cfgSrv *config.Server,
) *Product {
	return &Product{
		s3Client:        s3Client,
		s3PresignClient: s3PresignClient,
		cfgSrv:          cfgSrv,
	}
}

var _ application.ProductObjectStorage = (*Product)(nil)

func (p *Product) GetUploadImageURL(ctx context.Context) (*http.UploadImageURLResponseDto, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, ToDomainErrorFromS3(err)
	}
	idStr := id.String()
	url, err := p.s3PresignClient.PresignDeleteObject(
		ctx,
		&s3.DeleteObjectInput{
			Bucket: aws.String(p.cfgSrv.S3Bucket),
			Key:    aws.String(S3ProductImageFolderTemp + idStr),
		},
		s3.WithPresignExpires(10*time.Minute),
	)
	if err != nil {
		return nil, ToDomainErrorFromS3(err)
	}
	return &http.UploadImageURLResponseDto{
		URL: url.URL,
		Key: idStr,
	}, nil
}

func (p *Product) GetDeleteImageURL(ctx context.Context, imageID uuid.UUID) (*http.DeleteImageURLResponseDto, error) {
	url, err := p.s3PresignClient.PresignDeleteObject(
		ctx,
		&s3.DeleteObjectInput{
			Bucket: aws.String(p.cfgSrv.S3Bucket),
			Key:    aws.String(S3ProductImageFolder + imageID.String()),
		},
		s3.WithPresignExpires(10*time.Minute),
	)
	if err != nil {
		return nil, ToDomainErrorFromS3(err)
	}
	return &http.DeleteImageURLResponseDto{
		URL: url.URL,
	}, nil
}

func (p *Product) PersistImageFromTemp(ctx context.Context, key string, imageID uuid.UUID) error {
	_, err := p.s3Client.CopyObject(ctx, &s3.CopyObjectInput{
		Bucket:     aws.String(p.cfgSrv.S3Bucket),
		CopySource: aws.String(p.cfgSrv.S3Bucket + "/" + S3ProductImageFolderTemp + key),
		Key:        aws.String(S3ProductImageFolder + imageID.String()),
	})
	if err != nil {
		return ToDomainErrorFromS3(err)
	}
	_, err = p.s3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(p.cfgSrv.S3Bucket),
		Key:    aws.String(S3ProductImageFolderTemp + key),
	})
	if err != nil {
		return ToDomainErrorFromS3(err)
	}
	return nil
}
