package binding

import (
	"basic-crud-go/internal/app/admin/auth/dto"
	utilUser "basic-crud-go/internal/app/admin/user/util"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

// ValidateLoginDTO validate DTO LoginUserDTO
func ValidateLoginDTO(c *gin.Context) (*dto.LoginUserDTO, error) {
	var input dto.LoginUserDTO
	// Bind body JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		return nil, errors.New("invalid or missing JSON body")
	}
	// Validate email from JSON
	if !utilUser.IsEmailValid(input.Email) {
		return nil, fmt.Errorf("invalid email address: %s", input.Email)
	}
	return &input, nil
}
