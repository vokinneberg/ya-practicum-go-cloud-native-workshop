package logging

import (
	"context"

	"go.uber.org/zap"
)

type ctxKey string

const (
	Key ctxKey = "ZapLogger"
)

// CtxWithLogger returns a new context that stores the logger provided.
func CtxWithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, Key, logger)
}

// FromContext returns the logger from the context.
func FromContext(ctx context.Context) *zap.Logger {
	logger, ok := ctx.Value(Key).(*zap.Logger)
	if !ok {
		return zap.L()
	}
	return logger
}
