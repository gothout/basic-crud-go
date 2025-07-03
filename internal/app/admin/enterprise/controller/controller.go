package controller

import "github.com/gin-gonic/gin"

type EnterpriseController interface {
	Ping(ctx *gin.Context)
}
