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
	"github.com/go-mumu/cs-go/library/config"
	"github.com/go-mumu/cs-go/library/mysql"
	"github.com/go-mumu/cs-go/proto/pb"
	provider2 "github.com/go-mumu/cs-go/service/container/provider"
	server2 "github.com/go-mumu/cs-go/service/server"
	"github.com/google/wire"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// App 全局应用程序
type App struct {
	DefMysql *mysql.DefMysql
	Server   *server2.Server
	Handler  *provider2.Handler
}

// InitApp 注入全局应用程序
func InitApp() (*App, func(), error) {
	panic(
		wire.Build(
			wire.Struct(new(App), "*"),
			mysql.InitDef,
			config.Init,
			server2.NewServer,
			provider2.HandlerProviderSet,
			provider2.DaoProviderSet,
		),
	)
}

func (a *App) Run() error {
	a.Server.SetGrpcAddr(config.C.Rpc.GrpcAddr)
	a.Server.SetHttpAddr(config.C.Rpc.HttpAddr)

	a.Server.SetGrpcHandlerTimeout(config.C.Rpc.GrpcHandlerTimeout)

	a.Server.SetHttpReadTimeout(config.C.Rpc.HttpReadTimeout)
	a.Server.SetHttpWriteTimeout(config.C.Rpc.HttpWriteTimeout)

	a.Server.SetGrpcIdleTimeout(config.C.Rpc.GrpcIdleTimeout)
	a.Server.SetHttpIdleTimeout(config.C.Rpc.HttpIdleTimeout)

	a.Server.SetMaxBodySize(config.C.Rpc.MaxBodySize)

	a.Server.SetGrpcRegister(func(s *grpc.Server) {
		pb.RegisterUserServiceServer(s, a.Handler.UserServiceHandler)
	})

	a.Server.SetHttpRegister(func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
		return server2.HttpRegisterFunc(ctx, mux, endpoint, opts,
			[]server2.HttpRegister{
				pb.RegisterUserServiceHandlerFromEndpoint,
			}...,
		)
	})

	return a.Server.Run()
}