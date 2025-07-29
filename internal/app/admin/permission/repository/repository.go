package repository

import (
	"basic-crud-go/internal/app/admin/permission/model"
	"context"
)

type PermissionRepository interface {
	ApplyPermissionUserBatch(ctx context.Context, userID string, codes []string) error
	Search(ctx context.Context, name string) ([]model.Permission, error)
	ReadByCode(ctx context.Context, code string) (*model.Permission, error)
	ReadAllPermissions(ctx context.Context, page, limit int) ([]model.Permission, error)
	ReadPermissionUserId(ctx context.Context, id string) ([]model.Permission, error)
}
