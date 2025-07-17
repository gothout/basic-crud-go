package dto

type ReadPermissionDTO struct {
	Name string `uri:"name" binding:"required"`
}
