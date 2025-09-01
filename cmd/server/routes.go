package server

import (
	_ "basic-crud-go/docs"
	authHandler "basic-crud-go/internal/app/admin/auth/handler"
	diAdmin "basic-crud-go/internal/app/admin/di"
	enterpriseHandler "basic-crud-go/internal/app/admin/enterprise/handler"
	permissionHandler "basic-crud-go/internal/app/admin/permission/handler"
	userHandler "basic-crud-go/internal/app/admin/user/handler"
	"basic-crud-go/internal/app/middleware"
	middlewareHandler "basic-crud-go/internal/app/middleware/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(router *gin.Engine) {
	// Instance admin controllers
	mw := middlewareHandler.InitAuthMiddleware()
	RegisterAdminRoutes(router, mw)
	InitSwagger(router)
}

func RegisterAdminRoutes(router *gin.Engine, mw *middleware.AuthMiddleware) {
	container := diAdmin.NewContainer()
	enterpriseHandler.RegisterEnterpriseRoutes(router, mw, container.GetEnterpriseController())
	userHandler.RegisterUserRoutes(router, mw, container.GetUserController())
	permissionHandler.RegisterPermissionRoutes(router, mw, container.GetPermissionController())
	authHandler.RegisterAuthRoutes(router)
}

func InitSwagger(router *gin.Engine) {
	router.StaticFS("/docs", gin.Dir("./docs", true))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
