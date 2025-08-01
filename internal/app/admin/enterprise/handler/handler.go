package handler

import (
	v1 "basic-crud-go/internal/app/admin/enterprise/handler/v1"
	mw "basic-crud-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterEnterpriseRoutes(router *gin.Engine, authMW *mw.Auth) {
	enterpriseGroup := router.Group("/enterprise")
	v1.RegisterV1Routes(enterpriseGroup, authMW)
}
