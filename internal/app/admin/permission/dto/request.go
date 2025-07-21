package dto

type SearchPermissionDTO struct {
	Query string `form:"query"`
}

type ReadPermissionDTO struct {
	Code string `form:"code"`
}

type ReadPermissionsDTO struct {
	Page  int `form:"page" binding:"omitempty"`
	Limit int `form:"limit" binding:"omitempty"`
}
