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

func New(level zapcore.LevelEnabler, w io.Writer, options ...zap.Option) *zap.Logger {
	if level == nil {
		level = defaultLevel
	}

	cfg := zapcore.EncoderConfig{
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

	enc := zapcore.NewJSONEncoder(cfg)
	return zap.New(zapcore.NewCore(enc, zapcore.AddSync(w), level),
		options...)
}

func NewStdOut(level zapcore.LevelEnabler, options ...zap.Option) *zap.Logger {
	return New(level, os.Stdout, options...)
}

func SetLogger(l *zap.Logger) {
	global = l
	globalSugared = l.Sugar()
}

func Logger() *zap.Logger {
	return global
}

func SetLevel(l zapcore.Level) {
	defaultLevel.SetLevel(l)
}

func Sugared() *zap.SugaredLogger {
	return globalSugared
}

func Close() {
	_ = Logger().Sync()
	_ = Sugared().Sync()
}
