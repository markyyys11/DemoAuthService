package profile

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Service interface {
	GetUserInfo(ctx context.Context) error
}

func NewService(repo *pgx.Conn) Service {
	return &service{
		repo: repo,
	}
}

type service struct {
	repo *pgx.Conn
}

func (s *service) GetUserInfo(ctx context.Context) error {
	return nil
}
