package controller

import "github.com/gin-gonic/gin"

type AdminController interface {
	Ping(ctx *gin.Context)
}
