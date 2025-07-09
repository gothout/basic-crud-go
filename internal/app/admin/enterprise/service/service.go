package service

import (
	"basic-crud-go/internal/app/admin/enterprise/model"
	"context"
)

type EnterpriseService interface {
	Ping(ctx context.Context) (string, error)
	Create(ctx context.Context, name, cnpj string) (model.Enterprise, error)
	Read(ctx context.Context, cnpj string) (model.Enterprise, error)
	ReadAllEnterprise(ctx context.Context, page, limit int) ([]model.Enterprise, error)
	Update(ctx context.Context, oldCnpj, newCnpj, newName string) (model.Enterprise, error)
}
