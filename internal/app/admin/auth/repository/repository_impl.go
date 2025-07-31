package repository

import (
	"basic-crud-go/internal/app/admin/auth/model"
	"context"
	"time"
)

type AuthRepository interface {
	//GetCredencial(userId string) (model.User, error)
	GenerateTokenUser(ctx context.Context, userId string, createdAt time.Time) (*model.TokenUser, error)
	GenerateTokenAPI(ctx context.Context, userId string, createdAt, ExpiresAt time.Time) (string, error)
}
