package controller

import "github.com/gin-gonic/gin"

type UserController interface {
	Ping(ctx *gin.Context)
	CreateUserHandler(ctx *gin.Context)
	ReadUserHandler(ctx *gin.Context)
	ReadUsersHandler(ctx *gin.Context)
	UpdateUserHandler(ctx *gin.Context)
}
