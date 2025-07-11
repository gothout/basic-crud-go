package binding

import (
	"basic-crud-go/internal/app/admin/user/dto"
	"basic-crud-go/internal/app/admin/user/util"
	"errors"
	"fmt"
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
