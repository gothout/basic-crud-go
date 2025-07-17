package dto

import "basic-crud-go/internal/app/admin/permission/model"

type ReadPermissionResponse struct {
	Name    string                   `json:"name"`
	Actions []model.PermissionAction `json:"actions"`
}
