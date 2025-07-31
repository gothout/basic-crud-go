package binding

import (
	"basic-crud-go/internal/app/admin/auth/dto"
	utilUser "basic-crud-go/internal/app/admin/user/util"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
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

// ValidateRefreshTokenUserDTO validates the RefreshTokenDTO input
func ValidateRefreshTokenUserDTO(c *gin.Context) (*dto.RefreshTokenUserDTO, error) {
	var input dto.RefreshTokenUserDTO

	// Bind JSON body to struct
	if err := c.ShouldBindJSON(&input); err != nil {
		return nil, errors.New("invalid or missing JSON body")
	}

	// Validate email format
	if !utilUser.IsEmailValid(input.Email) {
		return nil, fmt.Errorf("invalid email address: %s", input.Email)
	}

	return &input, nil
}

// ValidateLogoutUserDTO validate DTO LogoutUserDTO
func ValidateLogoutUserDTO(c *gin.Context) (*dto.LogoutUserDTO, error) {
	var input dto.LogoutUserDTO
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

// ValidateCreateTokenAPIDTO validates CreateTokenAPIDto input
func ValidateCreateTokenAPIDTO(c *gin.Context) (*dto.CreateTokenAPIDto, error) {
	var input dto.CreateTokenAPIDto

	// Bind JSON body
	if err := c.ShouldBindJSON(&input); err != nil {
		return nil, errors.New("invalid or missing JSON body")
	}

	// Validate email
	if !utilUser.IsEmailValid(input.Email) {
		return nil, fmt.Errorf("invalid email address: %s", input.Email)
	}

	// Validate expiration time is in the future
	if input.ExpiresAt.Before(time.Now()) {
		return nil, fmt.Errorf("expiration time must be in the future: %s", input.ExpiresAt.Format(time.RFC3339))
	}

	return &input, nil
}
