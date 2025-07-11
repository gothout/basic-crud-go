package service

import (
	"basic-crud-go/internal/app/admin/user/model"
	"context"
)

type UserService interface {
	Ping(ctx context.Context) (string, error)
	Create(ctx context.Context, enterpriseCnpj, number, firstName, lastName, email, password string) (*model.User, error)
}
