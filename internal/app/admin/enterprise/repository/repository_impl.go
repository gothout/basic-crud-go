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
func (r *enterpriseRepositoryImpl) CreateEnterpriseByCNPJ(ctx context.Context, name, cnpj string) (int64, error) {
	query := `
		INSERT INTO enterprise (name, cnpj, active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
	`

	now := time.Now()

	var id int64
	err := r.db.QueryRowContext(ctx, query, name, cnpj, false, now, now).Scan(&id)
	if err != nil {
		logger.Log(logger.Error, module, "CreateEnterpriseByCNPJ", err)
		return 0, err
	}

	return id, nil
}

// ReadAllEnterprise retrieves a paginated list of enterprises.
func (r *enterpriseRepositoryImpl) ReadAllEnterprise(ctx context.Context, page, limit int) ([]model.Enterprise, error) {
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
		logger.Log(logger.Error, module, "ReadAllEnterprise", err)
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
			logger.Log(logger.Error, module, "ReadAllEnterprise", err)
			return nil, err
		}
		enterprises = append(enterprises, ent)
	}

	if err := rows.Err(); err != nil {
		logger.Log(logger.Error, module, "ReadAllEnterprise", err)
		return nil, err
	}

	return enterprises, nil
}

// Read enteprise by CNPJ value
func (r *enterpriseRepositoryImpl) ReadEnterpriseByCNPJ(ctx context.Context, cnpj string) (model.Enterprise, error) {
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
		logger.Log(logger.Error, module, "ReadEnterpriseByCNPJ", err)
		return model.Enterprise{}, err
	}

	return enterprise, nil
}
