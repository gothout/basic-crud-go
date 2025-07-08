package repository

import (
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
