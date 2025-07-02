package v1

import (
	admin "basic-crud-go/internal/app/admin/handler/v1/admin"

	"github.com/gin-gonic/gin"
)

func RegisterV1Routes(router *gin.RouterGroup) {
	v1Group := router.Group("/v1")
	admin.RegisterAdminRoutes(v1Group)
}
