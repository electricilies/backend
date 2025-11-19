package http

import (
	"net/http"
	"strings"

	"backend/config"

	"github.com/Nerzal/gocloak/v13"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware interface {
	Handler() gin.HandlerFunc
}
