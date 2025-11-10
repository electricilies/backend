package handler

import (
	"context"
	"net/http"
	"time"

	"backend/config"

	"github.com/Nerzal/gocloak/v13"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
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
	dbConn         *pgxpool.Pool
	cfg            *config.Config
}

func NewHealthCheck(keycloakClient *gocloak.GoCloak,
	redisClient *redis.Client,
	s3Client *s3.Client,
	db *pgxpool.Pool,
	cfg *config.Config,
) HealthCheck {
	return &healthCheck{
		keycloakClient: keycloakClient,
		redisClient:    redisClient,
		s3Client:       s3Client,
		dbConn:         db,
		cfg:            cfg,
	}
}

func (h *healthCheck) Readiness(ctx *gin.Context) {
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	dbStatus, dbErr := IsDbReady(c, h.dbConn)
	redisStatus, redisErr := IsRedisReady(c, h.redisClient)
	s3Status, s3Err := IsS3Ready(c, h.s3Client, h.cfg.S3Bucket)
	keycloakStatus, keycloakErr := IsKeycloakReady(c, h.keycloakClient, h.cfg.KcHttpManagmentPath)

	backendStatus := dbStatus && redisStatus && s3Status && keycloakStatus

	resp := gin.H{
		"backend": gin.H{
			"status": backendStatus,
		},
		"database": gin.H{
			"status": dbStatus,
			"reason": errorToString(dbErr),
		},
		"redis": gin.H{
			"status": redisStatus,
			"reason": errorToString(redisErr),
		},
		"s3": gin.H{
			"status": s3Status,
			"reason": errorToString(s3Err),
		},
		"keycloak": gin.H{
			"status": keycloakStatus,
			"reason": errorToString(keycloakErr),
		},
	}

	if backendStatus {
		ctx.JSON(http.StatusOK, resp)
	} else {
		ctx.JSON(http.StatusServiceUnavailable, resp)
	}
}

func (h *healthCheck) Liveness(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, time.Now())
}

func IsS3Ready(ctx context.Context, s3Client *s3.Client, s3Bucket string) (bool, error) {
	if s3Client == nil {
		return false, ErrNilClient("s3")
	}
	exist, err := s3Client.HeadBucket(
		ctx,
		&s3.HeadBucketInput{Bucket: aws.String(s3Bucket)},
	)
	if err != nil {
		return false, err
	}
	if exist == nil {
		return false, ErrNoResponse("s3")
	}
	return true, nil
}

func IsDbReady(ctx context.Context, dbConn *pgxpool.Pool) (bool, error) {
	if dbConn == nil {
		return false, ErrNilClient("database")
	}
	err := dbConn.Ping(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

func IsKeycloakReady(ctx context.Context, keycloakClient *gocloak.GoCloak, kcManagmentPath string) (bool, error) {
	if keycloakClient == nil {
		return false, ErrNilClient("keycloak")
	}
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, kcManagmentPath+"/health/ready", nil)
	client := &http.Client{Timeout: 2 * time.Second} // TODO: move http client to helper
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != 200 {
		return false, ErrStatusCode("keycloak", resp.StatusCode)
	}
	return true, nil
}

func IsRedisReady(ctx context.Context, redisClient *redis.Client) (bool, error) {
	if redisClient == nil {
		return false, ErrNilClient("redis")
	}
	err := redisClient.Ping(ctx).Err()
	if err != nil {
		return false, err
	}
	return true, nil
}

func errorToString(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

func ErrNilClient(name string) error {
	return &HealthError{msg: name + " client is nil"}
}

func ErrNoResponse(name string) error {
	return &HealthError{msg: name + " no response"}
}

func ErrStatusCode(name string, code int) error {
	return &HealthError{msg: name + " status code: " + http.StatusText(code)}
}

type HealthError struct {
	msg string
}

func (e *HealthError) Error() string {
	return e.msg
}
