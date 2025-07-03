package handler

import (
	_ "basic-crud-go/docs"
	v1 "basic-crud-go/internal/app/admin/user/handler/v1"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RegisterUserRoutes register all routes domain user
func RegisterUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/user")
	// init Swagger.
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1.RegisterV1Routes(userGroup)
}
