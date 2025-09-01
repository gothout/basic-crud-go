package util

import (
	"basic-crud-go/internal/app/admin/middleware/model"
	"basic-crud-go/internal/configuration/logger"
	"fmt"
	"github.com/gin-gonic/gin"
)

const (
	CtxIdentityKey = "identity"
	module         = "Admin-Middleware-Util"
)

func GetIdentity(ctx *gin.Context) (*model.UserIndentity, error) {
	v, ok := ctx.Get(CtxIdentityKey)
	if !ok || v == nil {
		logger.Log(logger.Error, module, "GetIdentity", "Missing identity in context")
		return nil, fmt.Errorf("Missing identity in context")
	}
	id, ok := v.(*model.UserIndentity)
	if !ok || id == nil {
		logger.Log(logger.Error, module, "GetIdentity", "Invalid identity in context")
		return nil, fmt.Errorf("Invalid identity in context")
	}
	return id, nil
}

func HasPermission(perms *[]model.UserPermissions, code string) bool {
	if perms == nil {
		return false
	}
	for _, p := range *perms {
		if p.Permission == nil {
			continue
		}
		// p.Permission é *permModel.Permission
		if p.Permission.Code == code { // ou (*p.Permission).Code == code
			return true
		}
	}
	return false
}
