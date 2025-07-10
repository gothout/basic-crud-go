package handler

import (
	_ "basic-crud-go/docs"
	v1 "basic-crud-go/internal/app/admin/user/handler/v1"

	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes register all routes domain user
func RegisterUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/user")
	v1.RegisterV1Routes(userGroup)
}
