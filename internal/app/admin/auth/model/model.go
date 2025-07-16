package model

import (
	"time"
)

type Token struct {
	Id        int64
	UserId    string
	Token     string
	CreatedAt time.Time
}
