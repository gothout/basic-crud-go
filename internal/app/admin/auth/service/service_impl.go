package service

import (
	"basic-crud-go/internal/app/admin/auth/model"
	"context"
	"time"
)

type AuthService interface {
	LoginUser(ctx context.Context, email, senha string) (*model.UserIdentity, error)
	RefreshTokenUser(ctx context.Context, email, token string) bool
	LogoutUser(ctx context.Context, email, token string) bool
	GenerateTokenAPI(ctx context.Context, email, token string, expiresAt time.Time) (string, error)
}
