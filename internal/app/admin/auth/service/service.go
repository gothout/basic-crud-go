package service

import (
	"basic-crud-go/internal/app/admin/auth/model"
	"basic-crud-go/internal/app/admin/auth/repository"
	permService "basic-crud-go/internal/app/admin/permission/service"
	userService "basic-crud-go/internal/app/admin/user/service"
	security "basic-crud-go/internal/app/util/password"
	"context"
	"time"
)

const module string = "Admin-Auth-Service"

type authServiceImpl struct {
	repo        repository.AuthRepository
	userService userService.UserService
	permService permService.PermissionService
}

func NewAuthService(repo repository.AuthRepository, userService userService.UserService, permService permService.PermissionService) AuthService {
	return &authServiceImpl{
		repo:        repo,
		userService: userService,
		permService: permService,
	}
}

func (s *authServiceImpl) LoginUser(ctx context.Context, email, password string) (*model.UserIdentity, error) {
	// validate email user
	user, enterprise, err := s.userService.Read(ctx, email)
	if err != nil {
		return nil, err
	}

	// validate password user
	err = security.Compare(user.Password, password)
	if err != nil {
		return nil, err
	}

	// get permissions user
	perms, err := s.permService.ReadPermissionsUser(ctx, email)
	if err != nil {
		return nil, err
	}

	// define created time
	createdAt := time.Now()

	// create token
	userToken, err := s.repo.GenerateTokenUser(ctx, user.Id, createdAt)
	if err != nil {
		return nil, err
	}

	identity := &model.UserIdentity{
		User:        user,
		Enterprise:  enterprise,
		Permissions: &perms,
		TokenUser:   userToken,
		TokenApi:    nil,
	}

	return identity, nil
}
