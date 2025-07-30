package service

import (
	"basic-crud-go/internal/app/admin/auth/model"
	"context"
)

type AuthService interface {
	LoginUser(ctx context.Context, email, senha string) (*model.UserIdentity, error)
}
