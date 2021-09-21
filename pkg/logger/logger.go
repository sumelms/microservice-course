package logger

import (
	"os"

	"github.com/go-kit/kit/log"
)

func NewLogger() log.Logger {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.NewSyncLogger(logger)
	logger = log.With(logger,
		"service", os.Args[0],
		"time:", log.DefaultTimestampUTC,
		"caller", log.DefaultCaller,
	)

	return logger
}
