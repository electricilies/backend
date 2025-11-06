package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"backend/config"
	"backend/internal/constant"

	"github.com/Nerzal/gocloak/v13"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Auth interface {
	Handler() gin.HandlerFunc
}

type authMiddleware struct {
	keycloakClient *gocloak.GoCloak
	clientId       string
	clientSecret   string
	realm          string
}

func NewJWTVerify(keycloakClient *gocloak.GoCloak) Auth {
	return &authMiddleware{
		keycloakClient: keycloakClient,
		clientId:       config.Cfg.KcClientId,
		clientSecret:   config.Cfg.KcClientSecret,
		realm:          config.Cfg.KcRealm,
	}
}

func (j *authMiddleware) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			return
		}
		token := parts[1]
		rptResult, err := j.keycloakClient.RetrospectToken(ctx, token, j.clientId, j.clientSecret, j.realm)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Failed to introspect token", "detail": err.Error()})
			return
		}
		if rptResult == nil || !*rptResult.Active {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Inactive or invalid token"})
			return
		}
		tokens, _, err := j.keycloakClient.DecodeAccessToken(ctx, token, j.realm)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Cannot decode access token", "detail": err.Error()})
			return
		}

		claims, _ := tokens.Claims.(jwt.MapClaims)
		sub := claims["sub"].(string)

		info, err := j.keycloakClient.GetUserByID(ctx, token, j.realm, sub)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Cannot get user info", "detail": err.Error()})
			return
		}
		if !*info.Enabled {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User is banned"})
			return
		}
		ctx.Set(constant.TokenKey, token)
		fmt.Println("Token in middleware:", token)
		c := context.WithValue(ctx.Request.Context(), constant.TokenKey, token)
		ctx.Request = ctx.Request.WithContext(c)
		ctx.Set("claims", claims)
		ctx.Next()
	}
}
