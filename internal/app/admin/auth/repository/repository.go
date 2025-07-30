package repository

import (
	"basic-crud-go/internal/app/admin/auth/model"
	"basic-crud-go/internal/app/admin/auth/util"
	"basic-crud-go/internal/configuration/logger"
	"context"
	"database/sql"
	"fmt"
	"time"
)

const module string = "Admin-Auth-Repository"

type authRepositoryImpl struct {
	db *sql.DB
}

func NewAuthRepositoryImpl(db *sql.DB) AuthRepository {
	return &authRepositoryImpl{
		db: db,
	}
}

func (r *authRepositoryImpl) GenerateTokenUser(ctx context.Context, userId string, createdAt time.Time) (*model.TokenUser, error) {
	// generate secure token
	token, err := util.GenerateSecureToken(32)
	if err != nil {
		logger.Log(logger.Error, module, "CreateToken", fmt.Errorf("error generating token %w", err))
		return nil, fmt.Errorf("error generating token")
	}

	query := `
		INSERT INTO admin_user_token (user_id, token, created_at)
		VALUES ($1, $2, $3)
		RETURNING id, user_id, token, created_at;
	`
	var result model.TokenUser
	err = r.db.QueryRowContext(ctx, query, userId, token, createdAt).Scan(&result.Id, &result.UserId, &result.Token, &result.CreatedAt)
	if err != nil {
		logger.Log(logger.Error, module, "InsertToken", fmt.Errorf("error inserting token %w", err))
		return nil, fmt.Errorf("error inserting token")
	}

	return &result, nil
}
