package middleware

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/Nerzal/gocloak/v13"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Auth interface {
	Handler() gin.HandlerFunc
}

type auth struct {
	keycloakClient *gocloak.GoCloak
	keycloakHost   string
	clientId       string
	clientSecret   string
	realm          string
}

func (j *auth) Handler() gin.HandlerFunc {
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
		roles := extractRoles(claims, j.clientId)
		requiredRoles := info.ClientRoles
		requiredRoleNames := make([]string, len(*requiredRoles))
		for c, r := range *requiredRoles {
			if c == j.clientId {
				copy(requiredRoleNames, r)
			}
		}
		if equal := reflect.DeepEqual(roles, requiredRoleNames); !equal {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "you don't have enough roles"})
			return

		}
		c.Next()
	}
}

func extractRoles(claims jwt.MapClaims, clientID string) []string {
	var roles []string

	resAccess, _ := claims["resource_access"].(map[string]interface{})
	if resAccess == nil {
		return roles
	}

	clientRes, _ := resAccess[clientID].(map[string]interface{})
	if clientRes == nil {
		return roles
	}

	resRoles, _ := clientRes["roles"].([]interface{})
	for _, r := range resRoles {
		if role, _ := r.(string); role != "" {
			roles = append(roles, role)
		}
	}

	return roles
}
