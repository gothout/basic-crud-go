package permission

import (
	entRepo "basic-crud-go/internal/app/admin/enterprise/repository"
	entService "basic-crud-go/internal/app/admin/enterprise/service"
	"basic-crud-go/internal/app/admin/permission/controller"
	"basic-crud-go/internal/app/admin/permission/repository"
	"basic-crud-go/internal/app/admin/permission/service"
	userRepo "basic-crud-go/internal/app/admin/user/repository"
	userService "basic-crud-go/internal/app/admin/user/service"
	"basic-crud-go/internal/infrastructure/db/postgres"
	"github.com/gin-gonic/gin"
)

func RegisterPermissionRoutes(router *gin.RouterGroup) {
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
		group.GET("search/", permController.Search)
		group.GET("read/", permController.Read)
		group.GET("", permController.ReadAll)
		group.POST("apply", permController.Apply)
	}

}
