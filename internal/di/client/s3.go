package client

import (
	beconf "backend/config"
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/viper"
)

func NewS3() *s3.Client {
	cfg, err := config.LoadDefaultConfig(
		context.Background(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(beconf.Cfg.S3AccessKey, beconf.Cfg.S3SecretKey, "")), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	return s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
		o.BaseEndpoint = aws.String(viper.GetString(beconf.Cfg.S3Endpoint))
	})
}
