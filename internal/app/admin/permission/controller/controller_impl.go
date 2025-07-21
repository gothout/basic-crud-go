package controller

import (
	"basic-crud-go/internal/app/admin/permission/binding"
	"basic-crud-go/internal/app/admin/permission/dto"
	"basic-crud-go/internal/app/admin/permission/service"
	"basic-crud-go/internal/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const module string = "Admin-Permission-Controller"

type permissionControllerImpl struct {
	service service.PermissionService
}

func NewPermissionController(s service.PermissionService) PermissionController {
	return &permissionControllerImpl{
		service: s,
	}
}

// ReadAll godoc
// @Summary      Read permissions
// @Description  Retrieve a paginated list of permissions names
// @Tags         Permission
// @Accept       json
// @Produce      json
// @Param        page   query     int     false   "Page number (min 1)"
// @Param        limit  query     int     false   "Items per page (default: 10)"
// @Success      200   {object}  dto.ReadPermissionsResponse
// @Failure      400   {object}  rest_err.RestErr
// @Failure      500   {object}  rest_err.RestErr
// @Router       /permission/v1/ [get]
func (c *permissionControllerImpl) ReadAll(ctx *gin.Context) {
	req := binding.ValidatePermissionsDTO(ctx)

	perms, err := c.service.ReadAllPermissions(ctx, req.Page, req.Limit)
	if err != nil {
		restErr := rest_err.NewInternalServerError("error fetching permissions", nil)
		ctx.JSON(restErr.Code, restErr)
		return
	}

	var responses []dto.ReadPermissionResponse
	for _, perm := range perms {
		responses = append(responses, dto.ReadPermissionResponse{
			Code:        perm.Code,
			Description: perm.Description,
		})
	}

	ctx.JSON(http.StatusOK, dto.ReadPermissionsResponse{
		Permissions: responses,
	})
}

// Read godoc
// @Summary      Read permission
// @Description  Read permissions by full code name
// @Tags         Permission
// @Accept       json
// @Produce      json
// @Param        code   query     string     true   "Read code (min 4 characters)"
// @Success      200   {object}  dto.ReadPermissionResponse
// @Failure      400   {object}  rest_err.RestErr
// @Failure      500   {object}  rest_err.RestErr
// @Router       /permission/v1/read [get]
func (c *permissionControllerImpl) Read(ctx *gin.Context) {
	req, err := binding.ValidateReadPermissioDTO(ctx)
	if err != nil {
		restErr := rest_err.NewBadRequestValidationError("invalid query", []rest_err.Causes{
			rest_err.NewCause("validation", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	perms, err := c.service.ReadByCode(ctx, req.Code)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			restErr := rest_err.NewNotFoundError("code not found")
			ctx.JSON(restErr.Code, restErr)
			return
		}
		restErr := rest_err.NewInternalServerError("error fetching permission", nil)
		ctx.JSON(restErr.Code, restErr)
		return
	}

	ctx.JSON(http.StatusOK, dto.ReadPermissionResponse{
		Code:        perms.Code,
		Description: perms.Description,
	})
}

// Search godoc
// @Summary      Search permissions
// @Description  Search permissions by partial or full code name
// @Tags         Permission
// @Accept       json
// @Produce      json
// @Param        query   query     string     true   "Search query (min 4 characters)"
// @Success      200   {object}  dto.ReadPermissionsResponse
// @Failure      400   {object}  rest_err.RestErr
// @Failure      500   {object}  rest_err.RestErr
// @Router       /permission/v1/search [get]
func (c *permissionControllerImpl) Search(ctx *gin.Context) {
	req, err := binding.ValidateSearchPermissionDTO(ctx)
	if err != nil {
		restErr := rest_err.NewBadRequestValidationError("invalid query", []rest_err.Causes{
			rest_err.NewCause("validation", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	perms, err := c.service.Search(ctx, req.Query)

	var responses []dto.ReadPermissionResponse
	for _, perm := range perms {
		responses = append(responses, dto.ReadPermissionResponse{
			Code:        perm.Code,
			Description: perm.Description,
		})
	}

	ctx.JSON(http.StatusOK, dto.ReadPermissionsResponse{
		Permissions: responses,
	})
}

// Apply godoc
// @Summary      Apply permissions to user
// @Description  Apply a batch of permissions to a user by email
// @Tags         Permission
// @Accept       json
// @Produce      json
// @Param        payload  body      dto.ApplyPermissionBatchDTO  true  "Email and permission codes"
// @Success      204
// @Failure      400   {object}  rest_err.RestErr
// @Failure      500   {object}  rest_err.RestErr
// @Router       /permission/v1/apply [post]
func (c *permissionControllerImpl) Apply(ctx *gin.Context) {
	req, err := binding.ValidateApplyPermissionBatchDTO(ctx)
	if err != nil {
		restErr := rest_err.NewBadRequestValidationError("invalid input", []rest_err.Causes{
			rest_err.NewCause("validation", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	err = c.service.ApplyPermissionUserBatch(ctx, req.Email, req.Codes)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			restErr := rest_err.NewNotFoundError(err.Error())
			ctx.JSON(restErr.Code, restErr)
			return
		}
		restErr := rest_err.NewInternalServerError("failed to apply permissions", nil)
		ctx.JSON(restErr.Code, restErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}
