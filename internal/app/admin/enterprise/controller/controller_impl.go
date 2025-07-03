package controller

import (
	"basic-crud-go/internal/app/admin/enterprise/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type enterpriseController struct {
	service service.EnterpriseService
}

func NewEnterpriseController(s service.EnterpriseService) EnterpriseController {
	return &enterpriseController{
		service: s,
	}
}

// Ping godoc
// @Summary      Healthcheck do Enterprise
// @Description  Retorna um pong para verificar se o serviço enterprise está ativo
// @Tags         Enterprise
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /enterprise/v1/ping [get]
func (c *enterpriseController) Ping(ctx *gin.Context) {
	result, _ := c.service.Ping(ctx.Request.Context())

	ctx.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}
