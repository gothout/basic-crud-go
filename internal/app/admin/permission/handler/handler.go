package handler

import (
	v1 "basic-crud-go/internal/app/admin/permission/handler/v1"
	"github.com/gin-gonic/gin"
)

func RegisterPermissionRoutes(router *gin.Engine) {
	permissionGroup := router.Group("/permission")
	v1.RegisterV1Routes(permissionGroup)
}
