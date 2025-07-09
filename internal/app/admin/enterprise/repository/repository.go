package repository

import (
	"basic-crud-go/internal/app/admin/enterprise/model"
	"context"
)

type EnterpriseRepository interface {
	CreateEnterpriseByCNPJ(ctx context.Context, name, cnpj string) (int64, error)
	ReadAllEnterprise(ctx context.Context, page, limit int) ([]model.Enterprise, error)
	ReadEnterpriseByCNPJ(ctx context.Context, cnpj string) (model.Enterprise, error)
	Update(ctx context.Context, id int64, newCnpj, newName string) (model.Enterprise, error)
}
