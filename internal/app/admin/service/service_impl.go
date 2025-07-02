package service

import (
	"basic-crud-go/internal/infrastructure/db/postgres"
	"context"
	"database/sql"
)

type adminService struct {
	db *sql.DB
}

func NewAdminService() AdminService {
	return &adminService{
		db: postgres.GetDB(), // busca a instância do banco diretamente
	}
}

func (s *adminService) Ping(ctx context.Context) (string, error) {
	return "pong", nil
}
