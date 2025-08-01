package repository

import (
	"basic-crud-go/internal/app/admin/user/model"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, enterpriseId int64, number, firstName, lastName, email, password string) (*model.User, error)
	ReadAll(ctx context.Context, page, limit int) ([]model.UserExtend, error)
	ReadUsersByEnterpriseID(ctx context.Context, entepriseId int64, page, limit int) ([]model.UserExtend, error)
	Read(ctx context.Context, email string) (*model.User, error)
	ReadById(ctx context.Context, id string) (*model.User, error)
	Update(ctx context.Context, user model.User) (*model.User, error)
	Delete(ctx context.Context, id string) (bool, error)
}
