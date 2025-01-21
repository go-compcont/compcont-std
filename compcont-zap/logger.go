package compcontzap

import (
	"context"
	"net/http"

	"go.uber.org/zap"
)

var (
	defaultLogger *zap.Logger
)

type ctxKeyLogger struct{}

func WithRequest(req *http.Request, logger *zap.Logger) {
	*req = *req.WithContext(WithContext(req.Context(), logger))
}

func WithContext(ctx context.Context, logger *zap.Logger) context.Context {
	// ctx中已经有一个logger了，不再添加新logger
	if _, ok := ctx.Value(ctxKeyLogger{}).(*zap.Logger); ok {
		return ctx
	}
	return context.WithValue(ctx, ctxKeyLogger{}, logger)
}

func FromContext(ctx context.Context) *zap.Logger {
	val, ok := ctx.Value(ctxKeyLogger{}).(*zap.Logger)
	if !ok {
		return GetDefault()
	}
	return val
}

func SetDefault(l *zap.Logger) {
	defaultLogger = l
}

func GetDefault() *zap.Logger {
	if defaultLogger == nil {
		defaultLogger, _ = zap.NewDevelopmentConfig().Build()
	}
	return defaultLogger
}
