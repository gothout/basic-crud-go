package dto

type ReadPermissionDTO struct {
	Query string `form:"query"`
}

type ReadPermissionsDTO struct {
	Page  int `form:"page" binding:"omitempty"`
	Limit int `form:"limit" binding:"omitempty"`
}
