package controller

import "github.com/gin-gonic/gin"

type PermissionController interface {
	ReadAll(ctx *gin.Context)
}
