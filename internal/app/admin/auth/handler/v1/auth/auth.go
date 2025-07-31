package auth

import (
	authenticationController "basic-crud-go/internal/app/admin/auth/controller"
	authenticationRepo "basic-crud-go/internal/app/admin/auth/repository"
	authenticationService "basic-crud-go/internal/app/admin/auth/service"
	enterpriseRepository "basic-crud-go/internal/app/admin/enterprise/repository"
	enterpriseService "basic-crud-go/internal/app/admin/enterprise/service"
	permUserRepository "basic-crud-go/internal/app/admin/permission/repository"
	permUserService "basic-crud-go/internal/app/admin/permission/service"
	userRepository "basic-crud-go/internal/app/admin/user/repository"
	userService "basic-crud-go/internal/app/admin/user/service"
	"basic-crud-go/internal/infrastructure/db/postgres"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.RouterGroup) {
	db := postgres.GetDB()

	// Enterprise layer
	entRepo := enterpriseRepository.NewRepositoryImpl(db)
	entSvc := enterpriseService.NewEnterpriseService(entRepo)

	// User layer
	userRepo := userRepository.NewUserRepositoryImpl(db)
	userSvc := userService.NewUserService(userRepo, entSvc)

	// Permission User layer
	permUserRepo := permUserRepository.NewRepositoryImpl(db)
	permUserSvc := permUserService.NewPermissionService(permUserRepo, userSvc)

	// Auth layer
	authRepo := authenticationRepo.NewAuthRepositoryImpl(db)
	authSvc := authenticationService.NewAuthService(authRepo, userSvc, permUserSvc)
	authCtrl := authenticationController.NewAuthController(authSvc)

	// Routes
	userGroup := router.Group("")
	{
		userGroup.POST("/login", authCtrl.AuthLoginHandler)
		userGroup.POST("/refresh", authCtrl.AuthRefreshHandler)
		userGroup.POST("/logout", authCtrl.AuthLogoutHandler)
		userGroup.POST("/token", authCtrl.AuthCreateTokenHandler)
	}
}
