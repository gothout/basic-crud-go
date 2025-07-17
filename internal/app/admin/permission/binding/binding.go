package binding

import (
	"basic-crud-go/internal/app/admin/permission/dto"
	"errors"
	"github.com/gin-gonic/gin"
)

// ValidateReadPermissioDTO read permission by name
func ValidateReadPermissionDTO(c *gin.Context) (*dto.ReadPermissionDTO, error) {
	var input dto.ReadPermissionDTO

	// Faz o bind do query param "query"
	if err := c.ShouldBindQuery(&input); err != nil {
		return nil, errors.New("query parameter is required")
	}

	// Validação mínima manual
	if len(input.Query) < 4 {
		return nil, errors.New("query must be at least 4 characters")
	}

	return &input, nil
}

// ValidatePermissionsDTO binds optional query parameters for reading users
func ValidatePermissionsDTO(c *gin.Context) *dto.ReadPermissionsDTO {
	var input dto.ReadPermissionsDTO
	_ = c.ShouldBindQuery(&input)
	return &input
}
