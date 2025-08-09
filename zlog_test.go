package zlog

import (
	"bytes"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newTestLogger() (*bytes.Buffer, *zap.Logger) {
	var buf bytes.Buffer
	return &buf, New(zapcore.DebugLevel, &buf)
}

func TestNew(t *testing.T) {
	_, logger := newTestLogger()
	if logger == nil {
		t.Fatal("Expected non-nil logger")
	}
}

func TestNewStdOut(t *testing.T) {
	logger := NewStdOut(zapcore.DebugLevel)
	if logger == nil {
		t.Fatal("Expected non-nil logger")
	}
}

func TestSetLogger(t *testing.T) {
	_, logger := newTestLogger()
	SetLogger(logger)
	if Logger() != logger {
		t.Fatal("Expected global logger to be set")
	}
	if Sugared() == nil {
		t.Fatal("Expected global sugared logger to be set")
	}
}

func TestSetLevel(t *testing.T) {
	SetLevel(zapcore.InfoLevel)
	if defaultLevel.Level() != zapcore.InfoLevel {
		t.Fatalf("Expected log level to be %v, got %v", zapcore.InfoLevel, defaultLevel.Level())
	}
}

func TestLoggerAndSugared(t *testing.T) {
	t.Run("Logger", func(t *testing.T) {
		if Logger() == nil {
			t.Fatal("Expected non-nil global logger")
		}
	})
	t.Run("Sugared", func(t *testing.T) {
		if Sugared() == nil {
			t.Fatal("Expected non-nil global sugared logger")
		}
	})
}

func TestClose(t *testing.T) {
	// Should not panic
	Close()
}
