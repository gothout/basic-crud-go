package dto

type ReadPermissionDTO struct {
	Name string `uri:"name" binding:"required"`
}
type ReadPermissionsDTO struct {
	Page  int `form:"page" binding:"omitempty"`
	Limit int `form:"limit" binding:"omitempty"`
}
