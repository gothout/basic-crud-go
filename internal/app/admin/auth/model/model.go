package model

import (
	entModel "basic-crud-go/internal/app/admin/enterprise/model"
	permModel "basic-crud-go/internal/app/admin/permission/model"
	userModel "basic-crud-go/internal/app/admin/user/model"
	"time"
)

type TokenUser struct {
	Id        int64
	UserId    string
	Token     string
	CreatedAt time.Time
}

type TokenApi struct {
	Id        int64
	UserId    string
	Token     string
	CreatedAt time.Time
	ExpiresAt time.Time
}

type UserIdentity struct {
	User        *userModel.User
	Enterprise  *entModel.Enterprise
	Permissions *[]permModel.Permission
	TokenUser   *TokenUser
	TokenApi    *TokenApi
}
