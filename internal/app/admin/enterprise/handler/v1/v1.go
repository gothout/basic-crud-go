package v1

import (
	enterprise "basic-crud-go/internal/app/admin/enterprise/handler/v1/enterprise"
	mw "basic-crud-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterV1Routes(router *gin.RouterGroup, authMW *mw.Auth) {
	v1Group := router.Group("/v1")
	enterprise.RegisterEnterpriseRoutes(v1Group, authMW)
}
