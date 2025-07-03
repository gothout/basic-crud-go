package handler

import (
	_ "basic-crud-go/docs" // esse import embute os comentários do swag
	v1 "basic-crud-go/internal/app/admin/user/handler/v1"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RegisterUserRoutes registra todas as rotas da área usuario
func RegisterUserRoutes(router *gin.Engine) {
	adminGroup := router.Group("/user")
	// Inicia Swagger.
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1.RegisterV1Routes(adminGroup)
}
