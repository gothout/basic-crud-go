package middleware

import (
	"database/sql"
	"strings"

	authRepo "basic-crud-go/internal/app/admin/auth/repository"
	authService "basic-crud-go/internal/app/admin/auth/service"
	entRepo "basic-crud-go/internal/app/admin/enterprise/repository"
	entService "basic-crud-go/internal/app/admin/enterprise/service"
	adminmw "basic-crud-go/internal/app/admin/middleware/service"
	permRepo "basic-crud-go/internal/app/admin/permission/repository"
	permService "basic-crud-go/internal/app/admin/permission/service"
	userRepo "basic-crud-go/internal/app/admin/user/repository"
	userService "basic-crud-go/internal/app/admin/user/service"
	"basic-crud-go/internal/configuration/logger"
	"basic-crud-go/internal/configuration/rest_err"

	"github.com/gin-gonic/gin"
)

// Auth is a reusable middleware handler that validates API keys
// and checks for specific permissions.
type Auth struct {
	service adminmw.MiddlewareService
}

// NewAuth creates a new Auth middleware instance with an existing service.
func NewAuth(service adminmw.MiddlewareService) *Auth {
	return &Auth{service: service}
}

// NewAuthMiddleware initializes all required dependencies using the provided
// database connection and returns a ready-to-use Auth middleware instance.
func NewAuthMiddleware(db *sql.DB) *Auth {
	entRepository := entRepo.NewRepositoryImpl(db)
	entSvc := entService.NewEnterpriseService(entRepository)
	userRepository := userRepo.NewUserRepositoryImpl(db)
	userSvc := userService.NewUserService(userRepository, entSvc)
	permRepository := permRepo.NewRepositoryImpl(db)
	permSvc := permService.NewPermissionService(permRepository, userSvc)
	authRepository := authRepo.NewAuthRepositoryImpl(db)
	authSvc := authService.NewAuthService(authRepository, userSvc, permSvc)
	mwSvc := adminmw.NewMiddlewareService(userSvc, authSvc, permSvc)
	return &Auth{service: mwSvc}
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
