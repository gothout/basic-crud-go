package repository

import (
	"basic-crud-go/internal/app/admin/user/model"
	"basic-crud-go/internal/configuration/logger"
	"context"
	"database/sql"
)

const module string = "User-Repository"

type userRepositoryImpl struct {
	db *sql.DB
}

// NewUserRepositoryImpl create new instance an repository user
func NewUserRepositoryImpl(db *sql.DB) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

// Create user
func (r *userRepositoryImpl) Create(ctx context.Context, enterpriseId int64, number, firstName, lastName, email, password string) (*model.User, error) {
	query := `
		INSERT INTO "user" (
			enterprise_id,
			number,
			first_name,
			last_name,
			email,
			password
		) VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, enterprise_id, number, first_name, last_name, email, password, created_at, updated_at;
	`

	row := r.db.QueryRowContext(ctx, query,
		enterpriseId,
		number,
		firstName,
		lastName,
		email,
		password,
	)

	var user model.User
	err := row.Scan(
		&user.Id,
		&user.EnterpriseId,
		&user.Number,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		logger.Log(logger.Error, module, "Create", err)
		return nil, err
	}

	return &user, nil
}
