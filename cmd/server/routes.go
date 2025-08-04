package server

import (
	_ "basic-crud-go/docs"
	authHandler "basic-crud-go/internal/app/admin/auth/handler"
	authRepo "basic-crud-go/internal/app/admin/auth/repository"
	authService "basic-crud-go/internal/app/admin/auth/service"
	enterpriseHandler "basic-crud-go/internal/app/admin/enterprise/handler"
	enterpriseRepo "basic-crud-go/internal/app/admin/enterprise/repository"
	enterpriseService "basic-crud-go/internal/app/admin/enterprise/service"
	mwService "basic-crud-go/internal/app/admin/middleware/service"
	permissionHandler "basic-crud-go/internal/app/admin/permission/handler"
	permissionRepo "basic-crud-go/internal/app/admin/permission/repository"
	permissionService "basic-crud-go/internal/app/admin/permission/service"
	userHandler "basic-crud-go/internal/app/admin/user/handler"
	userRepo "basic-crud-go/internal/app/admin/user/repository"
	userService "basic-crud-go/internal/app/admin/user/service"
	middleware "basic-crud-go/internal/app/middleware"
	"basic-crud-go/internal/infrastructure/db/postgres"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(router *gin.Engine) {
	InitSwagger(router)
	mw := initAuthMiddleware()
	userHandler.RegisterUserRoutes(router)
	enterpriseHandler.RegisterEnterpriseRoutes(router, mw)
	permissionHandler.RegisterPermissionRoutes(router)
	authHandler.RegisterAuthRoutes(router)
}

func InitSwagger(router *gin.Engine) {
	router.StaticFS("/docs", gin.Dir("./docs", true))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func initAuthMiddleware() *middleware.AuthMiddleware {
	db := postgres.GetDB()

	entRepo := enterpriseRepo.NewRepositoryImpl(db)
	entSvc := enterpriseService.NewEnterpriseService(entRepo)

	uRepo := userRepo.NewUserRepositoryImpl(db)
	uSvc := userService.NewUserService(uRepo, entSvc)

	permRepo := permissionRepo.NewRepositoryImpl(db)
	permSvc := permissionService.NewPermissionService(permRepo, uSvc)

	aRepo := authRepo.NewAuthRepositoryImpl(db)
	aSvc := authService.NewAuthService(aRepo, uSvc, permSvc)

	mwSvc := mwService.NewMiddlewareService(uSvc, aSvc, permSvc)
	return middleware.NewAuthMiddleware(mwSvc)
}
