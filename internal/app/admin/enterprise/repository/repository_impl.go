package repository

import (
	"basic-crud-go/internal/app/admin/enterprise/model"
	"basic-crud-go/internal/configuration/logger"
	"context"
	"database/sql"
	"time"
)

const module string = "Enterprise-Repository"

type enterpriseRepositoryImpl struct {
	db *sql.DB
}

// NewRepositoryImpl cria uma nova instância do repositório de empresas.
func NewRepositoryImpl(db *sql.DB) EnterpriseRepository {
	return &enterpriseRepositoryImpl{
		db: db,
	}
}

// Create enterprise
func (r *enterpriseRepositoryImpl) Create(ctx context.Context, name, cnpj string) (int64, error) {
	query := `
		INSERT INTO enterprise (name, cnpj, active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
	`

	now := time.Now()

	var id int64
	err := r.db.QueryRowContext(ctx, query, name, cnpj, false, now, now).Scan(&id)
	if err != nil {
		logger.Log(logger.Error, module, "Create", err)
		return 0, err
	}

	return id, nil
}

// ReadAllEnterprise retrieves a paginated list of enterprises.
func (r *enterpriseRepositoryImpl) ReadAll(ctx context.Context, page, limit int) ([]model.Enterprise, error) {
	var enterprises []model.Enterprise

	offset := (page - 1) * limit

	query := `
		SELECT id, name, cnpj, active, created_at, updated_at
		FROM enterprise
		ORDER BY id DESC
		LIMIT $1 OFFSET $2;
	`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		logger.Log(logger.Error, module, "ReadAll", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var ent model.Enterprise
		if err := rows.Scan(
			&ent.Id,
			&ent.Name,
			&ent.Cnpj,
			&ent.Active,
			&ent.CreateAt,
			&ent.UpdateAt,
		); err != nil {
			logger.Log(logger.Error, module, "ReadAll", err)
			return nil, err
		}
		enterprises = append(enterprises, ent)
	}

	if err := rows.Err(); err != nil {
		logger.Log(logger.Error, module, "ReadAll", err)
		return nil, err
	}

	return enterprises, nil
}

// Read enteprise by CNPJ value
func (r *enterpriseRepositoryImpl) Read(ctx context.Context, cnpj string) (model.Enterprise, error) {
	var enterprise model.Enterprise

	query := `
		SELECT id, name, cnpj, active, created_at, updated_at 
		FROM enterprise 
		WHERE cnpj = $1;
	`
	err := r.db.QueryRowContext(ctx, query, cnpj).Scan(
		&enterprise.Id,
		&enterprise.Name,
		&enterprise.Cnpj,
		&enterprise.Active,
		&enterprise.CreateAt,
		&enterprise.UpdateAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return model.Enterprise{}, nil
		}
		logger.Log(logger.Error, module, "Read", err)
		return model.Enterprise{}, err
	}

	return enterprise, nil
}

func (r *enterpriseRepositoryImpl) ReadById(ctx context.Context, id int64) (*model.Enterprise, error) {
	var enterprise model.Enterprise

	query := `
		SELECT id, name, cnpj, active, created_at, updated_at 
		FROM enterprise 
		WHERE id = $1;
	`
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&enterprise.Id,
		&enterprise.Name,
		&enterprise.Cnpj,
		&enterprise.Active,
		&enterprise.CreateAt,
		&enterprise.UpdateAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return &model.Enterprise{}, nil
		}
		logger.Log(logger.Error, module, "Read", err)
		return &model.Enterprise{}, err
	}

	return &enterprise, nil
}

// Update enterprise by CNPJ value
func (r *enterpriseRepositoryImpl) Update(ctx context.Context, id int64, newCnpj, newName string) (model.Enterprise, error) {
	var updatedEnterprise model.Enterprise

	query := `
		UPDATE enterprise
		SET name = $1, cnpj = $2, updated_at = NOW()
		WHERE id = $3
		RETURNING id, name, cnpj, active, created_at, updated_at;
	`
	err := r.db.QueryRowContext(ctx, query, newName, newCnpj, id).Scan(
		&updatedEnterprise.Id,
		&updatedEnterprise.Name,
		&updatedEnterprise.Cnpj,
		&updatedEnterprise.Active,
		&updatedEnterprise.CreateAt,
		&updatedEnterprise.UpdateAt,
	)
	if err != nil {
		logger.Log(logger.Error, module, "Update", err)
		return model.Enterprise{}, err
	}

	return updatedEnterprise, nil
}

// Delete enterprise ByCNPJ
func (r *enterpriseRepositoryImpl) Delete(ctx context.Context, id int64) (bool, error) {
	query := `
		DELETE FROM enterprise WHERE id = $1 RETURNING true;
	`
	var ok bool
	err := r.db.QueryRowContext(ctx, query, id).Scan(&ok)
	if err != nil {
		logger.Log(logger.Error, module, "Delete", err)
		return false, err
	}
	return ok, nil
}
