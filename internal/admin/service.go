package admin

import (
	"context"
	"qolibaba/internal/admin/port"
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) SayHello(ctx context.Context, name string) (string, error) {
	return "admin: hello " + name, nil
}
