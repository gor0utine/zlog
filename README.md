# zlog

`zlog` is a logging library for Go, built on top of `zap` from Uber. It provides a simple interface for logging with different levels and outputs.

## Installation

To install `zlog`, use `go get`:

```sh
go get github.com/pr1sonmike/zlog
```

## Usage

```go
import (
    "github.com/pr1sonmike/zlog"
    "go.uber.org/zap/zapcore"
)

func main() {
    logger := zlog.New(zapcore.DebugLevel, os.Stdout)
    zlog.SetLogger(logger)
	
	zlog.Info("info log one",
        "fieldOne", "test",
        "fieldTwo", "test2")

    defer zlog.Close()
}
```

## Testing

```sh
go test ./...
```