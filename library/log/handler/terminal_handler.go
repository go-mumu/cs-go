/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-07-25
 * File: terminal_handler.go
 * Desc: 终端处理
 */

package handler

import (
	"context"
	"encoding/json"
	"github.com/fatih/color"
	"golang.org/x/exp/slog"
	"io"
	"log"
)

type TerminalHandlerOptions struct {
	Opts slog.HandlerOptions
}

type TerminalHandler struct {
	slog.Handler
	l *log.Logger
}

func (h *TerminalHandler) Handle(ctx context.Context, r slog.Record) error {
	level := r.Level.String() + ":"

	switch r.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	fields := make(map[string]interface{}, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()

		return true
	})

	b, err := json.Marshal(fields)
	if err != nil {
		return err
	}

	timeStr := r.Time.Format("[2006-01-02 15:04:05.000]")
	timeStr = color.GreenString(timeStr)

	msg := color.CyanString(r.Message)

	h.l.Println(timeStr, level, msg, color.WhiteString(string(b)))

	return nil
}

func NewTerminalHandler(
	out io.Writer,
	opts TerminalHandlerOptions,
) *TerminalHandler {
	h := &TerminalHandler{
		Handler: slog.NewJSONHandler(out, &opts.Opts),
		l:       log.New(out, "", 0),
	}

	return h
}
