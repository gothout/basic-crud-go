package service

import (
	"basic-crud-go/internal/infrastructure/db/postgres"
	"context"
	"database/sql"
)

type userService struct {
	db *sql.DB
}

func NewUserService() UserService {
	return &userService{
		db: postgres.GetDB(), // busca a instância do banco diretamente
	}
}

func (s *userService) Ping(ctx context.Context) (string, error) {
	return "pong", nil
}
