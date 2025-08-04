package v1

import (
	permission "basic-crud-go/internal/app/admin/permission/handler/v1/permission"
	"basic-crud-go/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterV1Routes(router *gin.RouterGroup, mw *middleware.AuthMiddleware) {
	v1Group := router.Group("/v1")
	permission.RegisterPermissionRoutes(v1Group, mw)
}
