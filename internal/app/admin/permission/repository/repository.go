package repository

import (
	"basic-crud-go/internal/app/admin/permission/model"
	"context"
)

type PermissionRepository interface {
	Read(ctx context.Context, moduleID int64) (*model.ModulePermission, error)
	ReadModuleByName(ctx context.Context, name string) (*model.ModulePermission, error)
	ReadAllModules(ctx context.Context, page, limit int) ([]model.ModulePermission, error)
}
