//go:build wireinject

/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-07-26
 * File: injector.go
 * Desc: 注入器
 */

package container

import (
	"context"
	"github.com/go-mumu/cs-go/library/common/flags"
	"github.com/go-mumu/cs-go/library/config"
	"github.com/go-mumu/cs-go/library/mysql"
	"github.com/go-mumu/cs-go/proto/pb"
	"github.com/go-mumu/cs-go/service/container/provider"
	"github.com/go-mumu/cs-go/service/server"
	"github.com/google/wire"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"strconv"
)

// App 全局应用程序
type App struct {
	DefMysql *mysql.DefMysql
	Server   *server.Server
	Handler  *provider.Handler
}

// InitApp 注入全局应用程序
func InitApp() (*App, func(), error) {
	panic(
		wire.Build(
			wire.Struct(new(App), "*"),
			provider.MysqlProviderSet,
			provider.ServerProviderSet,
			provider.HandlerProviderSet,
			provider.DaoProviderSet,
		),
	)
}

func (app *App) Run() error {
	app.Server.SetGRPCIp(flags.GRPCServiceIp)
	app.Server.SetGRPCPort(strconv.FormatInt(int64(flags.GRPCServicePort), 10))
	app.Server.SetHTTPIp(flags.HTTPServiceIp)
	app.Server.SetHTTPPort(strconv.FormatInt(int64(flags.HTTPServicePort), 10))

	app.Server.SetGRPCHandlerTimeout(config.V.GetInt("server.grpc_handler_timeout"))
	app.Server.SetHTTPHandlerTimeout(config.V.GetInt("server.http_handler_timeout"))

	app.Server.SetHTTPReadTimeout(config.V.GetInt("server.http_read_timeout"))
	app.Server.SetHTTPWriteTimeout(config.V.GetInt("server.http_write_timeout"))

	app.Server.SetMaxConnectionIdle(config.V.GetInt("server.max_connection_idle"))
	app.Server.SetHTTPIdleTimeout(config.V.GetInt("server.http_idle_time_out"))

	app.Server.SetMaxMsgSize(config.V.GetInt("server.max_msg_size_byte"))

	app.Server.SetGRPCRegister(func(s *grpc.Server) {
		pb.RegisterUserServiceServer(s, app.Handler.UserServiceHandler)
	})

	app.Server.SetHTTPRegister(func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
		return server.HTTPRegisterFunc(ctx, mux, endpoint, opts,
			[]server.HTTPRegister{
				pb.RegisterUserServiceHandlerFromEndpoint,
			}...,
		)
	})

	return app.Server.Run()
}
