package controller

import "github.com/gin-gonic/gin"

type UserController interface {
	Ping(ctx *gin.Context)
	CreateUserHandler(ctx *gin.Context)
}
