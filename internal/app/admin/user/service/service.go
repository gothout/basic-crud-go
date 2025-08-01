package service

import (
	entModel "basic-crud-go/internal/app/admin/enterprise/model"
	"basic-crud-go/internal/app/admin/user/dto"
	"basic-crud-go/internal/app/admin/user/model"
	"context"
)

type UserService interface {
	Ping(ctx context.Context) (string, error)
	Create(ctx context.Context, enterpriseCnpj, number, firstName, lastName, email, password string) (*model.User, error)
	ReadAll(ctx context.Context, page, limit int) ([]model.UserExtend, error)
	ReadByCnpj(ctx context.Context, cnpj string, page, limit int) ([]model.UserExtend, error)
	Read(ctx context.Context, email string) (*model.User, *entModel.Enterprise, error)
	ReadById(ctx context.Context, id string) (*model.User, *entModel.Enterprise, error)
	Update(ctx context.Context, dto dto.UpdateUserDTO) (*model.User, *entModel.Enterprise, error)
	Delete(ctx context.Context, email string) (bool, error)
}
