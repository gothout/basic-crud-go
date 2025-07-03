package user

import (
	controller "basic-crud-go/internal/app/admin/user/controller"
	service "basic-crud-go/internal/app/admin/user/service"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	//Service
	svc := service.NewUserService()
	//Conteroller
	ctrl := controller.NewUserController(svc)

	group := router.Group("/")
	{
		group.GET("/ping", ctrl.Ping)
	}

}
