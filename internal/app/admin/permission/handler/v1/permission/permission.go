package permission

import (
	"basic-crud-go/internal/app/admin/permission/controller"
	"basic-crud-go/internal/app/admin/permission/repository"
	"basic-crud-go/internal/app/admin/permission/service"
	"basic-crud-go/internal/infrastructure/db/postgres"
	"github.com/gin-gonic/gin"
)

func RegisterPermissionRoutes(router *gin.RouterGroup) {
	// Repository
	repo := repository.NewRepositoryImpl(postgres.GetDB())
	// Service
	svc := service.NewPermissionService(repo)
	// Controller
	ctrl := controller.NewPermissionController(svc)
	group := router.Group("/")
	{
		group.GET("read/:name", ctrl.Read)
		group.GET("", ctrl.ReadAll)
	}
}
