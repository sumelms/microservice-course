package logger

import (
	"fmt"
	"os"

	kitlog "github.com/go-kit/log"
)

type Logger interface {
	Log(keyvals ...interface{})
	With(keyvals ...interface{}) Logger
	Logger() kitlog.Logger
}

type logger struct {
	logger kitlog.Logger
}

func NewLogger() Logger {
	l := kitlog.NewLogfmtLogger(os.Stderr)
	l = kitlog.NewSyncLogger(l)
	l = kitlog.With(l,
		"service", os.Args[0],
		"time:", kitlog.DefaultTimestampUTC,
		"caller", kitlog.DefaultCaller,
	)

	return logger{logger: l}
}

func (l logger) Log(keyvals ...interface{}) {
	if err := l.logger.Log(keyvals...); err != nil {
		fmt.Println("Erro de log:", err)
	}
}

func (l logger) With(keyvals ...interface{}) Logger {
	return &logger{logger: kitlog.With(l.logger, keyvals...)}
}

func (l logger) Logger() kitlog.Logger {
	return l.logger
}
