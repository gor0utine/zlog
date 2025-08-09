package zlog

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	global        *zap.Logger
	globalSugared *zap.SugaredLogger
	defaultLevel  = zap.NewAtomicLevelAt(zap.ErrorLevel)
)

func init() {
	SetLogger(NewStdOut(defaultLevel))
}

// encoderConfig returns the default zapcore.EncoderConfig.
func encoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

// New creates a new zap.Logger with the given level and writer.
func New(level zapcore.LevelEnabler, w io.Writer, options ...zap.Option) *zap.Logger {
	if level == nil {
		level = defaultLevel
	}
	enc := zapcore.NewJSONEncoder(encoderConfig())
	core := zapcore.NewCore(enc, zapcore.AddSync(w), level)
	return zap.New(core, options...)
}

// NewStdOut creates a new zap.Logger that writes to stdout.
func NewStdOut(level zapcore.LevelEnabler, options ...zap.Option) *zap.Logger {
	return New(level, os.Stdout, options...)
}

// SetLogger sets the global logger and its sugared version.
func SetLogger(l *zap.Logger) {
	global = l
	globalSugared = l.Sugar()
}

// Logger returns the global logger.
func Logger() *zap.Logger {
	return global
}

// SetLevel sets the default logging level.
func SetLevel(l zapcore.Level) {
	defaultLevel.SetLevel(l)
}

// Sugared returns the global sugared logger.
func Sugared() *zap.SugaredLogger {
	return globalSugared
}

// Close flushes any buffered log entries.
func Close() {
	_ = global.Sync()
	_ = globalSugared.Sync()
}
