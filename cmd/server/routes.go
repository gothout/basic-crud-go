package server

import (
	_ "basic-crud-go/docs"
	authHandler "basic-crud-go/internal/app/admin/auth/handler"
	enterpriseHandler "basic-crud-go/internal/app/admin/enterprise/handler"
	permissionHandler "basic-crud-go/internal/app/admin/permission/handler"
	userHandler "basic-crud-go/internal/app/admin/user/handler"
	middlewareHandler "basic-crud-go/internal/app/middleware/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(router *gin.Engine) {
	InitSwagger(router)
	mw := middlewareHandler.InitAuthMiddleware()
	userHandler.RegisterUserRoutes(router, mw)
	enterpriseHandler.RegisterEnterpriseRoutes(router, mw)
	permissionHandler.RegisterPermissionRoutes(router, mw)
	authHandler.RegisterAuthRoutes(router)
}

func InitSwagger(router *gin.Engine) {
	router.StaticFS("/docs", gin.Dir("./docs", true))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
