package enterprise

import (
	controller "basic-crud-go/internal/app/admin/enterprise/controller"
	service "basic-crud-go/internal/app/admin/enterprise/service"

	"github.com/gin-gonic/gin"
)

func RegisterEnterpriseRoutes(router *gin.RouterGroup) {
	//Service
	svc := service.NewEnterpriseService()
	//Conteroller
	ctrl := controller.NewEnterpriseController(svc)

	group := router.Group("/")
	{
		group.GET("/ping", ctrl.Ping)
	}

}
