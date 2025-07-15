package controller

import (
	entepriseDto "basic-crud-go/internal/app/admin/enterprise/dto"
	"basic-crud-go/internal/app/admin/user/binding"
	"basic-crud-go/internal/app/admin/user/dto"
	"basic-crud-go/internal/app/admin/user/service"
	"basic-crud-go/internal/app/admin/user/util"
	"basic-crud-go/internal/configuration/rest_err"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type userController struct {
	service service.UserService
}

func NewUserController(s service.UserService) UserController {
	return &userController{
		service: s,
	}
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
	var Cnpj string = util.RemoveNonDigits(req.Cnpj)
	var Number string = util.RemoveNonDigits(req.Number)

	// Create user
	created, err := c.service.Create(ctx,
		Cnpj,
		Number,
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
		Cnpj:      Cnpj,
		FirstName: created.FirstName,
		LastName:  created.LastName,
		Email:     created.Email,
		Number:    created.Number,
	})
}

// ReadUsers godoc
// @Summary      List users
// @Description  Retrieve a paginated list of users
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        page   query     int     false   "Page number (min 1)"
// @Param        limit  query     int     false   "Items per page (default: 10)"
// @Success      200    {object}  dto.ReadUsersResponse
// @Failure      400    {object}  rest_err.RestErr
// @Failure      404    {object}  rest_err.RestErr
// @Failure      500    {object}  rest_err.RestErr
// @Router       /user/v1/read [get]
func (c *userController) ReadUsersHandler(ctx *gin.Context) {
	req := binding.ValidateReadUsersDTO(ctx)

	users, err := c.service.ReadAll(ctx, req.Page, req.Limit)
	if err != nil {
		restErr := rest_err.NewInternalServerError("failed to fetch users", []rest_err.Causes{
			rest_err.NewCause("read users", "error"),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	var result []dto.ReadUserResponse
	for _, ent := range users {
		result = append(result, dto.ReadUserResponse{
			Id:        ent.Id,
			FirstName: ent.FirstName,
			LastName:  ent.LastName,
			Email:     ent.Email,
			Number:    ent.Number,
			CreatedAt: ent.CreatedAt,
			UpdatedAt: ent.UpdatedAt,
			Enterprise: entepriseDto.ReadEnterpriseResponse{
				Name:      ent.Enterprise.Name,
				Cnpj:      ent.Enterprise.Cnpj,
				CreatedAt: ent.Enterprise.CreateAt,
				UpdatedAt: ent.Enterprise.UpdateAt,
			},
		})
	}

	ctx.JSON(http.StatusOK, dto.ReadUsersResponse{
		Users: result,
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

// UpdateUser godoc
// @Summary      Update user
// @Description  Update user by email (URI param) and update fields from body
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        email    path      string              true  "User email"
// @Param        request  body      dto.UpdateUserDTO   true  "User update data"
// @Success      200      {object}  dto.UpdateUserResponse
// @Failure      400      {object}  rest_err.RestErr
// @Failure      404      {object}  rest_err.RestErr
// @Failure      500      {object}  rest_err.RestErr
// @Router       /user/v1/{email} [put]
func (c *userController) UpdateUserHandler(ctx *gin.Context) {
	req, err := binding.ValidateUpdateUserDTO(ctx)
	if err != nil {
		restErr := rest_err.NewBadRequestValidationError("invalid request", []rest_err.Causes{
			rest_err.NewCause("validation", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	updatedUser, enterprise, err := c.service.Update(ctx, *req)
	if err != nil && strings.Contains(err.Error(), "not found") {
		restErr := rest_err.NewNotFoundError(fmt.Sprintf("user %s not found", req.Email))
		ctx.JSON(restErr.Code, restErr)
		return
	}

	if err != nil && strings.Contains(err.Error(), "user_email_key") {
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("it is not possible to update to this new email %s", req.Email))
		ctx.JSON(restErr.Code, restErr)
		return
	}
	if err != nil {
		restErr := rest_err.NewInternalServerError("failed to update user", nil)
		ctx.JSON(restErr.Code, restErr)
		return
	}

	ctx.JSON(http.StatusOK, &dto.UpdateUserResponse{
		UpdatedUser: dto.ReadUserResponse{
			Id:        updatedUser.Id,
			FirstName: updatedUser.FirstName,
			LastName:  updatedUser.LastName,
			Email:     updatedUser.Email,
			Number:    updatedUser.Number,
			CreatedAt: updatedUser.CreatedAt,
			UpdatedAt: updatedUser.UpdatedAt,
			Enterprise: entepriseDto.ReadEnterpriseResponse{
				Name:      enterprise.Name,
				Cnpj:      enterprise.Cnpj,
				CreatedAt: enterprise.CreateAt,
				UpdatedAt: enterprise.UpdateAt,
			},
		},
	})
}

// DeleteUser godoc
// @Summary      Delete user
// @Description  Delete user by email
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        email 	   path     string  true "User email"
// @Success      204
// @Failure      400      {object}  rest_err.RestErr
// @Failure		 404	  {object}	rest_err.RestErr
// @Failure      500      {object}  rest_err.RestErr
// @Router       /user/v1/{email} [delete]
func (c *userController) DeleteUserHandler(ctx *gin.Context) {
	req, err := binding.ValidateDeleteUserDTO(ctx)
	if err != nil {
		restErr := rest_err.NewBadRequestValidationError("invalid request body", []rest_err.Causes{
			rest_err.NewCause("validation", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// read user
	_, err = c.service.Delete(ctx, req.Email)
	if err != nil && strings.Contains(err.Error(), "user not found") {
		restErr := rest_err.NewNotFoundError(fmt.Sprintf("user %s not found", req.Email))
		ctx.JSON(restErr.Code, restErr)
		return
	}
	if err != nil {
		restErr := rest_err.NewInternalServerError("error delete user", nil)
		ctx.JSON(restErr.Code, restErr)
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
