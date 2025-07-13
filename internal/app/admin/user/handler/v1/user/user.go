package user

import (
	enterpriseRepo "basic-crud-go/internal/app/admin/enterprise/repository"
	enterpriseService "basic-crud-go/internal/app/admin/enterprise/service"
	controller "basic-crud-go/internal/app/admin/user/controller"
	repository "basic-crud-go/internal/app/admin/user/repository"
	service "basic-crud-go/internal/app/admin/user/service"
	"basic-crud-go/internal/infrastructure/db/postgres"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
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
		userGroup.GET("/ping", userCtrl.Ping)
		userGroup.POST("", userCtrl.CreateUserHandler)
		userGroup.GET(":email", userCtrl.ReadUserHandler)
		userGroup.GET("read", userCtrl.ReadUsersHandler)
		userGroup.PUT(":email", userCtrl.UpdateUserHandler)
	}
}
