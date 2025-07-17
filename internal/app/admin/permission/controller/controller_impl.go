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

// Read godoc
// @Summary      Read permission
// @Description  Read permission by name
// @Tags         Permission
// @Accept       json
// @Produce      json
// @Param        name  path      string  true  "name of permission"
// @Success      200   {object}  dto.ReadPermissionResponse
// @Failure      400   {object}  rest_err.RestErr
// @Failure      500   {object}  rest_err.RestErr
// @Router       /permission/v1/read/{name} [get]
func (c *permissionControllerImpl) Read(ctx *gin.Context) {
	req, err := binding.ValidateReadPermissioDTO(ctx)
	if err != nil {
		restErr := rest_err.NewBadRequestValidationError("invalid request body", []rest_err.Causes{
			rest_err.NewCause("validation", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// Read
	mod, err := c.service.Read(ctx, req.Name)
	if err != nil {
		restErr := rest_err.NewNotFoundError("module not found")
		ctx.JSON(restErr.Code, restErr)
		return
	}

	ctx.JSON(http.StatusOK, &dto.ReadPermissionResponse{
		Name:    mod.Name,
		Actions: mod.Actions,
	})
}
