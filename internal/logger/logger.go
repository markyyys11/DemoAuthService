package logger

import "fmt"

var def Logger

type Logger interface {
	Hint(msg string)
	Warn(msg string)
	Err(msg string)
}

func InitLogger(logger Logger) {
	def = logger
}

func Hint(msg string, a ...any) {
	def.Hint(fmt.Sprintf(msg, a...))
}

func Warn(msg string, a ...any) {
	def.Warn(fmt.Sprintf(msg, a...))
}

func Err(msg string, a ...any) {
	def.Err(fmt.Sprintf(msg, a...))
}
