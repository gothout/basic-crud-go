package service

import "context"

type EnterpriseService interface {
	Ping(ctx context.Context) (string, error)
}
