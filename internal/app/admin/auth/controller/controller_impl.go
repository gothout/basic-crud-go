package controller

import (
	"basic-crud-go/internal/app/admin/auth/binding"
	"basic-crud-go/internal/app/admin/auth/dto"
	"basic-crud-go/internal/app/admin/auth/service"
	entDTO "basic-crud-go/internal/app/admin/enterprise/dto"
	permDto "basic-crud-go/internal/app/admin/permission/dto"
	userDTO "basic-crud-go/internal/app/admin/user/dto"
	"basic-crud-go/internal/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"net/http"
)

type authController struct {
	service service.AuthService
}

func NewAuthController(s service.AuthService) AuthController {
	return &authController{
		service: s,
	}
}

// AuthLoginHandler godoc
// @Summary      Login
// @Description  Authenticate user and return user data, token and permissions
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request  body      dto.LoginUserDTO  true  "Login credentials"
// @Success      200      {object}  dto.LoginUserResponse
// @Failure      400      {object}  rest_err.RestErr
// @Failure      403      {object}  rest_err.RestErr
// @Router       /auth/login [post]
func (ac *authController) AuthLoginHandler(ctx *gin.Context) {
	req, err := binding.ValidateLoginDTO(ctx)
	if err != nil {
		restErr := rest_err.NewBadRequestValidationError("Invalid request body", []rest_err.Causes{
			rest_err.NewCause("validation", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}
	identity, err := ac.service.LoginUser(ctx, req.Email, req.Password)
	if err != nil {
		restErr := rest_err.NewForbiddenError("Unable to process your login")
		ctx.JSON(restErr.Code, restErr)
		return
	}
	var permissionDTOs []permDto.ReadPermissionResponse
	if identity.Permissions != nil {
		for _, p := range *identity.Permissions {
			permissionDTOs = append(permissionDTOs, permDto.ReadPermissionResponse{
				Code:        p.Code,
				Description: p.Description,
			})
		}
	}

	ctx.JSON(http.StatusOK, &dto.LoginUserResponse{
		User: userDTO.ReadUserResponse{
			Id:        identity.User.Id,
			FirstName: identity.User.FirstName,
			LastName:  identity.User.LastName,
			Email:     identity.User.Email,
			Number:    identity.User.Number,
			CreatedAt: identity.User.CreatedAt,
			UpdatedAt: identity.User.UpdatedAt,
			Enterprise: entDTO.ReadEnterpriseResponse{
				Name:      identity.Enterprise.Name,
				Cnpj:      identity.Enterprise.Cnpj,
				CreatedAt: identity.User.CreatedAt,
				UpdatedAt: identity.User.UpdatedAt,
			},
		},
		Token:       identity.TokenUser.Token,
		Permissions: permDto.ReadPermissionsResponse{Permissions: permissionDTOs},
	})
}

// AuthRefreshHandler godoc
// @Summary      Refresh token
// @Description  Reset token expiration if it's still valid
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request  body      dto.RefreshTokenUserDTO  true  "Refresh credentials"
// @Success      200      {object}  dto.RefreshTokenUserResponse
// @Failure      400      {object}  rest_err.RestErr
// @Failure      403      {object}  rest_err.RestErr
// @Router       /auth/refresh [post]
func (ac *authController) AuthRefreshHandler(ctx *gin.Context) {
	req, err := binding.ValidateRefreshTokenUserDTO(ctx)
	if err != nil {
		restErr := rest_err.NewBadRequestValidationError("Invalid request body", []rest_err.Causes{
			rest_err.NewCause("validation", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}
	// extract token header
	authHeader := ctx.GetHeader("Authorization")
	if len(authHeader) < 8 || authHeader[:7] != "Bearer " {
		restErr := rest_err.NewBadRequestError("Missing or malformed Authorization header")
		ctx.JSON(restErr.Code, restErr)
		return
	}
	token := authHeader[7:]
	// Use service to refresh token (validates against DB and cache)
	success := ac.service.RefreshTokenUser(ctx, req.Email, token)
	if !success {
		restErr := rest_err.NewForbiddenError("Invalid credentials or token")
		ctx.JSON(restErr.Code, restErr)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Token refreshed successfully"})
}

// AuthLogoutHandler godoc
// @Summary      Logout user
// @Description  Logs user out of the system by invalidating their token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request  body      dto.LogoutUserDTO  true  "Logout request (only email)"
// @Success      204
// @Failure      400  {object}  rest_err.RestErr
// @Failure      403  {object}  rest_err.RestErr
// @Router       /auth/logout [post]
func (ac *authController) AuthLogoutHandler(ctx *gin.Context) {
	req, err := binding.ValidateLogoutUserDTO(ctx)
	if err != nil {
		restErr := rest_err.NewBadRequestValidationError("Invalid request body", []rest_err.Causes{
			rest_err.NewCause("validation", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}
	// extract token header
	authHeader := ctx.GetHeader("Authorization")
	if len(authHeader) < 8 || authHeader[:7] != "Bearer " {
		restErr := rest_err.NewBadRequestError("Missing or malformed Authorization header")
		ctx.JSON(restErr.Code, restErr)
		return
	}
	token := authHeader[7:]

	// execute logout
	if !ac.service.LogoutUser(ctx, req.Email, token) {
		restErr := rest_err.NewForbiddenError("Invalid logout")
		ctx.JSON(restErr.Code, restErr)
		return
	}
	ctx.Status(http.StatusNoContent)
}
