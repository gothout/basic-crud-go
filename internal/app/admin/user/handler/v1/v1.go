package v1

import (
	user "basic-crud-go/internal/app/admin/user/handler/v1/user"
	"basic-crud-go/internal/app/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterV1Routes(router *gin.RouterGroup, mw *middleware.AuthMiddleware) {
	v1Group := router.Group("/v1")
	user.RegisterUserRoutes(v1Group, mw)
}
