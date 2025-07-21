package binding

import (
	"basic-crud-go/internal/app/admin/permission/dto"
	userUtil "basic-crud-go/internal/app/admin/user/util"
	"errors"
	"github.com/gin-gonic/gin"
)

// ValidateSearchPermissionDTO search permission by name
func ValidateSearchPermissionDTO(c *gin.Context) (*dto.SearchPermissionDTO, error) {
	var input dto.SearchPermissionDTO

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

// ValidateReadPermissioDTO read permission by code
func ValidateReadPermissioDTO(c *gin.Context) (*dto.ReadPermissionDTO, error) {
	var input dto.ReadPermissionDTO
	// Faz o bind do query param "code"
	if err := c.ShouldBindQuery(&input); err != nil {
		return nil, errors.New("code parameter is required")
	}

	// Validação mínima manual
	if len(input.Code) < 4 {
		return nil, errors.New("code must be at least 4 characters")
	}
	return &input, nil
}

// ValidatePermissionsDTO binds optional query parameters for reading users
func ValidatePermissionsDTO(c *gin.Context) *dto.ReadPermissionsDTO {
	var input dto.ReadPermissionsDTO
	_ = c.ShouldBindQuery(&input)
	return &input
}

// ValidateApplyPermissionBatchDTO binds and validates batch permission request
func ValidateApplyPermissionBatchDTO(c *gin.Context) (*dto.ApplyPermissionBatchDTO, error) {
	var input dto.ApplyPermissionBatchDTO

	// Bind do corpo JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		return nil, errors.New("invalid JSON body or missing required fields")
	}

	// Validação manual extra (caso queira reforçar)
	if len(input.Codes) == 0 {
		return nil, errors.New("at least one permission code must be provided")
	}
	if input.Email == "" {
		return nil, errors.New("email is required")
	}

	if !userUtil.IsEmailValid(input.Email) {
		return nil, errors.New("invalid email")
	}

	return &input, nil
}
