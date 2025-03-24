package main

import (
	"log/slog"

	"github.com/fl4vis/pretty_slog/pretty"
)

func main() {
	// logger := slog.New(pretty.NewHandler(nil))
	logger := slog.New(pretty.NewHandler(&slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))

	logger.Info("Hello World", "hi", "hi", "hi", "i", "h", "h")
	logger.Error("Hello World", "hi", "hi", "hi", "i")
}
