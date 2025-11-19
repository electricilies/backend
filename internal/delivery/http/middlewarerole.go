package http

import (
	"net/http"

	"backend/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserRole string

type RoleMiddleware interface {
	Handler([]UserRole) gin.HandlerFunc
}

type RoleMiddlewareImpl struct {
	srvCfg *config.Server
}

var _ RoleMiddleware = &RoleMiddlewareImpl{}

func ProvideRoleMiddleware(srvCfg *config.Server) *RoleMiddlewareImpl {
	return &RoleMiddlewareImpl{
		srvCfg: srvCfg,
	}
}

func (m *RoleMiddlewareImpl) Handler(rolesAllowed []UserRole) gin.HandlerFunc {
	set := make(map[UserRole]struct{})
	for _, requiredRole := range rolesAllowed {
		set[requiredRole] = struct{}{}
	}
	return func(ctx *gin.Context) {
		claimsInterface, exists := ctx.Get("claims")
		if !exists {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "No JWT claims found",
				},
			)
			return
		}

		claims, ok := claimsInterface.(jwt.MapClaims)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid JWT claims format"})
			return
		}

		userRoles := extractRolesFromClaims(claims)
		if len(userRoles) == 0 {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "No role found in JWT claims"})
			return
		}

		allowed := false
		for _, role := range userRoles {
			if _, exists := set[UserRole(role)]; exists {
				ctx.Set("user_role", role)
				allowed = true
				break
			}
		}
		if !allowed {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			return
		}

		ctx.Next()
	}
}

func extractRolesFromClaims(claims jwt.MapClaims) []string {
	var roles []string
	if rawRoles, ok := claims["roles"]; ok {
		switch v := rawRoles.(type) {
		case []interface{}:
			for _, r := range v {
				if s, ok := r.(string); ok {
					roles = append(roles, s)
				}
			}
		case []string:
			roles = append(roles, v...)
		}
	}
	if role, ok := claims["role"].(string); ok && role != "" {
		roles = append(roles, role)
	}
	return roles
}
