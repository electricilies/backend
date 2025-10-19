package client

import (
	"context"
	"fmt"
	"log"

	"backend/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/viper"
)

func NewS3() *s3.Client {
	cfg, err := awsconfig.LoadDefaultConfig(
		context.Background(),
		awsconfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(config.Cfg.S3AccessKey, config.Cfg.S3SecretKey, "")), awsconfig.WithRegion(config.Cfg.S3RegionName))
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
		o.BaseEndpoint = aws.String(viper.GetString(config.Cfg.S3Endpoint))
	})
	output, err := client.ListBuckets(context.Background(), &s3.ListBucketsInput{})
	if err != nil {
		fmt.Println("Nooooo")
	} else {
		fmt.Println("Yes")
		for bucket := range output.Buckets {
			fmt.Println("Bucket", bucket)
		}
	}
	return client
}
