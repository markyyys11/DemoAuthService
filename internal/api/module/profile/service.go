package profile

import (
	"context"
	"database/sql"
)

type Service interface {
	GetUserInfo(ctx context.Context) error
}

func NewService(repo *sql.Conn) Service {
	return &service{
		repo: repo,
	}
}

type service struct {
	repo *sql.Conn
}

func (s *service) GetUserInfo(ctx context.Context) error {
	return nil
}
