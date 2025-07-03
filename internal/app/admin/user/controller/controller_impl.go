package controller

import (
	"basic-crud-go/internal/app/admin/user/service"
	"net/http"

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
// @Summary      Healthcheck do Admin
// @Description  Retorna um pong para verificar se o serviço admin está ativo
// @Tags         Admin
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /user/v1/ping [get]
func (c *userController) Ping(ctx *gin.Context) {
	result, _ := c.service.Ping(ctx.Request.Context())

	ctx.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}
