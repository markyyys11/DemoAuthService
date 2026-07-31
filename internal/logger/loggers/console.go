package loggers

import (
	"DemoAuthService/internal/logger"
	"log/slog"
	"os"
)

type console struct {
}

func NewConsoleLogger() logger.Logger {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))
	return &console{}
}

func (l *console) Hint(msg string) {
	slog.Default().Info(msg)
}

func (l *console) Warn(msg string) {
	slog.Default().Warn(msg)
}

func (l *console) Err(msg string) {
	slog.Default().Error(msg)
}
