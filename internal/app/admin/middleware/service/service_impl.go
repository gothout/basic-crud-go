package service

import (
	authService "basic-crud-go/internal/app/admin/auth/service"
	"basic-crud-go/internal/app/admin/middleware/model"
	permissionService "basic-crud-go/internal/app/admin/permission/service"
	userService "basic-crud-go/internal/app/admin/user/service"
	"context"
)

type middlewareServiceImpl struct {
	userService       userService.UserService
	authService       authService.AuthService
	permissionService permissionService.PermissionService
}

func NewMiddlewareService(userService userService.UserService, authService authService.AuthService, permissionService permissionService.PermissionService) MiddlewareService {
	return &middlewareServiceImpl{
		userService:       userService,
		authService:       authService,
		permissionService: permissionService,
	}
}

func (m middlewareServiceImpl) ValidateApiKey(ctx context.Context, apiKey string) (*model.UserIndentity, error) {
	// validate token
	userId, err := m.authService.GetUserIdByAPIKey(ctx, apiKey)
	if err != nil {
		return nil, err
	}
	// get identity
	user, enterprise, err := m.userService.ReadById(ctx, userId)
	if err != nil {
		return nil, err
	}
	// get permissions user
	perms, err := m.permissionService.ReadPermissionsUser(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	// convert to []UserPermissions
	var permissions []model.UserPermissions
	for _, perm := range perms {
		p := perm
		permissions = append(permissions, model.UserPermissions{
			Permission: &p,
		})
	}

	identity := &model.UserIndentity{
		User:        user,
		Enterprise:  enterprise,
		Permissions: &permissions,
	}

	return identity, nil
}

// HasPermission checks if any of the required permission codes exist in the user's permissions.
func (m middlewareServiceImpl) HasPermission(requiredCodes []string, permissions *[]model.UserPermissions) bool {
	if permissions == nil {
		return false
	}

	for _, p := range *permissions {
		if p.Permission == nil {
			continue
		}
		for _, code := range requiredCodes {
			if p.Permission.Code == code {
				return true
			}
		}
	}

	return false
}
