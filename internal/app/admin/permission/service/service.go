package service

import (
	"basic-crud-go/internal/app/admin/permission/model"
	"context"
)

type PermissionService interface {
	Read(ctx context.Context, name string) (*model.ModulePermission, error)
}
