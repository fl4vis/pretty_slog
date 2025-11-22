package pretty_slog

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"runtime"
	"sync"
)

const TIME_FORMAT = "[15:04:05.000]"

type Handler struct {
	h slog.Handler
	b *bytes.Buffer
	m *sync.Mutex
}

func NewHandler(opts *slog.HandlerOptions) *Handler {
	if opts == nil {
		opts = &slog.HandlerOptions{}
	}
	b := &bytes.Buffer{}
	return &Handler{
		b: b,
		h: slog.NewJSONHandler(b, &slog.HandlerOptions{
			Level:       opts.Level,
			AddSource:   opts.AddSource,
			ReplaceAttr: suppressDefaults(opts.ReplaceAttr),
		}),
		m: &sync.Mutex{},
	}
}

func (h *Handler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.h.Enabled(ctx, level)
}

func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &Handler{h.h.WithAttrs(attrs), h.b, h.m}
}

func (h *Handler) WithGroup(name string) slog.Handler {
	return &Handler{h.h.WithGroup(name), h.b, h.m}
}

func (h *Handler) Handle(ctx context.Context, r slog.Record) error {
	level := r.Level.String() + ":"

	switch r.Level {
	case slog.LevelDebug:
		level = Colorize(GRAY, level)
	case slog.LevelInfo:
		level = Colorize(CYAN, level)
	case slog.LevelWarn:
		level = Colorize(YELLOW, level)
	case slog.LevelError:
		level = Colorize(RED, level)
	}

	attrs, err := h.computeAttrs(ctx, r)
	if err != nil {
		return err
	}

	// If Error log, add file and line info
	if r.Level == slog.LevelError {
		pc, file, line, ok := runtime.Caller(3) // Adjust depth as needed
		if ok {
			fn := runtime.FuncForPC(pc) // Get function name
			attrs["source"] = map[string]any{
				"file":     file,
				"function": fn.Name(),
				"line":     line,
			}
		}
	}

	// Extract custom attributes from the log call
	r.Attrs(func(a slog.Attr) bool {
		attrs[a.Key] = a.Value.Any()
		return true
	})

	bytes, err := json.MarshalIndent(attrs, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshal attributes: %w", err)
	}

	fmt.Println(
		Colorize(PURPLE, r.Time.Format(TIME_FORMAT)),
		level,
		Colorize(WHITE, r.Message),
		ColorizeJSON(bytes),
	)

	return nil
}

func (h *Handler) computeAttrs(
	ctx context.Context,
	r slog.Record,
) (map[string]any, error) {
	h.m.Lock()
	defer func() {
		h.b.Reset()
		h.m.Unlock()
	}()

	if err := h.h.Handle(ctx, r); err != nil {
		return nil, fmt.Errorf("error on inner handler's Handle: %w", err)
	}

	var attrs map[string]any
	err := json.Unmarshal(h.b.Bytes(), &attrs)
	if err != nil {
		return nil, fmt.Errorf("error on inner handler's Unmarshal: %w", err)
	}

	return attrs, nil
}

func suppressDefaults(next func([]string, slog.Attr) slog.Attr) func([]string, slog.Attr) slog.Attr {
	return func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.TimeKey ||
			a.Key == slog.LevelKey ||
			a.Key == slog.MessageKey {
			return slog.Attr{}
		}
		if next == nil {
			return a
		}
		return next(groups, a)
	}
}
