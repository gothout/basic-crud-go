package service

import (
	"context"
)

type AdminService interface {
	Ping(ctx context.Context) (string, error)
}
