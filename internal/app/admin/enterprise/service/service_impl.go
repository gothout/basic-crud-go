package service

import (
	"basic-crud-go/internal/app/admin/enterprise/model"
	"basic-crud-go/internal/app/admin/enterprise/repository"
	"basic-crud-go/internal/configuration/logger"
	"basic-crud-go/internal/infrastructure/db/postgres"
	"context"
	"fmt"
	"time"
)

const module string = "Enterprise-Service"

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

// ReadAllEnterprise retrieves a paginated list of enterprises.
func (s *enterpriseServiceImpl) ReadAllEnterprise(ctx context.Context, page, limit int) ([]model.Enterprise, error) {
	var enterprises []model.Enterprise

	if page < 1 {
		logger.LogWithAutoFuncName(logger.Info, module, "page out of range. Defaulting to 1.")
		page = 1
	}

	if limit <= 0 || limit > 10 {
		logger.LogWithAutoFuncName(logger.Info, module, "limit out of range. Defaulting to 10.")
		limit = 10
	}

	enterprises, err := s.repo.ReadAllEnterprise(ctx, page, limit)
	if err != nil {
		return nil, err
	}

	return enterprises, nil
}

// Read enteprise by cnpj
func (s *enterpriseServiceImpl) Read(ctx context.Context, cnpj string) (model.Enterprise, error) {
	Enterprise, err := s.repo.ReadEnterpriseByCNPJ(ctx, cnpj)
	if err != nil {
		return model.Enterprise{}, fmt.Errorf("erro ao ler empresa")
	}
	if Enterprise.Name == "" {
		return model.Enterprise{}, fmt.Errorf("empresa nao encontrada")
	}

	return Enterprise, nil
}
