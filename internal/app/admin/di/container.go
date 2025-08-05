package di

import (
	entCtrl "basic-crud-go/internal/app/admin/enterprise/controller"
	entRepo "basic-crud-go/internal/app/admin/enterprise/repository"
	entSvc "basic-crud-go/internal/app/admin/enterprise/service"

	permCtrl "basic-crud-go/internal/app/admin/permission/controller"
	permRepo "basic-crud-go/internal/app/admin/permission/repository"
	permSvc "basic-crud-go/internal/app/admin/permission/service"

	userCtrl "basic-crud-go/internal/app/admin/user/controller"
	userRepo "basic-crud-go/internal/app/admin/user/repository"
	userSvc "basic-crud-go/internal/app/admin/user/service"

	"basic-crud-go/internal/infrastructure/db/postgres"
)

type Container struct {
	enterpriseController entCtrl.EnterpriseController
	userController       userCtrl.UserController
	permissionController permCtrl.PermissionController
}

func NewContainer() *Container {
	db := postgres.GetDB()

	// Enterprise layer
	enterpriseRepository := entRepo.NewRepositoryImpl(db)
	enterpriseService := entSvc.NewEnterpriseService(enterpriseRepository)
	enterpriseController := entCtrl.NewEnterpriseController(enterpriseService)

	// User Layer
	userRepository := userRepo.NewUserRepositoryImpl(db)
	userService := userSvc.NewUserService(userRepository, enterpriseService)
	userController := userCtrl.NewUserController(userService)

	// Permission Layer
	permissionRepository := permRepo.NewRepositoryImpl(db)
	permissionService := permSvc.NewPermissionService(permissionRepository, userService)
	permissionController := permCtrl.NewPermissionController(permissionService)

	return &Container{
		enterpriseController: enterpriseController,
		userController:       userController,
		permissionController: permissionController,
	}
}

// Getters

func (c *Container) GetEnterpriseController() entCtrl.EnterpriseController {
	return c.enterpriseController
}

func (c *Container) GetUserController() userCtrl.UserController {
	return c.userController
}

func (c *Container) GetPermissionController() permCtrl.PermissionController {
	return c.permissionController
}
