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

// ValidateUpdateUserDTO binds optional query parameters for update user
func ValidateUpdateUserDTO(c *gin.Context) (*dto.UpdateUserDTO, error) {
	var input dto.UpdateUserDTO

	// Bind path param (email)
	if err := c.ShouldBindUri(&input); err != nil {
		return nil, errors.New("email (from URI) is required")
	}

	// Bind body JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		return nil, errors.New("invalid or missing JSON body")
	}

	// Validate email from URI
	if !util.IsEmailValid(input.Email) {
		return nil, fmt.Errorf("invalid email address (uri): %s", input.Email)
	}

	// Validate updated email (if provided)
	if input.EmailUpdated != "" && !util.IsEmailValid(input.EmailUpdated) {
		return nil, fmt.Errorf("invalid email address (new): %s", input.EmailUpdated)
	}

	// Validate phone (if provided)
	if input.Number != "" && !util.IsPhoneValid(input.Number) {
		return nil, fmt.Errorf("invalid phone number: %s", input.Number)
	}

	// sanit number
	input.Number = util.RemoveNonDigits(input.Number)

	return &input, nil
}
func ValidateDeleteUserDTO(c *gin.Context) (*dto.DeleteUserDTO, error) {
	var input dto.DeleteUserDTO
	// Bind path param
	if err := c.ShouldBindUri(&input); err != nil {
		return nil, errors.New("email is required")
	}
	if !util.IsEmailValid(input.Email) {
		return nil, fmt.Errorf("invalid email address: %s", input.Email)
	}

	return &input, nil
}

func ValidateReadUsersByCnpjDTO(c *gin.Context) (*dto.ReadUsersByCnpjDTO, error) {
	var input dto.ReadUsersByCnpjDTO

	if err := c.ShouldBindQuery(&input); err != nil {
		return nil, errors.New("invalid parameters in query")
	}

	if input.Cnpj == "" {
		return nil, errors.New("cnpj is required")
	}

	if err := util.ValidateCNPJ(input.Cnpj); err != nil {
		return nil, err
	}
	input.Cnpj = util.RemoveNonDigits(input.Cnpj)

	return &input, nil
}
