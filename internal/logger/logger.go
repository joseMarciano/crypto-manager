package logger

import (
	"context"
	"log/slog"
	"os"
	"sync"

	"github.com/joseMarciano/crypto-manager/internal/observability"
)

var log *slog.Logger
var once sync.Once

func Info(ctx context.Context, msg string, data ...any) {
	logger(ctx).InfoContext(ctx, msg, data...)
}

func Debug(ctx context.Context, msg string, data ...any) {
	logger(ctx).DebugContext(ctx, msg, data...)
}

func Error(ctx context.Context, msg string, data ...any) {
	logger(ctx).ErrorContext(ctx, msg, data...)
}

func Warn(ctx context.Context, msg string, data ...any) {
	logger(ctx).WarnContext(ctx, msg, data...)
}

func logger(ctx context.Context) *slog.Logger {
	l := setup()
	reqID, ok := observability.RequestIDFromContext(ctx)
	if ok {
		l = l.With(observability.RequestIDMetadataKey, reqID)
	}

	return l
}

func setup() *slog.Logger {
	once.Do(func() {
		log = slog.New(slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{Level: new(slog.LevelVar)},
		))
	})

	return log
}
