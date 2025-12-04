package http

import "github.com/gin-gonic/gin"

type FlushCacheHandler interface {
	Handler() gin.HandlerFunc
}
