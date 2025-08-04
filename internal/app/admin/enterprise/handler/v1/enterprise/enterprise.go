package enterprise

import (
	"basic-crud-go/internal/app/admin/enterprise/controller"
	"basic-crud-go/internal/app/admin/enterprise/repository"
	"basic-crud-go/internal/app/admin/enterprise/service"
	middleware "basic-crud-go/internal/app/middleware"
	"basic-crud-go/internal/infrastructure/db/postgres"

	"github.com/gin-gonic/gin"
)

func RegisterEnterpriseRoutes(router *gin.RouterGroup, mw *middleware.AuthMiddleware) {
	// Repository
	repo := repository.NewRepositoryImpl(postgres.GetDB())
	//Service
	svc := service.NewEnterpriseService(repo)
	//Controller
	ctrl := controller.NewEnterpriseController(svc)

	group := router.Group("/")
	{
		group.POST("", mw.AuthMiddleware("create-enterprise"), ctrl.CreateEnterpriseHandler)
		group.GET("read", mw.AuthMiddleware("read-enterprise"), ctrl.ReadEnterprisesHandler)
		group.GET("read/:cnpj", mw.AuthMiddleware("read-enterprise"), ctrl.ReadEnterpriseHandler)
		group.PUT("", mw.AuthMiddleware("update-enterprise"), ctrl.UpdateEnterpriseHandler)
		group.DELETE(":cnpj", mw.AuthMiddleware("delete-enterprise"), ctrl.DeleteEnterpriseHandler)
	}

}
