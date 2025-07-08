package repository

import (
	"basic-crud-go/internal/app/admin/enterprise/model"
	"context"
)

type EnterpriseRepository interface {
	CreateEnterpriseByCNPJ(ctx context.Context, name, cnpj string) (int64, error)
	ReadEnterpriseByCNPJ(ctx context.Context, cnpj string) (model.Enterprise, error)
}
