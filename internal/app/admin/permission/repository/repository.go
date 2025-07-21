package repository

import (
	"basic-crud-go/internal/app/admin/permission/model"
	"context"
)

type PermissionRepository interface {
	ApplyPermissionUser(ctx context.Context, userID string, code string) error
	Search(ctx context.Context, name string) ([]model.Permission, error)
	ReadByCode(ctx context.Context, code string) (*model.Permission, error)
	ReadAllPermissions(ctx context.Context, page, limit int) ([]model.Permission, error)
}
