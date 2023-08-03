/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-07-25
 * File: slog.go
 * Desc: 初始化日志
 */

package log

import (
	handler2 "github.com/go-mumu/cs-go/library/log/handler"
	"github.com/go-mumu/cs-go/library/log/writer"
	"golang.org/x/exp/slog"
	"os"
)

var Log *slog.Logger
var Cli *slog.Logger

func init() {
	Log = slog.New(
		handler2.NewCommonHandler(
			writer.FileWriter(),
			handler2.CommonHandlerOpts{
				Opts: slog.HandlerOptions{
					AddSource: true,
				},
			},
		),
	)

	Cli = slog.New(
		handler2.NewTerminalHandler(
			os.Stdout,
			handler2.TerminalHandlerOptions{
				Opts: slog.HandlerOptions{
					Level: slog.LevelDebug,
				},
			},
		),
	)
}
