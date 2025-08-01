package enterprise

import (
	authRepo "basic-crud-go/internal/app/admin/auth/repository"
	authService "basic-crud-go/internal/app/admin/auth/service"
	"basic-crud-go/internal/app/admin/enterprise/controller"
	entRepo "basic-crud-go/internal/app/admin/enterprise/repository"
	"basic-crud-go/internal/app/admin/enterprise/service"
	adminMW "basic-crud-go/internal/app/admin/middleware/service"
	permRepo "basic-crud-go/internal/app/admin/permission/repository"
	permService "basic-crud-go/internal/app/admin/permission/service"
	userRepo "basic-crud-go/internal/app/admin/user/repository"
	userService "basic-crud-go/internal/app/admin/user/service"
	mw "basic-crud-go/internal/app/middleware"
	"basic-crud-go/internal/infrastructure/db/postgres"

	"github.com/gin-gonic/gin"
)

func RegisterEnterpriseRoutes(router *gin.RouterGroup) {
	db := postgres.GetDB()
	// Enterprise layer
	entRepository := entRepo.NewRepositoryImpl(db)
	svc := service.NewEnterpriseService(entRepository)
	ctrl := controller.NewEnterpriseController(svc)

	// Middleware dependencies
	userRepository := userRepo.NewUserRepositoryImpl(db)
	userSvc := userService.NewUserService(userRepository, svc)
	permRepository := permRepo.NewRepositoryImpl(db)
	permSvc := permService.NewPermissionService(permRepository, userSvc)
	authRepository := authRepo.NewAuthRepositoryImpl(db)
	authSvc := authService.NewAuthService(authRepository, userSvc, permSvc)
	mwSvc := adminMW.NewMiddlewareService(userSvc, authSvc, permSvc)
	authMW := mw.NewAuth(mwSvc)

	group := router.Group("/")
	{
		group.POST("", authMW.Handler("create-enterprise"), ctrl.CreateEnterpriseHandler)
		group.GET("read", ctrl.ReadEnterprisesHandler)
		group.GET("read/:cnpj", ctrl.ReadEnterpriseHandler)
		group.PUT("", authMW.Handler("update-enterprise"), ctrl.UpdateEnterpriseHandler)
		group.DELETE(":cnpj", authMW.Handler("delete-enterprise"), ctrl.DeleteEnterpriseHandler)
	}

}
