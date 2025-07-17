package model

import (
	userModel "basic-crud-go/internal/app/admin/user/model"
	"time"
)

type Token struct {
	Id        int64
	UserId    string
	Token     string
	CreatedAt time.Time
}

type User struct {
	Permission string
	User       userModel.UserExtend
}
