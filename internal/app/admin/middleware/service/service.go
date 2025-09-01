package service

import (
	"basic-crud-go/internal/app/admin/middleware/model"
	"context"
)

type MiddlewareService interface {
	ValidateApiKey(ctx context.Context, apiKey string) (*model.UserIndentity, error)
	ValidateUserKey(ctx context.Context, userKey string) (*model.UserIndentity, error)
	HasPermission(requiredCodes []string, permissions *[]model.UserPermissions) bool
}
