package repository

import (
	"basic-crud-go/internal/app/admin/permission/model"
	"context"
)

type PermissionRepository interface {
	ReadAllPermissions(ctx context.Context, page, limit int) ([]model.Permission, error)
}
