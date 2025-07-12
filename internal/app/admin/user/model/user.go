package model

import (
	entModel "basic-crud-go/internal/app/admin/enterprise/model"
	"time"
)

type User struct {
	Id           string
	EnterpriseId int64
	Number       string
	FirstName    string
	LastName     string
	Email        string
	Password     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserExtend struct {
	User
	Enterprise entModel.Enterprise
}
