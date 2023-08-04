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
	"github.com/go-mumu/cs-go/library/common/consts"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc/metadata"
	"io"
)

type CommonHandler struct {
	slog.Handler
}

type CommonHandlerOpts struct {
	Opts slog.HandlerOptions
}

func (h *CommonHandler) Handle(ctx context.Context, r slog.Record) error {
	traceId := ""

	if outMd, ok := metadata.FromOutgoingContext(ctx); ok {
		if sliceTraceId := outMd.Get(consts.TraceId); len(sliceTraceId) > 0 {
			traceId = sliceTraceId[0]
		}
	}

	if inMd, ok := metadata.FromIncomingContext(ctx); ok {
		if sliceTraceId := inMd.Get(consts.TraceId); len(sliceTraceId) > 0 {
			traceId = sliceTraceId[0]
		}
	}

	if traceId != "" {
		r.Add(consts.TraceId, traceId)
	}

	err := h.Handler.Handle(ctx, r)
	if err != nil {
		return err
	}

	return nil
}

func NewCommonHandler(out io.Writer, opts CommonHandlerOpts) *CommonHandler {
	return &CommonHandler{
		Handler: slog.NewTextHandler(out, &opts.Opts),
	}
}
