package server

import (
	_ "basic-crud-go/docs"
	authHandler "basic-crud-go/internal/app/admin/auth/handler"
	enterpriseHandler "basic-crud-go/internal/app/admin/enterprise/handler"
	permissionHandler "basic-crud-go/internal/app/admin/permission/handler"
	userHandler "basic-crud-go/internal/app/admin/user/handler"
	"basic-crud-go/internal/infrastructure/db/postgres"
	middleware "basic-crud-go/internal/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(router *gin.Engine) {
	InitSwagger(router)
	authMW := middleware.NewAuthMiddleware(postgres.GetDB())
	userHandler.RegisterUserRoutes(router)
	enterpriseHandler.RegisterEnterpriseRoutes(router, authMW)
	permissionHandler.RegisterPermissionRoutes(router)
	authHandler.RegisterAuthRoutes(router)
}

func InitSwagger(router *gin.Engine) {
	router.StaticFS("/docs", gin.Dir("./docs", true))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
