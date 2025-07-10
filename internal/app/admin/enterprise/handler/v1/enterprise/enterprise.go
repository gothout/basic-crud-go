package enterprise

import (
	"basic-crud-go/internal/app/admin/enterprise/controller"
	"basic-crud-go/internal/app/admin/enterprise/service"

	"github.com/gin-gonic/gin"
)

func RegisterEnterpriseRoutes(router *gin.RouterGroup) {
	//Service
	svc := service.NewEnterpriseService()
	//Conteroller
	ctrl := controller.NewEnterpriseController(svc)

	group := router.Group("/")
	{
		group.GET("ping", ctrl.Ping)
		group.POST("", ctrl.CreateEnterpriseHandler)
		group.GET("read", ctrl.ReadEnterprisesHandler)
		group.GET("read/:cnpj", ctrl.ReadEnterpriseHandler)
		group.PUT("", ctrl.UpdateEnterpriseHandler)
		group.DELETE(":cnpj", ctrl.DeleteEnterpriseHandler)
	}

}
