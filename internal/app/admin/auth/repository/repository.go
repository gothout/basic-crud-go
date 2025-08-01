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

func (r *authRepositoryImpl) GenerateTokenAPI(ctx context.Context, userId string, createdAt, expiresAt time.Time) (string, error) {
	// generate secure token
	token, err := util.GenerateSecureToken(32)
	if err != nil {
		logger.Log(logger.Error, module, "GenerateTokenAPI", fmt.Errorf("error generating token %w", err))
		return "", fmt.Errorf("error generating token")
	}
	tokenApi := "api_" + token
	query := `
		INSERT INTO admin_api_token (user_id, token, created_at, end_date)
		VALUES ($1, $2, $3, $4);
	`

	_, err = r.db.ExecContext(ctx, query, userId, tokenApi, createdAt, expiresAt)
	if err != nil {
		logger.Log(logger.Error, module, "InsertTokenAPI", fmt.Errorf("error inserting API token %w", err))
		return "", fmt.Errorf("error inserting token")
	}

	return tokenApi, nil
}

func (r *authRepositoryImpl) GetUserIdByAPIKey(ctx context.Context, apiKey string) (string, error) {
	const query = `
		SELECT user_id
		FROM admin_api_token
		WHERE token = $1
	`

	var userId string
	err := r.db.QueryRowContext(ctx, query, apiKey).Scan(&userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("api key not found")
		}
		return "", fmt.Errorf("failed to query api key: %w", err)
	}

	return userId, nil
}
