/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-07-25
 * File: slog.go
 * Desc:
 */

package log

import (
	"github.com/go-mumu/cs-go/log/handler"
	"github.com/go-mumu/cs-go/log/writer"
	"golang.org/x/exp/slog"
	"os"
)

var Log *slog.Logger
var Cli *slog.Logger

func init() {
	Log = slog.New(
		handler.NewCommonHandler(
			writer.FileWriter(),
			handler.CommonHandlerOpts{
				Opts: slog.HandlerOptions{
					AddSource: true,
				},
			},
		),
	)

	Cli = slog.New(
		handler.NewTerminalHandler(
			os.Stdout,
			handler.TerminalHandlerOptions{
				Opts: slog.HandlerOptions{
					Level: slog.LevelDebug,
				},
			},
		),
	)
}
