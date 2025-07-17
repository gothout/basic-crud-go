package service

import (
	"basic-crud-go/internal/app/admin/permission/model"
	"basic-crud-go/internal/app/admin/permission/repository"
	"basic-crud-go/internal/configuration/logger"
	"context"
	"fmt"
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

func (s *permissionServiceImpl) Read(ctx context.Context, name string) (*model.ModulePermission, error) {
	// Get ID
	modId, err := s.repo.ReadModuleByName(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("module not found")
	}
	// Get informations by ID
	mod, err := s.repo.Read(ctx, modId.ID)
	if err != nil {
		return nil, nil
	}
	return mod, nil
}

func (s *permissionServiceImpl) ReadAll(ctx context.Context, page, limit int) ([]model.ModulePermission, error) {
	if page < 1 {
		logger.LogWithAutoFuncName(logger.Info, module, "page out of range. Defaulting to 1.")
		page = 1
	}

	if limit <= 0 || limit > 10 {
		logger.LogWithAutoFuncName(logger.Info, module, "limit out of range. Defaulting to 10.")
		limit = 10
	}
	mods, err := s.repo.ReadAllModules(ctx, page, limit)
	if err != nil {
		return nil, err
	}

	return mods, nil
}
