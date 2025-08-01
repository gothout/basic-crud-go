package service

import (
	"basic-crud-go/internal/app/admin/middleware/model"
	"context"
)

type MiddlewareService interface {
	ValidateApiKey(ctx context.Context, apiKey string) (*model.UserIndentity, error)
	HasPermission(requiredCode string, permissions *[]model.UserPermissions) bool
}
