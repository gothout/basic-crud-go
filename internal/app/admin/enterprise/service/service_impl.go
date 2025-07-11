package service

import (
	"basic-crud-go/internal/app/admin/enterprise/model"
	"basic-crud-go/internal/app/admin/enterprise/repository"
	"basic-crud-go/internal/configuration/logger"
	"context"
	"fmt"
	"strings"
	"time"
)

const module string = "Enterprise-Service"

type enterpriseServiceImpl struct {
	repo repository.EnterpriseRepository
}

func NewEnterpriseService(r repository.EnterpriseRepository) EnterpriseService {
	return &enterpriseServiceImpl{
		repo: r,
	}
}

func (s *enterpriseServiceImpl) Ping(ctx context.Context) (string, error) {
	return "pong", nil
}

// Create enterprise
func (s *enterpriseServiceImpl) Create(ctx context.Context, name, cnpj string) (model.Enterprise, error) {
	var Enterprise model.Enterprise
	_, err := s.repo.Create(ctx, name, cnpj)
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
func (s *enterpriseServiceImpl) ReadAll(ctx context.Context, page, limit int) ([]model.Enterprise, error) {
	var enterprises []model.Enterprise

	if page < 1 {
		logger.LogWithAutoFuncName(logger.Info, module, "page out of range. Defaulting to 1.")
		page = 1
	}

	if limit <= 0 || limit > 10 {
		logger.LogWithAutoFuncName(logger.Info, module, "limit out of range. Defaulting to 10.")
		limit = 10
	}

	enterprises, err := s.repo.ReadAll(ctx, page, limit)
	if err != nil {
		return nil, err
	}

	return enterprises, nil
}

// Read enteprise by cnpj
func (s *enterpriseServiceImpl) Read(ctx context.Context, cnpj string) (model.Enterprise, error) {
	Enterprise, err := s.repo.Read(ctx, cnpj)
	if err != nil {
		return model.Enterprise{}, fmt.Errorf("erro ao ler empresa")
	}
	if Enterprise.Name == "" {
		return model.Enterprise{}, fmt.Errorf("empresa nao encontrada")
	}

	return Enterprise, nil
}

// Read enteprise by Id
func (s *enterpriseServiceImpl) ReadById(ctx context.Context, id int64) (*model.Enterprise, error) {
	Enterprise, err := s.repo.ReadById(ctx, id)
	if err != nil {
		return &model.Enterprise{}, fmt.Errorf("erro ao ler empresa")
	}
	if Enterprise.Name == "" {
		return &model.Enterprise{}, fmt.Errorf("empresa nao encontrada")
	}

	return Enterprise, nil
}

// Update enterprise by cnpj
func (s *enterpriseServiceImpl) Update(ctx context.Context, oldCnpj, newCnpj, newName string) (model.Enterprise, error) {

	enterprise, err := s.Read(ctx, oldCnpj)
	if err != nil {
		return model.Enterprise{}, fmt.Errorf("enterprise not found")
	}

	if newCnpj == "" {
		newCnpj = enterprise.Cnpj
	}
	if newName == "" {
		newName = enterprise.Name
	}

	updatedEnterprise, err := s.repo.Update(ctx, enterprise.Id, newCnpj, newName)
	if err != nil {
		if strings.Contains(err.Error(), "enterprise_cnpj_key") {
			return model.Enterprise{}, fmt.Errorf("enterprise_cnpj_key")
		}
		return model.Enterprise{}, fmt.Errorf("enterprise not found")
	}

	return updatedEnterprise, nil
}

// Delete enterprise by cnpj
func (s *enterpriseServiceImpl) Delete(ctx context.Context, cnpj string) (bool, error) {
	enterprise, err := s.Read(ctx, cnpj)
	if err != nil {
		return false, fmt.Errorf("enterprise not found")
	}
	deleted, err := s.repo.Delete(ctx, enterprise.Id)
	if err != nil {
		return deleted, fmt.Errorf("error deleting enterprise")
	}

	return deleted, nil
}
