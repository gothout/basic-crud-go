package enterprise

import (
	"basic-crud-go/internal/app/admin/enterprise/controller"
	"basic-crud-go/internal/app/admin/enterprise/repository"
	"basic-crud-go/internal/app/admin/enterprise/service"
	"basic-crud-go/internal/infrastructure/db/postgres"

	"github.com/gin-gonic/gin"
)

func RegisterEnterpriseRoutes(router *gin.RouterGroup) {
	// Repository
	repo := repository.NewRepositoryImpl(postgres.GetDB())
	//Service
	svc := service.NewEnterpriseService(repo)
	//Conteroller
	ctrl := controller.NewEnterpriseController(svc)

	group := router.Group("/")
	{
		group.POST("", ctrl.CreateEnterpriseHandler)
		group.GET("read", ctrl.ReadEnterprisesHandler)
		group.GET("read/:cnpj", ctrl.ReadEnterpriseHandler)
		group.PUT("", ctrl.UpdateEnterpriseHandler)
		group.DELETE(":cnpj", ctrl.DeleteEnterpriseHandler)
	}

}
