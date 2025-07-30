package controller

import (
	"basic-crud-go/internal/app/admin/auth/binding"
	"basic-crud-go/internal/app/admin/auth/dto"
	"basic-crud-go/internal/app/admin/auth/service"
	entDTO "basic-crud-go/internal/app/admin/enterprise/dto"
	permDto "basic-crud-go/internal/app/admin/permission/dto"
	userDTO "basic-crud-go/internal/app/admin/user/dto"
	"basic-crud-go/internal/configuration/rest_err"
	"fmt"
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
		fmt.Println(err)
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
