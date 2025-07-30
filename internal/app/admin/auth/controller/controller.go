package controller

import "github.com/gin-gonic/gin"

type AuthController interface {
	AuthLoginHandler(ctx *gin.Context)
	AuthRefreshHandler(ctx *gin.Context)
}
