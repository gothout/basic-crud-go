package model

import "time"

type Enterprise struct {
	Id       int64
	Name     string
	Cnpj     string
	Active   bool
	CreateAt time.Time
	UpdateAt time.Time
}
