package repository

import (
	"basic-crud-go/internal/app/admin/permission/model"
	"basic-crud-go/internal/configuration/logger"
	"context"
	"database/sql"
	"fmt"
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

func (r *permissionRepositoryImpl) Read(ctx context.Context, moduleID int64) (*model.ModulePermission, error) {
	query := `
		SELECT am.id, am.name, aa.id, aa.name
		FROM admin_permission ap
		JOIN admin_module am ON am.id = ap.module_id
		JOIN admin_action aa ON aa.id = ap.action_id
		WHERE am.id = $1
	`

	rows, err := r.db.QueryContext(ctx, query, moduleID)
	if err != nil {
		logger.Log(logger.Error, module, "Read", fmt.Errorf("error querying permissions: %w", err))
		return nil, fmt.Errorf("failed to query permissions")
	}
	defer rows.Close()

	var modulePermission model.ModulePermission

	modulePermission.Actions = []model.PermissionAction{}

	for rows.Next() {
		var action model.PermissionAction
		err := rows.Scan(&modulePermission.ID, &modulePermission.Name, &action.ID, &action.Name)
		if err != nil {
			logger.Log(logger.Error, module, "Read", fmt.Errorf("error scanning row: %w", err))
			return nil, fmt.Errorf("failed to parse permission result")
		}
		modulePermission.Actions = append(modulePermission.Actions, action)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error in result set: %w", err)
	}

	return &modulePermission, nil
}

func (r *permissionRepositoryImpl) ReadModuleByName(ctx context.Context, name string) (*model.ModulePermission, error) {
	query := `
		SELECT id, name
		FROM admin_module
		WHERE name = $1
		LIMIT 1;
	`

	var modulePermission model.ModulePermission

	err := r.db.QueryRowContext(ctx, query, name).Scan(&modulePermission.ID, &modulePermission.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("module with name '%s' not found", name)
		}
		logger.Log(logger.Error, module, "ReadModuleByName", err)
		return nil, err
	}

	return &modulePermission, nil
}

func (r *permissionRepositoryImpl) ReadAllModules(ctx context.Context, page, limit int) ([]model.ModulePermission, error) {
	var modules []model.ModulePermission
	offset := (page - 1) * limit

	query := `
		SELECT id, name
		FROM admin_module
		ORDER BY id ASC
		LIMIT $1 OFFSET $2;
	`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		logger.Log(logger.Error, module, "ReadAllModules", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var mod model.ModulePermission
		if err := rows.Scan(&mod.ID, &mod.Name); err != nil {
			logger.Log(logger.Error, module, "ReadAllModules", err)
			return nil, err
		}
		modules = append(modules, mod)
	}

	if err := rows.Err(); err != nil {
		logger.Log(logger.Error, module, "ReadAllModules", err)
		return nil, err
	}

	return modules, nil
}
