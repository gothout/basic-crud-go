package repository

import (
	"basic-crud-go/internal/app/admin/enterprise/model"
	"context"
)

type EnterpriseRepository interface {
	Create(ctx context.Context, name, cnpj string) (int64, error)
	ReadAll(ctx context.Context, page, limit int) ([]model.Enterprise, error)
	Read(ctx context.Context, cnpj string) (model.Enterprise, error)
	Update(ctx context.Context, id int64, newCnpj, newName string) (model.Enterprise, error)
	Delete(ctx context.Context, id int64) (bool, error)
}
