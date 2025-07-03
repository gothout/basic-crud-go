package v1

import (
	enterprise "basic-crud-go/internal/app/admin/enterprise/handler/v1/enterprise"

	"github.com/gin-gonic/gin"
)

func RegisterV1Routes(router *gin.RouterGroup) {
	v1Group := router.Group("/v1")
	enterprise.RegisterEnterpriseRoutes(v1Group)
}
