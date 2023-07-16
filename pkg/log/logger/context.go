package logger

import (
	"context"

	"github.com/gin-gonic/gin"
)

type ctxKeyType string

var LoggerCtxKey ctxKeyType = "logger"
var LoggerCtxFieldsKey ctxKeyType = "logger.fields"

const LogLevelOverRideKey = "log_level_override"

// FromContext returns the logger in the context.
// If no logger is present in the context, it'll return the default Logger instance
func FromContext(ctx context.Context) *Logger {
	var l *Logger

	switch ctx := ctx.(type) {
	case *gin.Context:
		log, ok := ctx.Get(string(LoggerCtxKey))
		if ok {
			l = log.(*Logger)
		} else {
			l = defaultLogger
		}

		f, ok := ctx.Get(string(LoggerCtxFieldsKey))
		if ok {
			fields, ok := f.([]Field)
			if ok {
				l = updateLoggerForLevelField(fields).With(fields...)
			}
		}

	default:
		var ok bool
		l, ok = ctx.Value(LoggerCtxKey).(*Logger)
		if !ok {
			l = defaultLogger
		}

		fields, ok := ctx.Value(LoggerCtxFieldsKey).([]Field)
		if ok {
			l = updateLoggerForLevelField(fields).With(fields...)
		}
	}
	return l
}

// ContextWithLogger creates a new context derived from the provided context with a Logger instance
//
// Deprecated: Use ContextWithFields to embed logger fields into the context subtree instead
func ContextWithLogger(ctx context.Context, l *Logger) context.Context {
	switch ctx := ctx.(type) {
	case *gin.Context:
		ctx.Set(string(LoggerCtxKey), l)
		return ctx
	default:
		return context.WithValue(ctx, LoggerCtxKey, l)
	}
}

// ContextWithFields creates a new context derived from the provided context with a set of fields
// which can be later retrieved by FromContext
func ContextWithFields(ctx context.Context, fields ...Field) context.Context {
	switch ctx := ctx.(type) {
	case *gin.Context:
		ctx.Set(string(LoggerCtxFieldsKey), fields)
		return ctx
	default:
		return context.WithValue(ctx, LoggerCtxFieldsKey, fields)
	}
}

// Use a different logger instance if log_level_override field is present in context
func updateLoggerForLevelField(fields []Field) *Logger {
	for _, f := range fields {
		if f.Key == LogLevelOverRideKey {
			l, err := NewLogger(Level(f.String))
			if err != nil {
				return defaultLogger
			}

			return l
		}
	}

	return defaultLogger
}
