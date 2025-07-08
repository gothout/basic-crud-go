package binding

import (
	"basic-crud-go/internal/app/admin/enterprise/dto"
	"basic-crud-go/internal/app/admin/enterprise/util"
	"errors"
	"strings"
)

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
