package service

import (
	"basic-crud-go/internal/app/admin/enterprise/model"
	"basic-crud-go/internal/app/admin/enterprise/repository"
	"basic-crud-go/internal/infrastructure/db/postgres"
	"context"
	"fmt"
	"time"
)

type enterpriseServiceImpl struct {
	repo repository.EnterpriseRepository
}

func NewEnterpriseService() EnterpriseService {
	return &enterpriseServiceImpl{
		repo: repository.NewRepositoryImpl(postgres.GetDB()),
	}
}

func (s *enterpriseServiceImpl) Ping(ctx context.Context) (string, error) {
	return "pong", nil
}

// Create enterprise
func (s *enterpriseServiceImpl) Create(ctx context.Context, name, cnpj string) (model.Enterprise, error) {
	var Enterprise model.Enterprise
	_, err := s.repo.CreateEnterpriseByCNPJ(ctx, name, cnpj)
	if err != nil {
		return Enterprise, fmt.Errorf("erro ao criar empresa")
	}

	Enterprise = model.Enterprise{
		Name:     name,
		Cnpj:     cnpj,
		Active:   true,
		CreateAt: time.Now(),
	}

	return Enterprise, nil
}
