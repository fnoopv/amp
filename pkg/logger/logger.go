package logger

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"sync"

	"goyave.dev/goyave/v5/util/errors"
)

type CustomHandlerOptions struct {
	Level slog.Leveler
}

type CustomHandler struct {
	opts   *CustomHandlerOptions
	mu     *sync.Mutex
	w      io.Writer
	attrs  []slog.Attr
	groups []string
}

func NewCustomSlogHandler(w io.Writer, opts *CustomHandlerOptions) *CustomHandler {
	if opts == nil {
		opts = &CustomHandlerOptions{}
	}
	return &CustomHandler{
		w:    w,
		mu:   &sync.Mutex{},
		opts: opts,
	}
}

func (h *CustomHandler) Handle(_ context.Context, r slog.Record) error {
	buf := bytes.NewBuffer(make([]byte, 0, 1024))

	buf.WriteString(r.Level.String())
	buf.WriteRune(' ')
	buf.WriteString(r.Message)

	// Add attrs and groups...

	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := h.w.Write(buf.Bytes())
	return errors.New(err)
}

func (h *CustomHandler) Enabled(_ context.Context, level slog.Level) bool {
	minLevel := slog.LevelInfo
	if h.opts.Level != nil {
		minLevel = h.opts.Level.Level()
	}
	return level >= minLevel
}

func (h *CustomHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	newAttrs := make([]slog.Attr, 0, len(h.attrs)+len(attrs))
	newAttrs = append(newAttrs, h.attrs...)
	newAttrs = append(newAttrs, attrs...)
	return &CustomHandler{
		opts:   h.opts,
		w:      h.w,
		mu:     h.mu,
		attrs:  newAttrs,
		groups: h.groups,
	}
}

func (h *CustomHandler) WithGroup(name string) slog.Handler {
	return &CustomHandler{
		opts:   h.opts,
		w:      h.w,
		mu:     h.mu,
		attrs:  append(make([]slog.Attr, 0, len(h.attrs)), h.attrs...),
		groups: append(h.groups, name),
	}
}
