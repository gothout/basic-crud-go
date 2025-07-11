package model

import "time"

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
