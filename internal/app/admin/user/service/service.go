package service

import (
	"context"
)

type UserService interface {
	Ping(ctx context.Context) (string, error)
}
