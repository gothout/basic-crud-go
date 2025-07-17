package v1

import (
	permission "basic-crud-go/internal/app/admin/permission/handler/v1/permission"
	"github.com/gin-gonic/gin"
)

func RegisterV1Routes(router *gin.RouterGroup) {
	v1Group := router.Group("/v1")
	permission.RegisterPermissionRoutes(v1Group)
}
