package handler

import (
	_ "basic-crud-go/docs"
	v1 "basic-crud-go/internal/app/admin/auth/handler/v1"

	"github.com/gin-gonic/gin"
)

// RegisterAuthRoutes register all routes domain auth
func RegisterAuthRoutes(router *gin.Engine) {
	userGroup := router.Group("/auth")
	v1.RegisterV1Routes(userGroup)
}
