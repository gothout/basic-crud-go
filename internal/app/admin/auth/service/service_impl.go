package service

import (
	"basic-crud-go/internal/app/admin/auth/model"
	"context"
)

type AuthService interface {
	LoginUser(ctx context.Context, email, senha string) (*model.UserIdentity, error)
	RefreshTokenUser(ctx context.Context, email, password, token string) bool
	LogoutUser(ctx context.Context, email string)
}
