package slogctxadapter

import (
	"context"
	"log/slog"

	"github.com/teeaa/slogctx"
	sqldblogger "github.com/teeaa/sqldb-logger"
)

type slogctxAdapter struct {
	logger *slogctx.Logger
}

// New creates a log adapter from sqldblogger.Logger to an slog.Logger one.
func New(logger *slogctx.Logger) sqldblogger.Logger {
	return &slogctxAdapter{logger: logger}
}

// Log implement sqldblogger.Logger and converts its levels to corresponding
// log/slog ones.
func (a *slogctxAdapter) Log(ctx context.Context, sqldbLevel sqldblogger.Level, msg string, data map[string]interface{}) {
	attrs := make([]slog.Attr, 0, len(data))
	for k, v := range data {
		attrs = append(attrs, slog.Any(k, v))
	}

	var level slog.Level
	switch sqldbLevel {
	case sqldblogger.LevelError:
		level = slog.LevelError
	case sqldblogger.LevelInfo:
		level = slog.LevelInfo
	case sqldblogger.LevelDebug:
		level = slog.LevelDebug
	default:
		// trace will use slog debug
		level = slog.LevelDebug
	}

	a.logger.LogAttrs(ctx, level, msg, attrs...)
}
