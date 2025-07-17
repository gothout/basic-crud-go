package binding

import (
	"basic-crud-go/internal/app/admin/permission/dto"
	"errors"
	"github.com/gin-gonic/gin"
)

// ValidateReadPermissioDTO read permission by name
func ValidateReadPermissioDTO(c *gin.Context) (*dto.ReadPermissionDTO, error) {
	var input dto.ReadPermissionDTO
	// Bind path param
	if err := c.ShouldBindUri(&input); err != nil {
		return nil, errors.New("name is required")
	}
	return &input, nil
}

// ValidatePermissionsDTO binds optional query parameters for reading users
func ValidatePermissionsDTO(c *gin.Context) *dto.ReadPermissionsDTO {
	var input dto.ReadPermissionsDTO
	_ = c.ShouldBindQuery(&input)
	return &input
}
