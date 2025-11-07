package client

import (
	"context"
	"log"

	"backend/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewS3() *s3.Client {
	cfg, err := awsconfig.LoadDefaultConfig(
		context.Background(),
		awsconfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(config.Cfg.S3AccessKey, config.Cfg.S3SecretKey, "")), awsconfig.WithRegion(config.Cfg.S3RegionName))
	if err != nil {
		log.Printf("failed to load config: %v", err)
		return nil
	}
	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
		o.BaseEndpoint = aws.String(config.Cfg.S3Endpoint)
	})
	exist, err := client.HeadBucket(
		context.Background(),
		&s3.HeadBucketInput{Bucket: aws.String(config.Cfg.S3Bucket)},
	)
	if err != nil || exist == nil {
		log.Printf("bucket not exist: %v", err)
	}

	return client
}

func NewS3Presign(s3Client *s3.Client) *s3.PresignClient {
	return s3.NewPresignClient(s3Client)
}
