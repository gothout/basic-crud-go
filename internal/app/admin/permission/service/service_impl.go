package service

import (
	"basic-crud-go/internal/app/admin/permission/model"
	"basic-crud-go/internal/app/admin/permission/repository"
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
