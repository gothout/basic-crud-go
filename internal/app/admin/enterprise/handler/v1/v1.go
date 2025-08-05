package v1

import (
	"basic-crud-go/internal/app/admin/enterprise/controller"
	enterprise "basic-crud-go/internal/app/admin/enterprise/handler/v1/enterprise"
	middleware "basic-crud-go/internal/app/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterV1Routes(router *gin.RouterGroup, mw *middleware.AuthMiddleware, ctrl controller.EnterpriseController) {
	v1Group := router.Group("/v1")
	enterprise.RegisterEnterpriseRoutes(v1Group, mw, ctrl)
}
