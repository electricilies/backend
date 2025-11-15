package middleware

import (
	"net/http"

	"backend/config"
	"backend/internal/constant"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Role interface {
	Handler([]constant.UserRole) gin.HandlerFunc
}

type RoleImpl struct {
	srvCfg *config.Server
}

func NewRole(requiredRoles []constant.UserRole, srvCfg *config.Server) Role {
	return &RoleImpl{
		srvCfg: srvCfg,
	}
}

func ProvideRole(srvCfg *config.Server) *RoleImpl {
	return &RoleImpl{
		srvCfg: srvCfg,
	}
}

func (r *RoleImpl) Handler(rolesAllowed []constant.UserRole) gin.HandlerFunc {
	set := make(map[constant.UserRole]struct{})
	for _, requiredRole := range rolesAllowed {
		set[requiredRole] = struct{}{}
	}
	return func(ctx *gin.Context) {
		claimsInterface, exists := ctx.Get("claims")
		if !exists {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "No JWT claims found, JWT middleware must be used before role middleware",
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
			if _, exists := set[constant.UserRole(role)]; exists {
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
