package server

import (
	userHandler "basic-crud-go/internal/app/admin/user/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	userHandler.RegisterUserRoutes(router)
}
