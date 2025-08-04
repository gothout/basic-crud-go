package v1

import (
	enterprise "basic-crud-go/internal/app/admin/enterprise/handler/v1/enterprise"
	middleware "basic-crud-go/internal/app/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterV1Routes(router *gin.RouterGroup, mw *middleware.AuthMiddleware) {
	v1Group := router.Group("/v1")
	enterprise.RegisterEnterpriseRoutes(v1Group, mw)
}
