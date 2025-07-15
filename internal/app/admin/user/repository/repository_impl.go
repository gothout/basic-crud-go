package repository

import (
	"basic-crud-go/internal/app/admin/user/model"
	"basic-crud-go/internal/configuration/logger"
	"context"
	"database/sql"
	"fmt"
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

// ReadAll retrieves a paginated list of users with enterprise CNPJ.
func (r *userRepositoryImpl) ReadAll(ctx context.Context, page, limit int) ([]model.UserExtend, error) {
	var users []model.UserExtend
	offset := (page - 1) * limit

	query := `
		SELECT 
			u.id,
			u.number,
			u.first_name,
			u.last_name,
			u.email,
			u.password,
			u.created_at,
			u.updated_at,
			e.name,
			e.cnpj,
			e.active,
			e.created_at,
			e.updated_at
		FROM "user" u
		JOIN enterprise e ON u.enterprise_id = e.id
		ORDER BY u.id DESC
		LIMIT $1 OFFSET $2;
	`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		logger.Log(logger.Error, module, "ReadAll", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.UserExtend
		if err := rows.Scan(
			&user.Id,
			&user.Number,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.Enterprise.Name,
			&user.Enterprise.Cnpj,
			&user.Enterprise.Active,
			&user.Enterprise.CreateAt,
			&user.Enterprise.UpdateAt,
		); err != nil {
			logger.Log(logger.Error, module, "ReadAll", err)
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		logger.Log(logger.Error, module, "ReadAll", err)
		return nil, err
	}

	return users, nil
}

// Read user by email
func (r *userRepositoryImpl) Read(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	query := `
		SELECT id, enterprise_id, number, first_name, last_name, email, password, created_at, updated_at
		FROM "user"
		WHERE email = $1
		LIMIT 1;
	`
	err := r.db.QueryRowContext(ctx, query, email).Scan(
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
		if err == sql.ErrNoRows {
			return &model.User{}, nil
		}
		logger.Log(logger.Error, module, "Read", err)
		return &model.User{}, err
	}
	return &user, nil
}

// Update user by Id
func (r *userRepositoryImpl) Update(ctx context.Context, user model.User) (*model.User, error) {
	query := `
		UPDATE "user"
		SET
			number = $1,
			first_name = $2,
			last_name = $3,
			email = $4,
			password = $5,
			updated_at = $6
		WHERE id = $7
		RETURNING id, enterprise_id, number, first_name, last_name, email, password, created_at, updated_at;
	`

	row := r.db.QueryRowContext(ctx, query,
		user.Number,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.UpdatedAt,
		user.Id,
	)

	var updatedUser model.User
	err := row.Scan(
		&updatedUser.Id,
		&updatedUser.EnterpriseId,
		&updatedUser.Number,
		&updatedUser.FirstName,
		&updatedUser.LastName,
		&updatedUser.Email,
		&updatedUser.Password,
		&updatedUser.CreatedAt,
		&updatedUser.UpdatedAt,
	)
	if err != nil {
		logger.Log(logger.Error, module, "Update", err)
		return nil, err
	}

	return &updatedUser, nil
}

func (r *userRepositoryImpl) Delete(ctx context.Context, id string) (bool, error) {
	query := `
		DELETE FROM "user" WHERE id = $1 RETURNING true;
	`

	var ok bool
	err := r.db.QueryRowContext(ctx, query, id).Scan(&ok)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, fmt.Errorf("user not found")
		}
		logger.Log(logger.Error, module, "Delete", err)
		return false, err
	}
	return true, nil
}
