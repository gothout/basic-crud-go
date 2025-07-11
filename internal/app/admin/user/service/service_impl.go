package service

import (
	entModel "basic-crud-go/internal/app/admin/enterprise/model"
	"basic-crud-go/internal/app/admin/enterprise/service"
	"basic-crud-go/internal/app/admin/user/model"
	"basic-crud-go/internal/app/admin/user/repository"
	util "basic-crud-go/internal/app/util/password"
	"basic-crud-go/internal/configuration/logger"
	"context"
	"fmt"
)

const module = "User-Service"

type userService struct {
	repo              repository.UserRepository
	enterpriseService service.EnterpriseService
}

func NewUserService(r repository.UserRepository, entService service.EnterpriseService) UserService {
	return &userService{
		repo:              r,
		enterpriseService: entService,
	}
}

func (s *userService) Ping(ctx context.Context) (string, error) {
	return "pong", nil
}
func (s *userService) Create(ctx context.Context, enterpriseCnpj, number, firstName, lastName, email, password string) (*model.User, error) {

	// Validate CNPJ
	enterprise, err := s.enterpriseService.Read(ctx, enterpriseCnpj)
	if err != nil {
		return nil, err
	}

	passwordHash, err := util.Hash(password)
	if err != nil {
		logger.Log(logger.Error, module, "Create", err)
		return nil, err
	}

	// Create user
	user, err := s.repo.Create(ctx, enterprise.Id, number, firstName, lastName, email, passwordHash)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return user, nil
}

func (s *userService) Read(ctx context.Context, email string) (*model.User, *entModel.Enterprise, error) {
	user, err := s.repo.Read(ctx, email)

	if err != nil {
		return nil, nil, err
	}
	enterprise, err := s.enterpriseService.ReadById(ctx, user.EnterpriseId)
	if err != nil {
		return nil, nil, err
	}

	return user, enterprise, nil
}
