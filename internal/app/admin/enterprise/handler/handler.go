package handler

import (
	_ "basic-crud-go/docs"
	v1 "basic-crud-go/internal/app/admin/enterprise/handler/v1"

	"github.com/gin-gonic/gin"
)

func RegisterEnterpriseRoutes(router *gin.Engine) {
	enterpriseGroup := router.Group("/enterprise")
	v1.RegisterV1Routes(enterpriseGroup)
}
