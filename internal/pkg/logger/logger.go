package logger

import (
	"go.uber.org/zap"
)

// Logger - кастомная обертка над zap.Logger
type Logger struct {
	zap *zap.Logger
}

func New() *Logger {
	return &Logger{}
}

func (l *Logger) Error(msg string, err error) {

}
