package server

import (
	_ "basic-crud-go/docs"
	enterpriseHandler "basic-crud-go/internal/app/admin/enterprise/handler"
	userHandler "basic-crud-go/internal/app/admin/user/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(router *gin.Engine) {
	InitSwagger(router)
	userHandler.RegisterUserRoutes(router)
	enterpriseHandler.RegisterEnterpriseRoutes(router)
}

func InitSwagger(router *gin.Engine) {
	router.StaticFS("/docs", gin.Dir("./docs", true))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
