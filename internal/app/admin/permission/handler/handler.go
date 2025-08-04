package handler

import (
	v1 "basic-crud-go/internal/app/admin/permission/handler/v1"
	"basic-crud-go/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterPermissionRoutes(router *gin.Engine, mw *middleware.AuthMiddleware) {
	permissionGroup := router.Group("/permission")
	v1.RegisterV1Routes(permissionGroup, mw)
}
