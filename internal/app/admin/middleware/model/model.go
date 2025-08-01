package model

import (
	entModel "basic-crud-go/internal/app/admin/enterprise/model"
	permModel "basic-crud-go/internal/app/admin/permission/model"
	userModel "basic-crud-go/internal/app/admin/user/model"
)

type UserIndentity struct {
	User        *userModel.User
	Enterprise  *entModel.Enterprise
	Permissions *[]UserPermissions
}

type UserPermissions struct {
	Permission *permModel.Permission
}
