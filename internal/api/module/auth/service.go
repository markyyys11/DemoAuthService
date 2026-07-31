package auth

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5"
)

type Service interface {
	Register(ctx context.Context) error
	Login(ctx context.Context) error
	Logout(ctx context.Context) error
}

func NewService(repo *pgx.Conn) Service {
	return &service{}
}

type service struct {
	repo *sql.Conn
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
