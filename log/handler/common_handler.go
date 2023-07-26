/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-07-25
 * File: common_handler.go
 * Desc: 通用处理
 */

package handler

import (
	"context"
	"github.com/go-mumu/cs-go/common/consts"
	"golang.org/x/exp/slog"
	"io"
)

type CommonHandler struct {
	slog.Handler
}

type CommonHandlerOpts struct {
	Opts slog.HandlerOptions
}

func (h *CommonHandler) Handle(ctx context.Context, r slog.Record) error {
	r.Add(consts.TraceId, ctx.Value(consts.TraceId))

	err := h.Handler.Handle(ctx, r)
	if err != nil {
		return err
	}

	return nil
}

func NewCommonHandler(out io.Writer, opts CommonHandlerOpts) *CommonHandler {
	return &CommonHandler{
		Handler: slog.NewJSONHandler(out, &opts.Opts),
	}
}
