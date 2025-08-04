package user

import (
	enterpriseRepo "basic-crud-go/internal/app/admin/enterprise/repository"
	enterpriseService "basic-crud-go/internal/app/admin/enterprise/service"
	controller "basic-crud-go/internal/app/admin/user/controller"
	repository "basic-crud-go/internal/app/admin/user/repository"
	service "basic-crud-go/internal/app/admin/user/service"
	"basic-crud-go/internal/app/middleware"
	"basic-crud-go/internal/infrastructure/db/postgres"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup, mw *middleware.AuthMiddleware) {
	db := postgres.GetDB()

	// Enterprise layer
	entRepo := enterpriseRepo.NewRepositoryImpl(db)
	entSvc := enterpriseService.NewEnterpriseService(entRepo)

	// User layer
	userRepo := repository.NewUserRepositoryImpl(db)
	userSvc := service.NewUserService(userRepo, entSvc)
	userCtrl := controller.NewUserController(userSvc)

	// Routes
	userGroup := router.Group("/")
	{
		userGroup.POST("", mw.AuthMiddleware("system", "create-user-admin", "create-user-enterprise"), userCtrl.CreateUserHandler)
		userGroup.GET(":email", mw.AuthMiddleware("system", "read-user", "read-user-enterprise"), userCtrl.ReadUserHandler)
		userGroup.GET("read", mw.AuthMiddleware("system", "read-user", "read-user-enterprise"), userCtrl.ReadUsersHandler)
		userGroup.GET("read/enterprise", mw.AuthMiddleware("system", "read-user", "read-user-enterprise"), userCtrl.ReadUsersByCnpjHandler)
		userGroup.PUT(":email", mw.AuthMiddleware("system", "update-enterprise-user"), userCtrl.UpdateUserHandler)
		userGroup.DELETE(":email", mw.AuthMiddleware("system", "delete-enterprise-user"), userCtrl.DeleteUserHandler)
	}
}
