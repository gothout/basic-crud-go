package service

import (
	"basic-crud-go/internal/app/admin/permission/model"
	"basic-crud-go/internal/app/admin/permission/repository"
	userService "basic-crud-go/internal/app/admin/user/service"
	"basic-crud-go/internal/configuration/logger"
	"context"
	"fmt"
)

const module string = "Admin-Permission-Service"

type permissionServiceImpl struct {
	repo        repository.PermissionRepository
	userService userService.UserService
}

func NewPermissionService(repo repository.PermissionRepository, userService userService.UserService) PermissionService {
	return &permissionServiceImpl{
		repo:        repo,
		userService: userService,
	}
}

func (s *permissionServiceImpl) ApplyPermissionUserBatch(ctx context.Context, email string, codes []string) error {

	// Verify user exist
	user, _, err := s.userService.Read(ctx, email)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	// verify codes
	for i := 0; i < len(codes); i++ {
		_, err := s.repo.ReadByCode(ctx, codes[i])
		if err != nil {
			return fmt.Errorf("permission code '%s' not found", codes[i])
		}
	}

	return s.repo.ApplyPermissionUserBatch(ctx, user.Id, codes)
}

func (s *permissionServiceImpl) Search(ctx context.Context, name string) ([]model.Permission, error) {
	permissions, err := s.repo.Search(ctx, name)
	if err != nil {
		return nil, err
	}

	return permissions, nil
}

func (s *permissionServiceImpl) ReadAllPermissions(ctx context.Context, page, limit int) ([]model.Permission, error) {
	if page < 1 {
		logger.LogWithAutoFuncName(logger.Info, module, "page out of range. Defaulting to 1.")
		page = 1
	}

	if limit <= 0 || limit > 10 {
		logger.LogWithAutoFuncName(logger.Info, module, "limit out of range. Defaulting to 10.")
		limit = 10
	}
	permissions, err := s.repo.ReadAllPermissions(ctx, page, limit)
	if err != nil {
		return nil, err
	}

	return permissions, nil
}

func (s *permissionServiceImpl) ReadByCode(ctx context.Context, code string) (*model.Permission, error) {
	// read by code
	perm, err := s.repo.ReadByCode(ctx, code)
	if err != nil {
		return perm, err
	}
	return perm, nil
}
