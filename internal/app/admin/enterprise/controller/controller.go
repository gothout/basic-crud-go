package controller

import "github.com/gin-gonic/gin"

type EnterpriseController interface {
	CreateEnterpriseHandler(ctx *gin.Context)
	ReadEnterprisesHandler(ctx *gin.Context)
	ReadEnterpriseHandler(ctx *gin.Context)
	UpdateEnterpriseHandler(ctx *gin.Context)
	DeleteEnterpriseHandler(ctx *gin.Context)
}
