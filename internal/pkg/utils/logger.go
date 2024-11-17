package utils

import (
	"context"
	"io"
	"os"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/tshubham7/eth-parser/internal/pkg/constants"
)

type contextKey string

const (
	KeyLogger contextKey = "LOGGER"

	KeyCorrelationId = "CORRELATION_ID"
)

func NewLogger(ctx context.Context) (context.Context, *logrus.Logger) {
	var log = logrus.New()
	log.Formatter = new(logrus.JSONFormatter)
	log.Level = getLogLevel(os.Getenv(constants.EnvLogLevel))
	log.Out = os.Stdout

	log = log.WithField(KeyCorrelationId, uuid.NewString()).Logger

	ctx = context.WithValue(ctx, KeyLogger, log)
	return ctx, log
}

func NewSilentLogger(ctx context.Context) (context.Context, *logrus.Logger) {
	var log = logrus.New()
	log.Formatter = new(logrus.JSONFormatter)
	log.Level = getLogLevel(os.Getenv(constants.EnvLogLevel))
	log.Out = io.Discard

	log = log.WithField(KeyCorrelationId, uuid.NewString()).Logger

	ctx = context.WithValue(ctx, KeyLogger, log)
	return ctx, log
}

func GetCurrentLogger(ctx context.Context) *logrus.Logger {
	if log := ctx.Value(KeyLogger); log != nil {
		return log.(*logrus.Logger)
	}

	_, log := NewLogger(ctx)
	return log
}

func getLogLevel(l string) logrus.Level {
	switch l {
	case "INFO":
		return logrus.InfoLevel
	case "DEBUG":
		return logrus.DebugLevel
	default:
		return logrus.TraceLevel
	}
}
