package service

import (
	"basic-crud-go/internal/app/admin/permission/model"
	"basic-crud-go/internal/app/admin/permission/repository"
	"basic-crud-go/internal/configuration/logger"
	"context"
)

const module string = "Admin-Permission-Service"

type permissionServiceImpl struct {
	repo repository.PermissionRepository
}

func NewPermissionService(repo repository.PermissionRepository) PermissionService {
	return &permissionServiceImpl{
		repo: repo,
	}
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
