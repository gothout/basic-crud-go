package repository

import (
	"basic-crud-go/internal/app/admin/enterprise/model"
	"context"
	"database/sql"
	"time"
)

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
		return 0, err
	}

	return id, nil
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
		return model.Enterprise{}, err
	}

	return enterprise, nil
}
