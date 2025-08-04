package handler

import (
	authRepo "basic-crud-go/internal/app/admin/auth/repository"
	authService "basic-crud-go/internal/app/admin/auth/service"
	enterpriseRepo "basic-crud-go/internal/app/admin/enterprise/repository"
	enterpriseService "basic-crud-go/internal/app/admin/enterprise/service"
	mwService "basic-crud-go/internal/app/admin/middleware/service"
	permissionRepo "basic-crud-go/internal/app/admin/permission/repository"
	permissionService "basic-crud-go/internal/app/admin/permission/service"
	userRepo "basic-crud-go/internal/app/admin/user/repository"
	userService "basic-crud-go/internal/app/admin/user/service"
	"basic-crud-go/internal/app/middleware"
	"basic-crud-go/internal/infrastructure/db/postgres"
)

func InitAuthMiddleware() *middleware.AuthMiddleware {
	db := postgres.GetDB()

	entRepo := enterpriseRepo.NewRepositoryImpl(db)
	entSvc := enterpriseService.NewEnterpriseService(entRepo)

	uRepo := userRepo.NewUserRepositoryImpl(db)
	uSvc := userService.NewUserService(uRepo, entSvc)

	permRepo := permissionRepo.NewRepositoryImpl(db)
	permSvc := permissionService.NewPermissionService(permRepo, uSvc)

	aRepo := authRepo.NewAuthRepositoryImpl(db)
	aSvc := authService.NewAuthService(aRepo, uSvc, permSvc)

	mwSvc := mwService.NewMiddlewareService(uSvc, aSvc, permSvc)
	return middleware.NewAuthMiddleware(mwSvc)
}
