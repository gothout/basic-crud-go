package admin

import (
	controller "basic-crud-go/internal/app/admin/controller"
	service "basic-crud-go/internal/app/admin/service"

	"github.com/gin-gonic/gin"
)

func RegisterAdminRoutes(router *gin.RouterGroup) {
	//Service
	svc := service.NewAdminService()
	//Conteroller
	ctrl := controller.NewAdminController(svc)

	group := router.Group("/")
	{
		group.GET("/ping", ctrl.Ping)
	}

}
