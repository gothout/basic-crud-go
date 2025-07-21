package controller

import "github.com/gin-gonic/gin"

type PermissionController interface {
	Search(ctx *gin.Context)
	Read(ctx *gin.Context)
	ReadAll(ctx *gin.Context)
}
