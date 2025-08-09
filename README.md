# zlog

`zlog` is a logging library for Go, built on top of `zap` from Uber. It provides a simple interface for structured logging with different levels and outputs.

## Installation

To install `zlog`, use `go get`:

```sh
go get github.com/gor0utine/zlog
```

## Usage

```go
import (
    "github.com/gor0utine/zlog"
    "go.uber.org/zap/zapcore"
    "os"
)

func main() {
    // Initialize logger with desired level and output
    logger := zlog.New(zlog.Config{
        Level:   zapcore.DebugLevel,
        Output:  os.Stdout,
        // ...other config options...
    })
    zlog.SetLogger(logger)

    // Simple log
    zlog.Info("info log one",
        "fieldOne", "test",
        "fieldTwo", "test2",
    )

    // Structured logging
    zlog.Debug("debug message", "user", "alice", "id", 42)
    zlog.Error("error occurred", "err", err)

    defer zlog.Close()
}
```

### Configuration

You can now configure `zlog` using the `zlog.Config` struct, which supports:
- `Level`: log level (e.g., `zapcore.InfoLevel`)
- `Output`: output destination (e.g., `os.Stdout`, file, etc.)
- Additional options as needed

## Testing

```sh
go test ./...
```
