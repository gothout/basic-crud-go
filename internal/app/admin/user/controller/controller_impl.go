package controller

import (
	entepriseDto "basic-crud-go/internal/app/admin/enterprise/dto"
	"basic-crud-go/internal/app/admin/user/binding"
	"basic-crud-go/internal/app/admin/user/dto"
	"basic-crud-go/internal/app/admin/user/service"
	"basic-crud-go/internal/app/admin/user/util"
	"basic-crud-go/internal/configuration/rest_err"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type userController struct {
	service service.UserService
}

func NewUserController(s service.UserService) UserController {
	return &userController{
		service: s,
	}
}

// Ping godoc
// @Summary      Healthcheck do User
// @Description  Retorna um pong para verificar se o serviço User está ativo
// @Tags         User
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /user/v1/ping [get]
func (c *userController) Ping(ctx *gin.Context) {
	result, _ := c.service.Ping(ctx.Request.Context())

	ctx.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}

// CreateUserHandler godoc
// @Summary      Create user
// @Description  Create user by CNPJ, name and email
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        request  body      dto.CreateUserDTO       true "User Data"
// @Success      201      {object}  dto.CreateUserResponse
// @Failure      400      {object}  rest_err.RestErr
// @Failure		 404	  {object}	rest_err.RestErr
// @Failure      500      {object}  rest_err.RestErr
// @Router       /user/v1/ [post]
func (c *userController) CreateUserHandler(ctx *gin.Context) {
	var req dto.CreateUserDTO

	// Parse request body
	if err := ctx.ShouldBindJSON(&req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("Missing or invalid required fields", []rest_err.Causes{
			rest_err.NewCause("body", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// Validate DTO
	if err := binding.ValidateCreateUserDTO(req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("Invalid request body", []rest_err.Causes{
			rest_err.NewCause("validation", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// Create user
	created, err := c.service.Create(ctx,
		util.RemoveNonDigits(req.Cnpj),
		util.RemoveNonDigits(req.Number),
		req.FirstName,
		req.LastName,
		req.Email,
		req.Password,
	)
	if err != nil && strings.Contains(err.Error(), "nao encontrada") {
		restErr := rest_err.NewNotFoundError(fmt.Sprintf("cnpj %s not found", util.RemoveNonDigits(req.Cnpj)))
		ctx.JSON(restErr.Code, restErr)
		return
	}
	if err != nil {
		restErr := rest_err.NewInternalServerError("Error creating user", nil)
		ctx.JSON(restErr.Code, restErr)
		return
	}

	ctx.JSON(http.StatusOK, &dto.CreateUserResponse{
		Cnpj:      req.Cnpj,
		FirstName: created.FirstName,
		LastName:  created.LastName,
		Email:     created.Email,
	})
}

// ReadUser godoc
// @Summary      Read user
// @Description  Read user by email
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        email 	   path     string  true "User email"
// @Success      201      {object}  dto.ReadUserResponse
// @Failure      400      {object}  rest_err.RestErr
// @Failure		 404	  {object}	rest_err.RestErr
// @Failure      500      {object}  rest_err.RestErr
// @Router       /user/v1/{email} [get]
func (c *userController) ReadUserHandler(ctx *gin.Context) {
	req, err := binding.ValidateReadUserDTO(ctx)
	if err != nil {
		restErr := rest_err.NewBadRequestValidationError("invalid request body", []rest_err.Causes{
			rest_err.NewCause("validation", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// read user
	user, enterprise, err := c.service.Read(ctx, req.Email)
	if err != nil && strings.Contains(err.Error(), "nao encontrada") {
		restErr := rest_err.NewNotFoundError(fmt.Sprintf("user %s not found", req.Email))
		ctx.JSON(restErr.Code, restErr)
		return
	}
	if err != nil {
		restErr := rest_err.NewInternalServerError("error read user", nil)
		ctx.JSON(restErr.Code, restErr)
		return
	}
	ctx.JSON(http.StatusOK, &dto.ReadUserResponse{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Number:    user.Number,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Enterprise: entepriseDto.ReadEnterpriseResponse{
			Name:      enterprise.Name,
			Cnpj:      enterprise.Cnpj,
			CreatedAt: enterprise.CreateAt,
			UpdatedAt: enterprise.UpdateAt,
		},
	})

}
