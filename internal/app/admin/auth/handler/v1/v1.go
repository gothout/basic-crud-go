package v1

import (
	auth "basic-crud-go/internal/app/admin/auth/handler/v1/auth"

	"github.com/gin-gonic/gin"
)

func RegisterV1Routes(router *gin.RouterGroup) {
	v1Group := router.Group("/")
	auth.RegisterAuthRoutes(v1Group)
}
