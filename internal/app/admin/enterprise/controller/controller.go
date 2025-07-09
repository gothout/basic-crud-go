package controller

import "github.com/gin-gonic/gin"

type EnterpriseController interface {
	Ping(ctx *gin.Context)
	CreateEnterpriseHandler(ctx *gin.Context)
	ReadEnterprisesHandler(ctx *gin.Context)
	ReadEnterpriseHandler(ctx *gin.Context)
}
