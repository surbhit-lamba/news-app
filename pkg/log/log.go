package log

import (
	"context"
	"news-app/pkg/log/logger"
)

// Initialize log module
func Initialize(ctx context.Context) {
	err := logger.Initialize(
		logger.Formatter("json"),
		logger.Level("debug"),
		logger.CallerSkip(2),
	)
	if err != nil {
		logger.Error("Error in initializing log module")
	}
}

// GetContextWithLogger Function adds Logger to the context so that request level fields could easily be logged
func GetContextWithLogger(parentContext context.Context, fields map[string]interface{}) context.Context {
	loggerObj := logger.WithFields(fields)
	return logger.ContextWithLogger(parentContext, loggerObj)
}

// DebugWithContext Debug logs a message at level Debug on the standard logger.
// DebugWithContext - requires logger to be added in context using log.ContextWithLogger
func DebugWithContext(ctx context.Context, args ...interface{}) {
	logger.FromContext(ctx).Debug(args...)
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// PrintWithContext Print logs a message at level Info on the standard logger.
// PrintWithContext - requires logger to be added in context using log.ContextWithLogger
func PrintWithContext(ctx context.Context, args ...interface{}) {
	logger.FromContext(ctx).Print(args...)
}

func Print(args ...interface{}) {
	logger.Print(args...)
}

// InfoWithContext Info logs a message at level Info on the standard logger.
// InfoWithContext - requires logger to be added in context using log.ContextWithLogger
func InfoWithContext(ctx context.Context, args ...interface{}) {
	logger.FromContext(ctx).Info(args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

// WarnWithContext Warn logs a message at level Warn on the standard logger.
// WarnWithContext - requires logger to be added in context using log.ContextWithLogger
func WarnWithContext(ctx context.Context, args ...interface{}) {
	logger.FromContext(ctx).Warn(args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// ErrorWithContext Error logs a message at level Error on the standard logger.
// ErrorWithContext - requires logger to be added in context using log.ContextWithLogger
func ErrorWithContext(ctx context.Context, args ...interface{}) {
	logger.FromContext(ctx).Error(args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

// PanicWithContext Panic logs a message at level Panic on the standard logger.
// PanicWithContext - requires logger to be added in context using log.ContextWithLogger
func PanicWithContext(ctx context.Context, args ...interface{}) {
	logger.FromContext(ctx).Panic(args...)
}

func Panic(args ...interface{}) {
	logger.Panic(args...)
}

// DebugfWithContext Debugf logs a message at level Debug on the standard logger.
// DebugfWithContext - requires logger to be added in context using log.ContextWithLogger
func DebugfWithContext(ctx context.Context, format string, args ...interface{}) {
	logger.FromContext(ctx).Debugf(format, args...)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

// PrintfWithContext Printf logs a message at level Info on the standard logger.
// PrintfWithContext - requires logger to be added in context using log.ContextWithLogger
func PrintfWithContext(ctx context.Context, format string, args ...interface{}) {
	logger.FromContext(ctx).Printf(format, args...)
}

func Printf(format string, args ...interface{}) {
	logger.Printf(format, args...)
}

// InfofWithContext Infof logs a message at level Info on the standard logger.
// InfofWithContext - requires logger to be added in context using log.ContextWithLogger
func InfofWithContext(ctx context.Context, format string, args ...interface{}) {
	logger.FromContext(ctx).Infof(format, args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// WarnfWithContext Warnf logs a message at level Warn on the standard logger.
// WarnfWithContext - requires logger to be added in context using log.ContextWithLogger
func WarnfWithContext(ctx context.Context, format string, args ...interface{}) {
	logger.FromContext(ctx).Warnf(format, args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

// ErrorfWithContext Errorf logs a message at level Error on the standard logger.
// ErrorfWithContext - requires logger to be added in context using log.ContextWithLogger
func ErrorfWithContext(ctx context.Context, format string, args ...interface{}) {
	logger.FromContext(ctx).Errorf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

// PanicfWithContext Panicf logs a message at level Panic on the standard logger.
// PanicfWithContext - requires logger to be added in context using log.ContextWithLogger
func PanicfWithContext(ctx context.Context, format string, args ...interface{}) {
	logger.FromContext(ctx).Panicf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}

// DebuglnWithContext Debugln logs a message at level Debug on the standard logger.
// DebuglnWithContext - requires logger to be added in context using log.ContextWithLogger
func DebuglnWithContext(ctx context.Context, args ...interface{}) {
	logger.FromContext(ctx).Debugln(args...)
}

func Debugln(args ...interface{}) {
	logger.Debugln(args...)
}

// PrintlnWithContext Println logs a message at level Info on the standard logger.
// PrintlnWithContext - requires logger to be added in context using log.ContextWithLogger
func PrintlnWithContext(ctx context.Context, args ...interface{}) {
	logger.FromContext(ctx).Println(args...)
}

func Println(args ...interface{}) {
	logger.Println(args...)
}

// InfolnWithContext Infoln logs a message at level Info on the standard logger.
// InfolnWithContext - requires logger to be added in context using log.ContextWithLogger
func InfolnWithContext(ctx context.Context, args ...interface{}) {
	logger.FromContext(ctx).Infoln(args...)
}

func Infoln(args ...interface{}) {
	logger.Infoln(args...)
}

// WarnlnWithContext Warnln logs a message at level Warn on the standard logger.
// WarnlnWithContext - requires logger to be added in context using log.ContextWithLogger
func WarnlnWithContext(ctx context.Context, args ...interface{}) {
	logger.FromContext(ctx).Warnln(args...)
}

func Warnln(args ...interface{}) {
	logger.Warnln(args...)
}

// WarninglnWithContext Warningln logs a message at level Warn on the standard logger.
// WarninglnWithContext - requires logger to be added in context using log.ContextWithLogger
func WarninglnWithContext(ctx context.Context, args ...interface{}) {
	logger.FromContext(ctx).Warnln(args...)
}

func Warningln(args ...interface{}) {
	logger.Warnln(args...)
}

// ErrorlnWithContext Errorln logs a message at level Error on the standard logger.
// ErrorlnWithContext - requires logger to be added in context using log.ContextWithLogger
func ErrorlnWithContext(ctx context.Context, args ...interface{}) {
	logger.FromContext(ctx).Errorln(args...)
}

func Errorln(args ...interface{}) {
	logger.Errorln(args...)
}

// PaniclnWithContext Panicln logs a message at level Panic on the standard logger.
// PaniclnWithContext - requires logger to be added in context using log.ContextWithLogger
func PaniclnWithContext(ctx context.Context, args ...interface{}) {
	logger.FromContext(ctx).Panicln(args...)
}

func Panicln(args ...interface{}) {
	logger.Panicln(args...)
}
