package controller

import (
	entepriseDto "basic-crud-go/internal/app/admin/enterprise/dto"
	mwUtil "basic-crud-go/internal/app/admin/middleware/util"
	"basic-crud-go/internal/app/admin/user/binding"
	"basic-crud-go/internal/app/admin/user/dto"
	"basic-crud-go/internal/app/admin/user/model"
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
// @Security     BearerAuth
// @Param        request  body      dto.CreateUserDTO       true "User Data"
// @Success      201      {object}  dto.CreateUserResponse
// @Failure      400      {object}  rest_err.RestErr
// @Failure      404      {object}  rest_err.RestErr
// @Failure      500      {object}  rest_err.RestErr
// @Router       /user/v1/ [post]
func (c *userController) CreateUserHandler(ctx *gin.Context) {
	var req dto.CreateUserDTO

	// 1) Parse request body
	if err := ctx.ShouldBindJSON(&req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("Missing or invalid required fields",
			[]rest_err.Causes{rest_err.NewCause("body", err.Error())})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// 2) Validate DTO
	if err := binding.ValidateCreateUserDTO(req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("Invalid request body",
			[]rest_err.Causes{rest_err.NewCause("validation", err.Error())})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// 3) Get identity from context (set by AuthMiddleware)
	identity, rerr := mwUtil.GetIdentity(ctx)
	if rerr != nil {
		restErr := rest_err.NewBadRequestValidationError("Invalid request body",
			[]rest_err.Causes{rest_err.NewCause("context", rerr.Error())})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// 4) Permission checks
	hasSystem := mwUtil.HasPermission(identity.Permissions, "system")
	hasCreateOwn := mwUtil.HasPermission(identity.Permissions, "create-user-enterprise")
	hasCreateAdmin := mwUtil.HasPermission(identity.Permissions, "create-user-admin")

	// User has no permission to create any user
	if !(hasSystem || hasCreateOwn || hasCreateAdmin) {
		restErr := rest_err.NewForbiddenError("you do not have permission to create users")
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// User can only create users in their own enterprise
	if hasCreateOwn && !hasSystem && !hasCreateAdmin {
		// if request CNPJ is different, deny
		if req.Cnpj != "" && util.RemoveNonDigits(req.Cnpj) != util.RemoveNonDigits(identity.Enterprise.Cnpj) {
			restErr := rest_err.NewForbiddenError("you can only create users for your own enterprise")
			ctx.JSON(restErr.Code, restErr)
			return
		}
		// force enterprise CNPJ from identity
		req.Cnpj = identity.Enterprise.Cnpj
	}

	// 5) Normalize inputs
	cnpj := util.RemoveNonDigits(req.Cnpj)
	number := util.RemoveNonDigits(req.Number)

	// 6) Call service
	created, err := c.service.Create(ctx,
		cnpj,
		number,
		req.FirstName,
		req.LastName,
		req.Email,
		req.Password,
	)
	if err != nil && strings.Contains(err.Error(), "nao encontrada") {
		restErr := rest_err.NewNotFoundError(fmt.Sprintf("cnpj %s not found", cnpj))
		ctx.JSON(restErr.Code, restErr)
		return
	}
	if err != nil {
		restErr := rest_err.NewInternalServerError("Error creating user", nil)
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// 7) Return 201 Created
	ctx.JSON(http.StatusCreated, &dto.CreateUserResponse{
		Cnpj:      cnpj,
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
// @Security     BearerAuth
// @Param        page   query     int     false   "Page number (min 1)"
// @Param        limit  query     int     false   "Items per page (default: 10)"
// @Success      200    {object}  dto.ReadUsersResponse
// @Failure      400    {object}  rest_err.RestErr
// @Failure      404    {object}  rest_err.RestErr
// @Failure      500    {object}  rest_err.RestErr
// @Router       /user/v1/read [get]
func (c *userController) ReadUsersHandler(ctx *gin.Context) {
	// 1) Parse query parameters
	req := binding.ValidateReadUsersDTO(ctx)

	// 2) Get identity from context (set by AuthMiddleware)
	identity, rerr := mwUtil.GetIdentity(ctx)
	if rerr != nil {
		restErr := rest_err.NewBadRequestValidationError("invalid request body",
			[]rest_err.Causes{rest_err.NewCause("context", rerr.Error())})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// 3) Permission checks
	hasSystem := mwUtil.HasPermission(identity.Permissions, "system")
	hasReadAny := mwUtil.HasPermission(identity.Permissions, "read-user")
	hasReadOwn := mwUtil.HasPermission(identity.Permissions, "read-user-enterprise")

	if !(hasSystem || hasReadAny || hasReadOwn) {
		restErr := rest_err.NewForbiddenError("you do not have permission to read users")
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// 4) Fetch users according to permission
	var (
		users []model.UserExtend
		err   error
	)
	if hasReadOwn && !hasSystem && !hasReadAny {
		cnpj := util.RemoveNonDigits(identity.Enterprise.Cnpj)
		users, err = c.service.ReadByCnpj(ctx, cnpj, req.Page, req.Limit)
	} else {
		users, err = c.service.ReadAll(ctx, req.Page, req.Limit)
	}
	if err != nil {
		restErr := rest_err.NewInternalServerError("failed to fetch users", []rest_err.Causes{
			rest_err.NewCause("read users", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// 5) Build response
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

	// 6) Return 200 OK
	ctx.JSON(http.StatusOK, dto.ReadUsersResponse{
		Users: result,
	})
}

// ReadUsers godoc
// @Summary      List users by CNPJ
// @Description  Retrieve a paginated list of users
// @Tags         User
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        cnpj   query     string  true    "CNPJ enterprise"
// @Param        page   query     int     false   "Page number (min 1)"
// @Param        limit  query     int     false   "Items per page (default: 10)"
// @Success      200    {object}  dto.ReadUsersResponse
// @Failure      400    {object}  rest_err.RestErr
// @Failure      404    {object}  rest_err.RestErr
// @Failure      500    {object}  rest_err.RestErr
// @Router       /user/v1/read/enterprise [get]
func (c *userController) ReadUsersByCnpjHandler(ctx *gin.Context) {
	// 1) Validate request parameters
	req, err := binding.ValidateReadUsersByCnpjDTO(ctx)
	if err != nil {
		restErr := rest_err.NewBadRequestValidationError("Invalid request body", []rest_err.Causes{
			rest_err.NewCause("validation", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// 2) Get identity from context (set by AuthMiddleware)
	identity, rerr := mwUtil.GetIdentity(ctx)
	if rerr != nil {
		restErr := rest_err.NewBadRequestValidationError("Invalid request body", []rest_err.Causes{
			rest_err.NewCause("context", rerr.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// 3) Permission checks
	hasSystem := mwUtil.HasPermission(identity.Permissions, "system")
	hasReadAny := mwUtil.HasPermission(identity.Permissions, "read-user")
	hasReadOwn := mwUtil.HasPermission(identity.Permissions, "read-user-enterprise")

	if !(hasSystem || hasReadAny || hasReadOwn) {
		restErr := rest_err.NewForbiddenError("you do not have permission to read users")
		ctx.JSON(restErr.Code, restErr)
		return
	}

	if hasReadOwn && !hasSystem && !hasReadAny {
		identityCnpj := util.RemoveNonDigits(identity.Enterprise.Cnpj)
		if req.Cnpj != identityCnpj {
			restErr := rest_err.NewForbiddenError("you can only read users from your own enterprise")
			ctx.JSON(restErr.Code, restErr)
			return
		}
	}

	// 4) Fetch users by CNPJ
	users, err := c.service.ReadByCnpj(ctx, req.Cnpj, req.Page, req.Limit)
	if err != nil {
		if strings.Contains(err.Error(), "nao encontrada") {
			restErr := rest_err.NewNotFoundError(fmt.Sprintf("enterprise %s not found", req.Cnpj))
			ctx.JSON(restErr.Code, restErr)
			return
		}
		restErr := rest_err.NewInternalServerError("failed to fetch users", []rest_err.Causes{
			rest_err.NewCause("read users", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// 5) Build response
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

	// 6) Return 200 OK
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
// @Security     BearerAuth
// @Param        email 	   path     string  true "User email"
// @Success      201      {object}  dto.ReadUserResponse
// @Failure      400      {object}  rest_err.RestErr
// @Failure		 404	  {object}	rest_err.RestErr
// @Failure      500      {object}  rest_err.RestErr
// @Router       /user/v1/{email} [get]
func (c *userController) ReadUserHandler(ctx *gin.Context) {
	// 1) Validate path parameter
	req, err := binding.ValidateReadUserDTO(ctx)
	if err != nil {
		restErr := rest_err.NewBadRequestValidationError("invalid request body", []rest_err.Causes{
			rest_err.NewCause("validation", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// 2) Get identity from context (set by AuthMiddleware)
	identity, rerr := mwUtil.GetIdentity(ctx)
	if rerr != nil {
		restErr := rest_err.NewBadRequestValidationError("invalid request body", []rest_err.Causes{
			rest_err.NewCause("context", rerr.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// 3) Read user
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

	// 4) Permission checks
	hasSystem := mwUtil.HasPermission(identity.Permissions, "system")
	hasReadAny := mwUtil.HasPermission(identity.Permissions, "read-user")
	hasReadOwn := mwUtil.HasPermission(identity.Permissions, "read-user-enterprise")

	if !(hasSystem || hasReadAny || hasReadOwn) {
		restErr := rest_err.NewForbiddenError("you do not have permission to read users")
		ctx.JSON(restErr.Code, restErr)
		return
	}

	if hasReadOwn && !hasSystem && !hasReadAny {
		identityCnpj := util.RemoveNonDigits(identity.Enterprise.Cnpj)
		userCnpj := util.RemoveNonDigits(enterprise.Cnpj)
		if identityCnpj != userCnpj {
			restErr := rest_err.NewForbiddenError("you can only read users from your own enterprise")
			ctx.JSON(restErr.Code, restErr)
			return
		}
	}

	// 5) Return 200 OK
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
// @Security     BearerAuth
// @Param        email    path      string              true  "User email"
// @Param        request  body      dto.UpdateUserDTO   true  "User update data"
// @Success      200      {object}  dto.UpdateUserResponse
// @Failure      400      {object}  rest_err.RestErr
// @Failure      404      {object}  rest_err.RestErr
// @Failure      500      {object}  rest_err.RestErr
// @Router       /user/v1/{email} [put]
func (c *userController) UpdateUserHandler(ctx *gin.Context) {
	// 1) Validate request body and path
	req, err := binding.ValidateUpdateUserDTO(ctx)
	if err != nil {
		restErr := rest_err.NewBadRequestValidationError("invalid request", []rest_err.Causes{
			rest_err.NewCause("validation", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// 2) Get identity from context (set by AuthMiddleware)
	identity, rerr := mwUtil.GetIdentity(ctx)
	if rerr != nil {
		restErr := rest_err.NewBadRequestValidationError("invalid request", []rest_err.Causes{
			rest_err.NewCause("context", rerr.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// 3) Permission checks
	hasSystem := mwUtil.HasPermission(identity.Permissions, "system")
	hasUpdateOwn := mwUtil.HasPermission(identity.Permissions, "update-enterprise-user")

	if !(hasSystem || hasUpdateOwn) {
		restErr := rest_err.NewForbiddenError("you do not have permission to update users")
		ctx.JSON(restErr.Code, restErr)
		return
	}

	if hasUpdateOwn && !hasSystem {
		_, enterprise, err := c.service.Read(ctx, req.Email)
		if err != nil && strings.Contains(err.Error(), "nao encontrada") {
			restErr := rest_err.NewNotFoundError(fmt.Sprintf("user %s not found", req.Email))
			ctx.JSON(restErr.Code, restErr)
			return
		}
		if err != nil {
			restErr := rest_err.NewInternalServerError("failed to fetch user", nil)
			ctx.JSON(restErr.Code, restErr)
			return
		}
		identityCnpj := util.RemoveNonDigits(identity.Enterprise.Cnpj)
		userCnpj := util.RemoveNonDigits(enterprise.Cnpj)
		if identityCnpj != userCnpj {
			restErr := rest_err.NewForbiddenError("you can only update users from your own enterprise")
			ctx.JSON(restErr.Code, restErr)
			return
		}
	}

	// 4) Update user
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

	// 5) Return 200 OK
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
// @Security     BearerAuth
// @Param        email 	   path     string  true "User email"
// @Success      204
// @Failure      400      {object}  rest_err.RestErr
// @Failure		 404	  {object}	rest_err.RestErr
// @Failure      500      {object}  rest_err.RestErr
// @Router       /user/v1/{email} [delete]
func (c *userController) DeleteUserHandler(ctx *gin.Context) {
	// 1) Validate path parameter
	req, err := binding.ValidateDeleteUserDTO(ctx)
	if err != nil {
		restErr := rest_err.NewBadRequestValidationError("invalid request body", []rest_err.Causes{
			rest_err.NewCause("validation", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// 2) Get identity from context (set by AuthMiddleware)
	identity, rerr := mwUtil.GetIdentity(ctx)
	if rerr != nil {
		restErr := rest_err.NewBadRequestValidationError("invalid request body", []rest_err.Causes{
			rest_err.NewCause("context", rerr.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// 3) Permission checks
	hasSystem := mwUtil.HasPermission(identity.Permissions, "system")
	hasDeleteOwn := mwUtil.HasPermission(identity.Permissions, "delete-enterprise-user")

	if !(hasSystem || hasDeleteOwn) {
		restErr := rest_err.NewForbiddenError("you do not have permission to delete users")
		ctx.JSON(restErr.Code, restErr)
		return
	}

	if hasDeleteOwn && !hasSystem {
		_, enterprise, err := c.service.Read(ctx, req.Email)
		if err != nil && strings.Contains(err.Error(), "nao encontrada") {
			restErr := rest_err.NewNotFoundError(fmt.Sprintf("user %s not found", req.Email))
			ctx.JSON(restErr.Code, restErr)
			return
		}
		if err != nil {
			restErr := rest_err.NewInternalServerError("failed to fetch user", nil)
			ctx.JSON(restErr.Code, restErr)
			return
		}
		identityCnpj := util.RemoveNonDigits(identity.Enterprise.Cnpj)
		userCnpj := util.RemoveNonDigits(enterprise.Cnpj)
		if identityCnpj != userCnpj {
			restErr := rest_err.NewForbiddenError("you can only delete users from your own enterprise")
			ctx.JSON(restErr.Code, restErr)
			return
		}
	}

	// 4) Delete user
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

	// 5) Return 204 No Content
	ctx.JSON(http.StatusNoContent, nil)
}
