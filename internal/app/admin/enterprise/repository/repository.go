package repository

import "context"

type EnterpriseRepository interface {
	CreateEnterpriseByCNPJ(ctx context.Context, name, cnpj string) (int64, error)
}
