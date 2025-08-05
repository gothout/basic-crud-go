package handler

import (
	"basic-crud-go/internal/app/admin/permission/controller"
	v1 "basic-crud-go/internal/app/admin/permission/handler/v1"
	"basic-crud-go/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterPermissionRoutes(router *gin.Engine, mw *middleware.AuthMiddleware, ctrl controller.PermissionController) {
	permissionGroup := router.Group("/permission")
	v1.RegisterV1Routes(permissionGroup, mw, ctrl)
}
