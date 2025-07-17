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

func (r *authRepositoryImpl) CreateToken(ctx context.Context, userId string) (*model.Token, error) {
	// generate secure token
	token, err := util.GenerateSecureToken(32)
	if err != nil {
		logger.Log(logger.Error, module, "CreateToken", fmt.Errorf("error generating token %w", err))
		return nil, fmt.Errorf("error generating token")
	}
	createdAt := time.Now()
	query := `
		INSERT INTO admin_token (user_id, token)
		VALUES ($1, $2)
		RETURNING id, user_id, token, created_at;
	`
	var result model.Token
	err = r.db.QueryRowContext(ctx, query, userId, token, createdAt).Scan(&result.Id, &result.UserId, &result.Token, &result.CreatedAt)
	return &result, nil
}

func (r *authRepositoryImpl) CreateUser(ctx context.Context, userId, permission string) error {
	query := `
		INSERT INTO user_permission (user_id, permission_id)
		SELECT $1, ap.id
		FROM admin_permission ap
		JOIN admin_module am ON am.id = ap.module_id
		WHERE am.name = $2
		ON CONFLICT (user_id, permission_id) DO NOTHING;
	`

	_, err := r.db.ExecContext(ctx, query, userId, permission)
	if err != nil {
		logger.Log(logger.Error, module, "CreateUser", fmt.Errorf("error assigning permissions to user: %w", err))
		return fmt.Errorf("failed to assign permissions to user")
	}

	return nil
}
