package service

import (
	"basic-crud-go/internal/infrastructure/db/postgres"
	"context"
	"database/sql"
)

type enterpriseService struct {
	db *sql.DB
}

func NewEnterpriseService() EnterpriseService {
	return &enterpriseService{
		db: postgres.GetDB(), // busca a instância do banco diretamente
	}
}

func (s *enterpriseService) Ping(ctx context.Context) (string, error) {
	return "pong", nil
}
