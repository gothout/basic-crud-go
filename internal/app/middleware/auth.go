package middleware

import (
	adminMiddlewareService "basic-crud-go/internal/app/admin/middleware/service"
	"basic-crud-go/internal/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	service adminMiddlewareService.MiddlewareService
}

func NewAuthMiddleware(service adminMiddlewareService.MiddlewareService) *AuthMiddleware {
	return &AuthMiddleware{service: service}
}

// AuthMiddleware validates API key and required permission(s).
func (m *AuthMiddleware) AuthMiddleware(requiredCodes ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if len(authHeader) < 8 || authHeader[:7] != "Bearer " {
			restErr := rest_err.NewBadRequestError("Missing or malformed Authorization header")
			ctx.AbortWithStatusJSON(restErr.Code, restErr)
			return
		}

		apiKey := authHeader[7:]
		identity, err := m.service.ValidateApiKey(ctx.Request.Context(), apiKey)
		if err != nil {
			restErr := rest_err.NewForbiddenError("Invalid credentials")
			ctx.AbortWithStatusJSON(restErr.Code, restErr)
			return
		}

		if len(requiredCodes) > 0 && !m.service.HasPermission(requiredCodes, identity.Permissions) {
			restErr := rest_err.NewForbiddenError("Permission denied")
			ctx.AbortWithStatusJSON(restErr.Code, restErr)
			return
		}

		ctx.Set("identity", identity)
		ctx.Next()
	}
}
