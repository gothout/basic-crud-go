package permission

import (
	"basic-crud-go/internal/app/admin/permission/controller"

	"basic-crud-go/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterPermissionRoutes(router *gin.RouterGroup, mw *middleware.AuthMiddleware, ctrl controller.PermissionController) {

	group := router.Group("/")
	{
		group.GET("search/", mw.AuthMiddleware("system", "read-permission"), ctrl.Search)
		group.GET("read/", mw.AuthMiddleware("system", "read-permission"), ctrl.Read)
		group.GET("", mw.AuthMiddleware("system", "read-permission"), ctrl.ReadAll)
		group.GET("user/:email", mw.AuthMiddleware("system", "read-permission"), ctrl.ReadUserPermission)
		group.POST("apply", mw.AuthMiddleware("system", "permission-apply-admin", "permission-apply-enterprise"), ctrl.Apply)
		group.DELETE("user/:email", mw.AuthMiddleware("system", "permission-apply-admin", "permission-apply-enterprise"), ctrl.RemoveBatch)
	}

}
