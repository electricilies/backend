package middleware

import (
	"backend/internal/constant"
	"context"
	"net/http"
	"strings"

	"github.com/Nerzal/gocloak/v13"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type JWTVerify interface {
	Handler() gin.HandlerFunc
}

type jwtVerify struct {
	keycloakClient *gocloak.GoCloak
	clientId       string
	clientSecret   string
	realm          string
}

func NewJWTVerify(keycloakClient *gocloak.GoCloak, clientId, clientSecret, realm string) JWTVerify {
	return &jwtVerify{
		keycloakClient: keycloakClient,
		clientId:       clientId,
		clientSecret:   clientSecret,
		realm:          realm,
	}
}

func (j *jwtVerify) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			return
		}
		token := parts[1]
		rptResult, err := j.keycloakClient.RetrospectToken(c, token, j.clientId, j.clientSecret, j.realm)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Failed to introspect token", "detail": err.Error()})
			return
		}
		if rptResult == nil || !*rptResult.Active {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Inactive or invalid token"})
			return
		}
		tokens, _, err := j.keycloakClient.DecodeAccessToken(c, token, j.realm)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Cannot decode access token", "detail": err.Error()})
			return
		}

		claims, _ := tokens.Claims.(jwt.MapClaims)
		sub := claims["sub"].(string)

		info, err := j.keycloakClient.GetUserByID(c, token, j.realm, sub)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Cannot get user info", "detail": err.Error()})
			return
		}
		if !*info.Enabled {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User is banned"})
			return
		}
		c.Set(constant.TokenKey, token)
		ctx := context.WithValue(c.Request.Context(), constant.TokenKey, token)
		c.Request = c.Request.WithContext(ctx)
		c.Set("claims", claims)
		c.Next()
	}
}
