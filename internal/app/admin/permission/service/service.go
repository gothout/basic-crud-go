package service

import (
	"basic-crud-go/internal/app/admin/permission/model"
	"context"
)

type PermissionService interface {
	Search(ctx context.Context, name string) ([]model.Permission, error)
	ReadAllPermissions(ctx context.Context, page, limit int) ([]model.Permission, error)
}
