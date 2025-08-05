package user

import (
	"basic-crud-go/internal/app/admin/user/controller"
	"basic-crud-go/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup, mw *middleware.AuthMiddleware, ctrl controller.UserController) {

	// Routes
	userGroup := router.Group("/")
	{
		userGroup.POST("", mw.AuthMiddleware("system", "create-user-admin", "create-user-enterprise"), ctrl.CreateUserHandler)
		userGroup.GET(":email", mw.AuthMiddleware("system", "read-user", "read-user-enterprise"), ctrl.ReadUserHandler)
		userGroup.GET("read", mw.AuthMiddleware("system", "read-user", "read-user-enterprise"), ctrl.ReadUsersHandler)
		userGroup.GET("read/enterprise", mw.AuthMiddleware("system", "read-user", "read-user-enterprise"), ctrl.ReadUsersByCnpjHandler)
		userGroup.PUT(":email", mw.AuthMiddleware("system", "update-enterprise-user"), ctrl.UpdateUserHandler)
		userGroup.DELETE(":email", mw.AuthMiddleware("system", "delete-enterprise-user"), ctrl.DeleteUserHandler)
	}
}
