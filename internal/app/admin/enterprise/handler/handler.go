package handler

import (
	controller "basic-crud-go/internal/app/admin/enterprise/controller"
	v1 "basic-crud-go/internal/app/admin/enterprise/handler/v1"
	middleware "basic-crud-go/internal/app/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterEnterpriseRoutes(router *gin.Engine, mw *middleware.AuthMiddleware, ctrl controller.EnterpriseController) {
	enterpriseGroup := router.Group("/enterprise")
	v1.RegisterV1Routes(enterpriseGroup, mw, ctrl)
}
