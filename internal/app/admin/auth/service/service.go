package service

import (
	"basic-crud-go/internal/app/admin/auth/model"
	"basic-crud-go/internal/app/admin/auth/repository"
	tokencache "basic-crud-go/internal/app/admin/auth/util/token_cache"
	permService "basic-crud-go/internal/app/admin/permission/service"
	userService "basic-crud-go/internal/app/admin/user/service"
	security "basic-crud-go/internal/app/util/password"
	"context"
	"fmt"
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
	// check if token exists in cache
	if identity, found := tokencache.GetToken(email); found {
		// validate password user
		err := security.Compare(identity.User.Password, password)
		if err != nil {
			return nil, err
		}
		return identity, nil
	}

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

	// save token in cache
	tokencache.SaveToken(email, identity)

	return identity, nil
}

func (s *authServiceImpl) RefreshTokenUser(ctx context.Context, email, token string) bool {
	// Validate user credentials from the database
	//user, _, err := s.userService.Read(ctx, email)
	//if err != nil {
	//	return false
	//}

	//if err := security.Compare(user.Password, password); err != nil {
	//return false
	//}

	// Check if the cached token matches and calculate time remaining
	identity, found := tokencache.GetToken(email)
	if !found || identity.TokenUser.Token != token {
		return false
	}

	ttl, valid := tokencache.GetRemainingTTL(email)
	if !valid || ttl <= 10*time.Minute {
		return false // do not refresh if token is near expiration or invalid
	}

	// Refresh token expiration
	return tokencache.RefreshToken(email, token)
}

func (s *authServiceImpl) LogoutUser(ctx context.Context, email, token string) bool {
	if !tokencache.Logout(email, token) {
		return false
	}
	return true
}

func (s *authServiceImpl) GenerateTokenAPI(ctx context.Context, email, token string, expiresAt time.Time) (string, error) {
	// Validate user credentials from the database
	user, _, err := s.userService.Read(ctx, email)
	if err != nil {
		return "", err
	}
	// Check if the cached token matches
	identity, found := tokencache.GetToken(email)
	if !found || identity.TokenUser.Token != token {
		return "", fmt.Errorf("not authorize")
	}

	dateNow := time.Now()

	// generate token
	tokenApi, err := s.repo.GenerateTokenAPI(ctx, user.Id, dateNow, expiresAt)
	if err != nil {
		return "", err
	}

	return tokenApi, nil
}

func (s *authServiceImpl) GetUserIdByAPIKey(ctx context.Context, apiKey string) (string, error) {

	token, err := s.repo.GetUserIdByAPIKey(ctx, apiKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *authServiceImpl) GetUserIdByUserKey(ctx context.Context, apiKey string) (string, error) {
	token, err := s.repo.GetValidUserIdByToken(ctx, apiKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *authServiceImpl) GetUserIdentiyByCache(ctx context.Context, userKey string) (*model.UserIdentity, error) {
	// check if token exists in cache
	if identity, found := tokencache.GetByUserToken(userKey); found {
		return identity, nil
	}
	return nil, nil
}
