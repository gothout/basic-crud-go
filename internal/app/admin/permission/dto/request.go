package dto

type SearchPermissionDTO struct {
	Query string `form:"query"`
}

type ReadPermissionDTO struct {
	Code string `form:"code"`
}

type ReadUserPermissionsDTO struct {
	Email string `uri:"email" binding:"required"`
}

type ReadPermissionsDTO struct {
	Page  int `form:"page" binding:"omitempty"`
	Limit int `form:"limit" binding:"omitempty"`
}

type ApplyPermissionBatchDTO struct {
	Email string   `json:"email" binding:"required,email"`
	Codes []string `json:"codes" binding:"required,min=1,dive,required"`
}
