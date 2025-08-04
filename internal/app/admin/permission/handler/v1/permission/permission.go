package permission

import (
	entRepo "basic-crud-go/internal/app/admin/enterprise/repository"
	entService "basic-crud-go/internal/app/admin/enterprise/service"
	"basic-crud-go/internal/app/admin/permission/controller"
	"basic-crud-go/internal/app/admin/permission/repository"
	"basic-crud-go/internal/app/admin/permission/service"
	userRepo "basic-crud-go/internal/app/admin/user/repository"
	userService "basic-crud-go/internal/app/admin/user/service"
	"basic-crud-go/internal/app/middleware"
	"basic-crud-go/internal/infrastructure/db/postgres"
	"github.com/gin-gonic/gin"
)

func RegisterPermissionRoutes(router *gin.RouterGroup, mw *middleware.AuthMiddleware) {
	db := postgres.GetDB()

	// Enterprise Repository
	entRep := entRepo.NewRepositoryImpl(db)
	// Enterprise Service
	entServ := entService.NewEnterpriseService(entRep)

	// User Repository
	userRep := userRepo.NewUserRepositoryImpl(db)
	// User Service
	userServ := userService.NewUserService(userRep, entServ)

	// Permission Repository
	permRepo := repository.NewRepositoryImpl(db)

	// Permission Service
	permService := service.NewPermissionService(permRepo, userServ)

	// Controller
	permController := controller.NewPermissionController(permService)

	group := router.Group("/")
	{
		group.GET("search/", mw.AuthMiddleware("system", "read-permission"), permController.Search)
		group.GET("read/", mw.AuthMiddleware("system", "read-permission"), permController.Read)
		group.GET("", mw.AuthMiddleware("system", "read-permission"), permController.ReadAll)
		group.GET("user/:email", mw.AuthMiddleware("system", "read-permission"), permController.ReadUserPermission)
		group.POST("apply", mw.AuthMiddleware("system", "permission-apply-admin", "permission-apply-enterprise"), permController.Apply)
		group.DELETE("user/:email", mw.AuthMiddleware("system", "permission-apply-admin", "permission-apply-enterprise"), permController.RemoveBatch)
	}

}
