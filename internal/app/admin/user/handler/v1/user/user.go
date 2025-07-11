package user

import (
	enterpriseRepo "basic-crud-go/internal/app/admin/enterprise/repository"
	entepriseService "basic-crud-go/internal/app/admin/enterprise/service"
	controller "basic-crud-go/internal/app/admin/user/controller"
	repository "basic-crud-go/internal/app/admin/user/repository"
	service "basic-crud-go/internal/app/admin/user/service"
	"basic-crud-go/internal/infrastructure/db/postgres"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	db := postgres.GetDB()
	// Repository Service
	repoEnterprise := enterpriseRepo.NewRepositoryImpl(db)
	// Service Enterprise
	entService := entepriseService.NewEnterpriseService(repoEnterprise)
	//Repository
	repo := repository.NewUserRepositoryImpl(db)
	//Service
	svc := service.NewUserService(repo, entService)
	//Conteroller
	ctrl := controller.NewUserController(svc)

	group := router.Group("/")
	{
		group.GET("/ping", ctrl.Ping)
		group.POST("", ctrl.CreateUserHandler)
	}

}
