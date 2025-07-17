package controller

import "github.com/gin-gonic/gin"

type PermissionController interface {
	Read(ctx *gin.Context)
}
