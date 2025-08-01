package binding

import (
	"basic-crud-go/internal/app/admin/enterprise/dto"
	"basic-crud-go/internal/app/admin/enterprise/util"
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

// Validate request create enteprise
func ValidateCreateEnterpriseDTO(input dto.CreateEnterpriseDTO) error {
	if strings.TrimSpace(input.Name) == "" {
		return errors.New("name is required")
	}
	if strings.TrimSpace(input.Cnpj) == "" {
		return errors.New("cnpj is required")
	}
	if err := util.ValidateCNPJ(input.Cnpj); err != nil {
		return err
	}
	return nil
}

// Validate request read enterprise
func ValidateReadEnterpriseDTO(c *gin.Context) (*dto.ReadEnterpriseDTO, error) {
	var input dto.ReadEnterpriseDTO

	// Bind path param
	if err := c.ShouldBindUri(&input); err != nil {
		return nil, errors.New("cnpj is required")
	}

	// clean cnpj
	input.Cnpj = util.RemoveNonDigits(input.Cnpj)

	// validate cnpj
	if strings.TrimSpace(input.Cnpj) == "" {
		return nil, errors.New("cnpj is required")
	}
	if err := util.ValidateCNPJ(input.Cnpj); err != nil {
		return nil, err
	}

	return &input, nil
}

// ValidateReadEnterprisesDTO binds optional query parameters for reading enterprises
func ValidateReadEnterprisesDTO(c *gin.Context) *dto.ReadEnterprisesDTO {
	var input dto.ReadEnterprisesDTO
	_ = c.ShouldBindQuery(&input) // Ignore error, as fields are optional
	return &input
}

// Validate request update enterprise by cnpj
func ValidateUpdateEnterpriseDTO(input dto.UpdateEnterpriseDTO) error {
	if strings.TrimSpace(input.Cnpj) == "" {
		return errors.New("cnpj is required")
	}
	if strings.TrimSpace(input.NewName) == "" && strings.TrimSpace(input.NewCnpj) == "" {
		return errors.New("newName or newCnpj is required")
	}
	if err := util.ValidateCNPJ(util.RemoveNonDigits(input.Cnpj)); err != nil {
		return err
	}
	if input.NewCnpj != "" {
		if err := util.ValidateCNPJ(util.RemoveNonDigits(input.NewCnpj)); err != nil {
			return err
		}
	}
	return nil
}

// Validate request delete enterprise by cnpj
func ValidateDeleteEnterpriseDTO(c *gin.Context) (*dto.DeleteEnterpriseDTO, error) {
	var input dto.DeleteEnterpriseDTO
	// Bind path param
	if err := c.ShouldBindUri(&input); err != nil {
		return nil, errors.New("cnpj is required")
	}

	// clean cnpj
	input.Cnpj = util.RemoveNonDigits(input.Cnpj)

	// validate cnpj
	if strings.TrimSpace(input.Cnpj) == "" {
		return nil, errors.New("cnpj is required")
	}
	if err := util.ValidateCNPJ(input.Cnpj); err != nil {
		return nil, err
	}

	return &input, nil
}
