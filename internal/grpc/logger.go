package grpc

import (
	"log/slog"
	"os"
)

type Logger struct {
	*slog.Logger
}

func NewLogger(level string) *Logger {
	var logLevel slog.Level
	switch level {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		Level: logLevel,
	}

	handler := slog.NewTextHandler(os.Stdout, opts)
	return &Logger{slog.New(handler)}
}

func (l *Logger) WithGroup(name string) *Logger {
	return &Logger{l.Logger.WithGroup(name)}
}

func (l *Logger) WithAttrs(attrs ...slog.Attr) *Logger {
	return &Logger{l.Logger.WithGroup("context").With(convertAttrsToAny(attrs)...)}
}

func convertAttrsToAny(attrs []slog.Attr) []any {
	result := make([]any, 0, len(attrs)*2)
	for _, attr := range attrs {
		result = append(result, attr.Key, attr.Value)
	}
	return result
}
