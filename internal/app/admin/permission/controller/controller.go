package controller

import "github.com/gin-gonic/gin"

type PermissionController interface {
	Search(ctx *gin.Context)
	Read(ctx *gin.Context)
	Apply(ctx *gin.Context)
	ReadAll(ctx *gin.Context)
	ReadUserPermission(ctx *gin.Context)
	RemoveBatch(ctx *gin.Context)
}
