package enterprise

import (
	"basic-crud-go/internal/app/admin/enterprise/controller"

	middleware "basic-crud-go/internal/app/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterEnterpriseRoutes(router *gin.RouterGroup, mw *middleware.AuthMiddleware, ctrl controller.EnterpriseController) {

	group := router.Group("/")
	{
		group.POST("", ctrl.CreateEnterpriseHandler)
		group.GET("read", mw.AuthMiddleware("system", "read-enterprise", "read-user-enterprise"), ctrl.ReadEnterprisesHandler)
		group.GET("read/:cnpj", mw.AuthMiddleware("system", "read-enterprise", "read-user-enterprise"), ctrl.ReadEnterpriseHandler)
		group.PUT("", mw.AuthMiddleware("system", "update-enterprise", "update-enterprise-user"), ctrl.UpdateEnterpriseHandler)
		group.DELETE(":cnpj", mw.AuthMiddleware("system", "delete-enterprise", "delete-enterprise-admin"), ctrl.DeleteEnterpriseHandler)
	}

}
