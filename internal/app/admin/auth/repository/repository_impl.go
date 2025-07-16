package repository

import (
	"basic-crud-go/internal/app/admin/auth/model"
	"context"
)

type AuthRepository interface {
	CreateToken(ctx context.Context, userId string) (*model.Token, error)
}
