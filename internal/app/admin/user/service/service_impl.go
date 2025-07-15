package service

import (
	entModel "basic-crud-go/internal/app/admin/enterprise/model"
	"basic-crud-go/internal/app/admin/enterprise/service"
	"basic-crud-go/internal/app/admin/user/dto"
	"basic-crud-go/internal/app/admin/user/model"
	"basic-crud-go/internal/app/admin/user/repository"
	util "basic-crud-go/internal/app/util/password"
	"basic-crud-go/internal/configuration/logger"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"
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

func (s *userService) ReadAll(ctx context.Context, page, limit int) ([]model.UserExtend, error) {
	var users []model.UserExtend

	if page < 1 {
		logger.LogWithAutoFuncName(logger.Info, module, "page out of range. Defaulting to 1.")
		page = 1
	}

	if limit <= 0 || limit > 10 {
		logger.LogWithAutoFuncName(logger.Info, module, "limit out of range. Defaulting to 10.")
		limit = 10
	}

	users, err := s.repo.ReadAll(ctx, page, limit)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *userService) ReadByCnpj(ctx context.Context, cnpj string, page, limit int) ([]model.UserExtend, error) {
	var users []model.UserExtend

	if page < 1 {
		logger.LogWithAutoFuncName(logger.Info, module, "page out of range. Defaulting to 1.")
		page = 1
	}

	if limit <= 0 || limit > 10 {
		logger.LogWithAutoFuncName(logger.Info, module, "limit out of range. Defaulting to 10.")
		limit = 10
	}

	enterprise, err := s.enterpriseService.Read(ctx, cnpj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("enterprise not found")
		}
		return nil, err
	}

	users, err = s.repo.ReadUsersByEnterpriseID(ctx, enterprise.Id, page, limit)
	if err != nil {
		return nil, err
	}

	return users, nil
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

func (s *userService) Update(ctx context.Context, dto dto.UpdateUserDTO) (*model.User, *entModel.Enterprise, error) {
	const method = "Update"

	// Fetch user and enterprise via service
	existingUser, enterprise, err := s.Read(ctx, dto.Email)
	if err != nil {
		logger.Log(logger.Info, module, method, fmt.Errorf("user or enterprise not found: %w", err))
		return nil, nil, fmt.Errorf("user not found")
	}

	// Prepare updatedUser from existing
	updatedUser := model.User{
		Id:           existingUser.Id,
		EnterpriseId: existingUser.EnterpriseId,
		Number:       existingUser.Number,
		FirstName:    existingUser.FirstName,
		LastName:     existingUser.LastName,
		Email:        existingUser.Email,
		Password:     existingUser.Password,
		CreatedAt:    existingUser.CreatedAt,
		UpdatedAt:    time.Now(),
	}

	// FirstName
	if strings.TrimSpace(dto.FirstName) != "" && dto.FirstName != existingUser.FirstName {
		logger.Log(logger.Info, module, method, fmt.Errorf("updating first name to: %s", dto.FirstName))
		updatedUser.FirstName = dto.FirstName
	}

	// LastName
	if strings.TrimSpace(dto.LastName) != "" && dto.LastName != existingUser.LastName {
		logger.Log(logger.Info, module, method, fmt.Errorf("updating last name to: %s", dto.LastName))
		updatedUser.LastName = dto.LastName
	}

	// EmailUpdated
	if strings.TrimSpace(dto.EmailUpdated) != "" && dto.EmailUpdated != existingUser.Email {
		logger.Log(logger.Info, module, method, fmt.Errorf("updating email to: %s", dto.EmailUpdated))
		updatedUser.Email = dto.EmailUpdated
	}

	// Number
	if strings.TrimSpace(dto.Number) != "" && dto.Number != existingUser.Number {
		logger.Log(logger.Info, module, method, fmt.Errorf("updating number to: %s", dto.Number))
		updatedUser.Number = dto.Number
	}

	// Password
	if strings.TrimSpace(dto.Password) != "" {
		logger.Log(logger.Info, module, method, fmt.Errorf("received password input"))

		if err := util.Compare(existingUser.Password, dto.Password); err != nil {
			logger.Log(logger.Info, module, method, fmt.Errorf("password is different, hashing new one"))

			passwordHash, err := util.Hash(dto.Password)
			if err != nil {
				logger.Log(logger.Error, module, method, fmt.Errorf("failed to hash password: %w", err))
				return nil, nil, fmt.Errorf("failed to hash password: %w", err)
			}
			updatedUser.Password = passwordHash
		} else {
			logger.Log(logger.Info, module, method, fmt.Errorf("password is the same, keeping hash"))
		}
	} else {
		logger.Log(logger.Info, module, method, fmt.Errorf("password is empty, keeping current"))
		updatedUser.Password = existingUser.Password
	}

	// Save updated user
	user, err := s.repo.Update(ctx, updatedUser)
	if err != nil {
		logger.Log(logger.Error, module, method, fmt.Errorf("failed to update user: %w", err))
		return nil, nil, err
	}

	return user, enterprise, nil
}

func (s *userService) Delete(ctx context.Context, email string) (bool, error) {
	user, err := s.repo.Read(ctx, email)

	if err != nil {
		return false, err
	}

	_, err = s.repo.Delete(ctx, user.Id)
	if err != nil {
		return false, fmt.Errorf("failed to delete user: %w", err)
	}

	return true, nil
}
