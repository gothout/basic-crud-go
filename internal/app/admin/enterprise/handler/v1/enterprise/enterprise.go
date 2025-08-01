package enterprise

import (
	"basic-crud-go/internal/app/admin/enterprise/controller"
	entRepo "basic-crud-go/internal/app/admin/enterprise/repository"
	"basic-crud-go/internal/app/admin/enterprise/service"
	"basic-crud-go/internal/infrastructure/db/postgres"
	mw "basic-crud-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterEnterpriseRoutes(router *gin.RouterGroup, authMW *mw.Auth) {
	db := postgres.GetDB()
	entRepository := entRepo.NewRepositoryImpl(db)
	svc := service.NewEnterpriseService(entRepository)
	ctrl := controller.NewEnterpriseController(svc)

	group := router.Group("/")
	{
		group.POST("", authMW.Handler("create-enterprise"), ctrl.CreateEnterpriseHandler)
		group.GET("read", ctrl.ReadEnterprisesHandler)
		group.GET("read/:cnpj", ctrl.ReadEnterpriseHandler)
		group.PUT("", authMW.Handler("update-enterprise"), ctrl.UpdateEnterpriseHandler)
		group.DELETE(":cnpj", authMW.Handler("delete-enterprise"), ctrl.DeleteEnterpriseHandler)
	}

}
