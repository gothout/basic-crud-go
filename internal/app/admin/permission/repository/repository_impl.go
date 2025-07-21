package repository

import (
	"basic-crud-go/internal/app/admin/permission/model"
	"basic-crud-go/internal/configuration/logger"
	"context"
	"database/sql"
	"fmt"
	"github.com/lib/pq"
)

const module string = "Admin-Permission-Repository"

type permissionRepositoryImpl struct {
	db *sql.DB
}

func NewRepositoryImpl(db *sql.DB) PermissionRepository {
	return &permissionRepositoryImpl{
		db: db,
	}
}

func (r *permissionRepositoryImpl) ApplyPermissionUserBatch(ctx context.Context, userID string, codes []string) error {
	if len(codes) == 0 {
		return nil
	}

	query := `
		INSERT INTO user_permission (user_id, permission_id)
		SELECT $1, id FROM admin_permission
		WHERE code = ANY($2)
		ON CONFLICT DO NOTHING;
	`

	_, err := r.db.ExecContext(ctx, query, userID, pq.Array(codes))
	if err != nil {
		logger.Log(logger.Error, module, "ApplyPermissionUserBatch", err)
		return err
	}

	return nil
}

func (r *permissionRepositoryImpl) ReadAllPermissions(ctx context.Context, page, limit int) ([]model.Permission, error) {
	var permissions []model.Permission
	offset := (page - 1) * limit

	query := `
		SELECT id, code, description
		FROM admin_permission
		ORDER BY id ASC
		LIMIT $1 OFFSET $2;
	`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		logger.Log(logger.Error, "permission", "ReadAllPermissions", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var perm model.Permission
		if err := rows.Scan(&perm.ID, &perm.Code, &perm.Description); err != nil {
			logger.Log(logger.Error, "permission", "ReadAllPermissions", err)
			return nil, err
		}
		permissions = append(permissions, perm)
	}

	if err := rows.Err(); err != nil {
		logger.Log(logger.Error, "permission", "ReadAllPermissions", err)
		return nil, err
	}

	return permissions, nil
}

func (r *permissionRepositoryImpl) Search(ctx context.Context, name string) ([]model.Permission, error) {
	var permissions []model.Permission

	query := `
		SELECT id, code, description
		FROM admin_permission
		WHERE code ILIKE '%' || $1 || '%'
		ORDER BY id ASC;
	`

	rows, err := r.db.QueryContext(ctx, query, name)
	if err != nil {
		logger.Log(logger.Error, module, "ReadByName", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var perm model.Permission
		if err := rows.Scan(&perm.ID, &perm.Code, &perm.Description); err != nil {
			logger.Log(logger.Error, module, "ReadByName", err)
			return nil, err
		}
		permissions = append(permissions, perm)
	}

	if err := rows.Err(); err != nil {
		logger.Log(logger.Error, module, "ReadByName", err)
		return nil, err
	}

	return permissions, nil
}

func (r *permissionRepositoryImpl) ReadByCode(ctx context.Context, code string) (*model.Permission, error) {
	var perm model.Permission

	query := `
		SELECT id, code, description
		FROM admin_permission
		WHERE code = $1
		LIMIT 1;
	`

	err := r.db.QueryRowContext(ctx, query, code).Scan(
		&perm.ID,
		&perm.Code,
		&perm.Description,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			//logger.Log(logger.Error, module, "ReadByCode", err)
			return &model.Permission{}, fmt.Errorf("code not found")
		}
		logger.Log(logger.Error, module, "ReadByCode", err)
		return &model.Permission{}, fmt.Errorf("code not found")
	}
	return &perm, nil
}
