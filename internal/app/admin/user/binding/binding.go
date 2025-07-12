package binding

import (
	"basic-crud-go/internal/app/admin/user/dto"
	"basic-crud-go/internal/app/admin/user/util"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func ValidateCreateUserDTO(input dto.CreateUserDTO) error {
	switch {
	case strings.TrimSpace(input.Cnpj) == "":
		return errors.New("cnpj is required")
	case strings.TrimSpace(input.FirstName) == "":
		return errors.New("first_name is required")
	case strings.TrimSpace(input.LastName) == "":
		return errors.New("last_name is required")
	case strings.TrimSpace(input.Email) == "":
		return errors.New("email is required")
	case strings.TrimSpace(input.Number) == "":
		return errors.New("number is required")
	case strings.TrimSpace(input.Password) == "":
		return errors.New("password is required")
	}

	err := util.ValidateCNPJ(input.Cnpj)
	if err != nil {
		return err
	}
	if !util.IsPhoneValid(input.Number) {
		return fmt.Errorf("invalid phone number: %s", input.Number)
	}
	if !util.IsEmailValid(input.Email) {
		return fmt.Errorf("invalid email address: %s", input.Email)
	}

	return nil
}

func ValidateReadUserDTO(c *gin.Context) (*dto.ReadUserDTO, error) {
	var input dto.ReadUserDTO
	// Bind path param
	if err := c.ShouldBindUri(&input); err != nil {
		return nil, errors.New("email is required")
	}
	if !util.IsEmailValid(input.Email) {
		return nil, fmt.Errorf("invalid email address: %s", input.Email)
	}

	return &input, nil
}

// ValidateReadUsersDTO binds optional query parameters for reading users
func ValidateReadUsersDTO(c *gin.Context) *dto.ReadUsersDTO {
	var input dto.ReadUsersDTO
	_ = c.ShouldBindQuery(&input)
	return &input
}
