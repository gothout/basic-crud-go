package repository

import (
	"basic-crud-go/internal/app/admin/user/model"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, enterpriseId int64, number, firstName, lastName, email, password string) (*model.User, error)
	Read(ctx context.Context, email string) (*model.User, error)
}
