package controller

import (
	"basic-crud-go/internal/app/admin/permission/binding"
	"basic-crud-go/internal/app/admin/permission/dto"
	"basic-crud-go/internal/app/admin/permission/service"
	"basic-crud-go/internal/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"net/http"
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
// @Description  Retrieve a paginated list of permissionss names
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
