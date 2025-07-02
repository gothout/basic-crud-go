package controller

import (
	"basic-crud-go/internal/app/admin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type adminController struct {
	service service.AdminService
}

func NewAdminController(s service.AdminService) AdminController {
	return &adminController{
		service: s,
	}
}

// Ping godoc
// @Summary      Healthcheck do Admin
// @Description  Retorna um pong para verificar se o serviço admin está ativo
// @Tags         Admin
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /admin/v1/ping [get]
func (c *adminController) Ping(ctx *gin.Context) {
	result, _ := c.service.Ping(ctx.Request.Context())

	ctx.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}
