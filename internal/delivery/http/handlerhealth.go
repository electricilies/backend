package http

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

type HealthHandler interface {
	Liveness(*gin.Context)
	Readiness(*gin.Context)
}


