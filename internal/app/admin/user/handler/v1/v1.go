package v1

import (
	admin "basic-crud-go/internal/app/admin/user/handler/v1/admin"

	"github.com/gin-gonic/gin"
)

func RegisterV1Routes(router *gin.RouterGroup) {
	v1Group := router.Group("/v1")
	admin.RegisterUserRoutes(v1Group)
}
