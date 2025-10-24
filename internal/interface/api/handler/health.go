package handler

import (
	"backend/config"
	"context"
	"net/http"
	"time"

	"github.com/Nerzal/gocloak/v13"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
)

type HealthCheck interface {
	Liveness(ctx *gin.Context)
	Readiness(ctx *gin.Context)
}

type healthCheck struct {
	keycloakClient *gocloak.GoCloak
	redisClient    *redis.Client
	s3Client       *s3.Client
	dbConn         *pgx.Conn
}

func NewHealthCheck(keycloakClient *gocloak.GoCloak,
	redisClient *redis.Client,
	s3Client *s3.Client,
	db *pgx.Conn,
) HealthCheck {
	return &healthCheck{
		keycloakClient: keycloakClient,
		redisClient:    redisClient,
		s3Client:       s3Client,
		dbConn:         db,
	}
}

func (h *healthCheck) Readiness(ctx *gin.Context) {
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	switch {
	// TODO: create struct all json format
	case IsDbReady(c, h.dbConn):
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"status": "not ready", "reason": "database not ready"})
		return
	case IsRedisReady(c, h.redisClient):
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"status": "not ready", "reason": "redis not ready"})
		return
	case IsS3Ready(c, h.s3Client):
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"status": "not ready", "reason": "minio not ready"})
		return
	case IsKeycloakReady(c, h.keycloakClient):
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"status": "not ready", "reason": "keycloak not ready"})
		return
	default:
		ctx.JSON(http.StatusOK, gin.H{"status": "ready"})
	}
}

func (h *healthCheck) Liveness(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, time.Now())
}

func IsS3Ready(ctx context.Context, s3Client *s3.Client) bool {
	if s3Client == nil {
		return false
	}
	exist, err := s3Client.HeadBucket(
		ctx,
		&s3.HeadBucketInput{Bucket: aws.String(config.Cfg.S3Bucket)},
	)
	return err == nil && exist != nil
}

func IsDbReady(ctx context.Context, dbConn *pgx.Conn) bool {
	if dbConn == nil {
		return false
	}
	err := dbConn.Ping(ctx)
	return err == nil
}

func IsKeycloakReady(ctx context.Context, keycloakClient *gocloak.GoCloak) bool {
	if keycloakClient == nil {
		return false
	}
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, config.Cfg.KcBasePath+"/health/ready", nil)
	client := &http.Client{Timeout: 2 * time.Second} // TODO: move http client to helper
	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer func() { _ = resp.Body.Close() }()
	return resp.StatusCode == 200
}

func IsRedisReady(ctx context.Context, redisClient *redis.Client) bool {
	if redisClient == nil {
		return false
	}
	return redisClient.Ping(ctx).Err() == nil
}
