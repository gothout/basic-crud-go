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
		logger.Log(logger.Error, module, "GenerateTokenUser", fmt.Errorf("error generating token %w", err))
		return nil, fmt.Errorf("error generating token")
	}

	query := `
		INSERT INTO admin_user_token (user_id, token, created_at)
		VALUES ($1, $2, $3)
		RETURNING id, user_id, token, created_at;
	`

	var result model.TokenUser
	err = r.db.QueryRowContext(ctx, query, userId, token, createdAt).
		Scan(&result.Id, &result.UserId, &result.Token, &result.CreatedAt)
	if err != nil {
		logger.Log(logger.Error, module, "GenerateTokenUser", fmt.Errorf("error inserting token %w", err))
		return nil, fmt.Errorf("error inserting token")
	}

	logger.Log(logger.Info, module, "GenerateTokenUser", fmt.Sprintf("user token created for user_id=%s at %s", result.UserId, result.CreatedAt))
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
		logger.Log(logger.Error, module, "GenerateTokenAPI", fmt.Errorf("error inserting API token %w", err))
		return "", fmt.Errorf("error inserting token")
	}

	logger.Log(logger.Info, module, "GenerateTokenAPI", fmt.Sprintf("API token created for user_id=%s expiring at %s", userId, expiresAt))
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
			logger.Log(logger.Warning, module, "GetUserIdByAPIKey", fmt.Errorf("api key not found"))
			return "", fmt.Errorf("api key not found")
		}
		logger.Log(logger.Error, module, "GetUserIdByAPIKey", fmt.Errorf("failed to query api key %w", err))
		return "", fmt.Errorf("failed to query api key: %w", err)
	}

	logger.Log(logger.Info, module, "GetUserIdByAPIKey", fmt.Sprintf("api key valid for user_id=%s", userId))
	return userId, nil
}

// GetValidUserIdByToken returns the userId if the token exists
// and has not expired (1h since createdAt based on backend clock).
// Returns ("", nil) if not found or expired.
func (r *authRepositoryImpl) GetValidUserIdByToken(ctx context.Context, token string) (string, error) {
	const q = `
		SELECT user_id, created_at
		FROM admin_user_token
		WHERE token = $1
		LIMIT 1;
	`

	var userId string
	var createdAt time.Time

	if err := r.db.QueryRowContext(ctx, q, token).Scan(&userId, &createdAt); err != nil {
		if err == sql.ErrNoRows {
			logger.Log(logger.Warning, module, "GetValidUserIdByToken", fmt.Errorf("token not found"))
			return "", nil
		}
		logger.Log(logger.Warning, module, "GetValidUserIdByToken", fmt.Errorf("failed to query token %w", err))
		return "", fmt.Errorf("failed to query user token: %w", err)
	}

	// Interpret DB timestamp (no timezone) as LOCAL time, not UTC
	createdLocal := time.Date(
		createdAt.Year(), createdAt.Month(), createdAt.Day(),
		createdAt.Hour(), createdAt.Minute(), createdAt.Second(), createdAt.Nanosecond(),
		time.Local,
	)

	// Compare using backend clock (local)
	if time.Since(createdLocal) > time.Hour {
		logger.Log(logger.Info, module, "GetValidUserIdByToken",
			fmt.Sprintf("token expired for user_id=%s created_at=%s (local=%s)", userId, createdAt, createdLocal))
		return "", nil
	}

	// logger.Log(logger.Info, module, "GetValidUserIdByToken",
	// 	fmt.Sprintf("token valid for user_id=%s created_at_local=%s", userId, createdLocal))
	return userId, nil
}
