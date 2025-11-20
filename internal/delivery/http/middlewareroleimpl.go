package http

import (
	"net/http"

	"backend/config"
	"backend/internal/application"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type RoleMiddlewareImpl struct {
	noJWTFoundErr              string
	invalidJWTErr              string
	noRoleFoundErr             string
	insufficientPermissionsErr string
	srvCfg                     *config.Server
}

var _ RoleMiddleware = &RoleMiddlewareImpl{}

func ProvideRoleMiddleware(srvCfg *config.Server) *RoleMiddlewareImpl {
	return &RoleMiddlewareImpl{
		srvCfg:                     srvCfg,
		noJWTFoundErr:              "no JWT token found",
		invalidJWTErr:              "invalid JWT token",
		noRoleFoundErr:             "no role found in JWT claims",
		insufficientPermissionsErr: "insufficient permissions",
	}
}

func (m *RoleMiddlewareImpl) Handler(rolesAllowed []application.UserRole) gin.HandlerFunc {
	set := make(map[application.UserRole]struct{})
	for _, requiredRole := range rolesAllowed {
		set[requiredRole] = struct{}{}
	}
	return func(ctx *gin.Context) {
		claimsInterface, exists := ctx.Get("claims")
		if !exists {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				NewError(m.noJWTFoundErr))
			return
		}

		claims, ok := claimsInterface.(jwt.MapClaims)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, NewError(m.invalidJWTErr))
			return
		}

		userRoles := extractRolesFromClaims(claims)
		if len(userRoles) == 0 {
			ctx.AbortWithStatusJSON(http.StatusForbidden, NewError(m.noRoleFoundErr))
			return
		}

		allowed := false
		for _, role := range userRoles {
			if _, exists := set[application.UserRole(role)]; exists {
				ctx.Set("user_role", role)
				allowed = true
				break
			}
		}
		if !allowed {
			ctx.AbortWithStatusJSON(http.StatusForbidden, NewError(m.insufficientPermissionsErr))
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
