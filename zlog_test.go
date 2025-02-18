package zlog

import (
	"bytes"
	"testing"

	"go.uber.org/zap/zapcore"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	logger := New(zapcore.DebugLevel, &buf)
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
	var buf bytes.Buffer
	logger := New(zapcore.DebugLevel, &buf)
	SetLogger(logger)
	if global != logger {
		t.Fatal("Expected global logger to be set")
	}
	if globalSugared == nil {
		t.Fatal("Expected global sugared logger to be set")
	}
}

func TestSetLevel(t *testing.T) {
	SetLevel(zapcore.InfoLevel)
	if defaultLevel.Level() != zapcore.InfoLevel {
		t.Fatalf("Expected log level to be %v, got %v", zapcore.InfoLevel, defaultLevel.Level())
	}
}

func TestLogger(t *testing.T) {
	if Logger() == nil {
		t.Fatal("Expected non-nil global logger")
	}
}

func TestSugared(t *testing.T) {
	if Sugared() == nil {
		t.Fatal("Expected non-nil global sugared logger")
	}
}

func TestClose(t *testing.T) {
	// Ensure no panic occurs
	Close()
}
