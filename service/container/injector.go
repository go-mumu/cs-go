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
	"github.com/go-mumu/cs-go/library/mysql"
	"github.com/go-mumu/cs-go/service/container/provider"
	"github.com/google/wire"
)

// App 全局应用程序
type App struct {
	DefMysql *mysql.DefMysql
	Server   *provider.Server
	Handler  *provider.Handler
}

// InitApp 注入全局应用程序
func InitApp() (*App, func(), error) {
	panic(
		wire.Build(
			wire.Struct(new(App), "*"),
			mysql.InitDef,
			provider.NewServer,
			provider.HandlerProviderSet,
			provider.DaoProviderSet,
		),
	)
}

/*func (a *App) Run() error {
	a.Server.SetGrpcAddr(config.V.GetString("rpc.grpc_addr"))
	a.Server.SetHttpAddr(config.V.GetString("rpc.http_addr"))

	a.Server.SetGrpcHandlerTimeout(config.V.GetInt("rpc.grpc_handler_timeout"))

	a.Server.SetHttpReadTimeout(config.V.GetInt("rpc.http_read_timeout"))
	a.Server.SetHttpWriteTimeout(config.V.GetInt("rpc.http_write_timeout"))

	a.Server.SetGrpcIdleTimeout(config.V.GetInt("rpc.grpc_idle_timeout"))
	a.Server.SetHttpIdleTimeout(config.V.GetInt("rpc.http_idle_timeout"))

	a.Server.SetMaxBodySize(config.V.GetInt("rpc.max_body_size"))

	a.Server.SetGrpcRegister(func(s *grpc.Server) {
		pb.RegisterUserServiceServer(s, a.Handler.UserServiceHandler)
	})

	a.Server.SetHttpRegister(func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
		return server.HttpRegisterFunc(ctx, mux, endpoint, opts,
			[]server.HttpRegister{
				pb.RegisterUserServiceHandlerFromEndpoint,
			}...,
		)
	})

	return a.Server.Run()
}*/
