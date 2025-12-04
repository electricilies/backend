package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type FlushCacheRedisHandler struct {
	redisClient *redis.Client
}

func ProvideFlushCacheRedisHandler(redisClient *redis.Client) *FlushCacheRedisHandler {
	return &FlushCacheRedisHandler{
		redisClient: redisClient,
	}
}

// FlushRedisCache godoc
//
//	@Summary		Flush Redis Cache
//	@Description	Flush all data from Redis cache
//	@Tags			Dev
//	@Accept			json
//	@Produce		json
//	@Success		204
//	@Failure		500	{object}	Error
//	@Router			/dev/flush-cache [post]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *FlushCacheRedisHandler) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := h.redisClient.FlushAll(c).Err()
		if err != nil {
			c.AbortWithStatusJSON(500, NewError(err.Error()))
			return
		}
		c.Status(http.StatusNoContent)
	}
}
