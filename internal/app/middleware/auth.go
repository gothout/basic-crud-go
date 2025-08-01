package middleware

import (
	adminmw "basic-crud-go/internal/app/admin/middleware/service"
	"basic-crud-go/internal/configuration/logger"
	"basic-crud-go/internal/configuration/rest_err"
	"strings"

	"github.com/gin-gonic/gin"
)

// Auth is a reusable middleware handler that validates API keys
// and checks for specific permissions.
type Auth struct {
	service adminmw.MiddlewareService
}

// NewAuth creates a new Auth middleware instance.
func NewAuth(service adminmw.MiddlewareService) *Auth {
	return &Auth{service: service}
}

// Handler returns a gin.HandlerFunc that validates the request's
// Authorization header and ensures the user has the required permission.
func (a *Auth) Handler(requiredCode string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) < 8 || !strings.HasPrefix(authHeader, "Bearer ") {
			restErr := rest_err.NewBadRequestError("Missing or malformed Authorization header")
			c.AbortWithStatusJSON(restErr.Code, restErr)
			return
		}
		apiKey := strings.TrimSpace(authHeader[7:])

		identity, err := a.service.ValidateApiKey(c.Request.Context(), apiKey)
		if err != nil {
			logger.LogWithAutoFuncName(logger.Warning, "AuthMiddleware", err.Error())
			restErr := rest_err.NewForbiddenError("Invalid or expired token")
			c.AbortWithStatusJSON(restErr.Code, restErr)
			return
		}

		if requiredCode != "" && !a.service.HasPermission(requiredCode, identity.Permissions) {
			restErr := rest_err.NewForbiddenError("Permission denied")
			c.AbortWithStatusJSON(restErr.Code, restErr)
			return
		}

		// store identity in context for handlers
		c.Set("identity", identity)
		c.Next()
	}
}
