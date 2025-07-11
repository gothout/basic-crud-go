package service

import (
	entModel "basic-crud-go/internal/app/admin/enterprise/model"
	"basic-crud-go/internal/app/admin/user/model"
	"context"
)

type UserService interface {
	Ping(ctx context.Context) (string, error)
	Create(ctx context.Context, enterpriseCnpj, number, firstName, lastName, email, password string) (*model.User, error)
	Read(ctx context.Context, email string) (*model.User, *entModel.Enterprise, error)
}
