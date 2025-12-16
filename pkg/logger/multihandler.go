package logger

import (
	"context"
	"log/slog"
)

type MultiHandler struct {
    handlers []slog.Handler
}

func (m *MultiHandler) Enabled(ctx context.Context, level slog.Level) bool {
    for _, h := range m.handlers {
        if h.Enabled(ctx, level) {
            return true
        }
    }
    return false
}

func (m *MultiHandler) Handle(ctx context.Context, r slog.Record) error {
    var err error
    for _, h := range m.handlers {
        if e := h.Handle(ctx, r); e != nil {
            err = e
        }
    }
    return err
}

func (m *MultiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
    hs := make([]slog.Handler, 0, len(m.handlers))
    for _, h := range m.handlers {
        hs = append(hs, h.WithAttrs(attrs))
    }
    return &MultiHandler{handlers: hs}
}

func (m *MultiHandler) WithGroup(name string) slog.Handler {
    hs := make([]slog.Handler, 0, len(m.handlers))
    for _, h := range m.handlers {
        hs = append(hs, h.WithGroup(name))
    }
    return &MultiHandler{handlers: hs}
}
