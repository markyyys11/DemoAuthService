package auth

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Service interface {
	Register(ctx context.Context) error
	Login(ctx context.Context) error
	Logout(ctx context.Context) error
	Refresh(ctx context.Context) error
}

func NewService(repo *pgx.Conn) Service {
	return &service{
		repo: repo,
	}
}

type service struct {
	repo *pgx.Conn
}

func (s *service) Register(ctx context.Context) error {
	return nil
}

func (s *service) Login(ctx context.Context) error {
	return nil
}

func (s *service) Logout(ctx context.Context) error {
	return nil
}

func (s *service) Refresh(ctx context.Context) error {
	return nil
}
