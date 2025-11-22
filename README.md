# pretty_slog Usage

A colorized slog handler with automatic source tracking for errors.

## Installation
```bash
go get github.com/fl4vis/pretty_slog
```

## Basic Usage

```go
package main

import (
    "log/slog"
    "github.com/fl4vis/pretty_slog"
)

func main() {
    logger := slog.New(pretty_slog.NewHandler(&slog.HandlerOptions{
        Level:     slog.LevelDebug,
        AddSource: false, // Must be false for error-only source tracking
    }))
    
    slog.SetDefault(logger)
    
    slog.Debug("Debug message")
    slog.Info("Info message")
    slog.Warn("Warning message")
    slog.Error("Error message") // Only this shows source info
}

```
<br>

### Important: AddSource Setting

Set `AddSource: false` when creating the handler. The handler automatically adds source information (file, line, function) only for *ERROR* level logs.

If you set `AddSource: true`, source information will appear on all log levels, which is usually not desired.


Example Output

```
[12:34:56.789] DEBUG: Debug message {}
[12:34:56.790] INFO: Info message {}
[12:34:56.791] WARN: Warning message {}
[12:34:56.792] ERROR: Error message {
  "source": {
    "file": "/path/to/main.go",
    "function": "main.main",
    "line": 18
  }
}
```
<br>

## Features

- Color-coded log levels (DEBUG=gray, INFO=cyan, WARN=yellow, ERROR=red)
- Colorized JSON attribute output
- Automatic source tracking on errors only
- Clean, readable format with timestamps

<br>

## Tests

```bash
 go test -v -run Visual
 go test -v -run Debug
```
