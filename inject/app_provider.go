/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-07-26
 * File: provider.go
 * Desc: provider
 */

package inject

import (
	"github.com/go-mumu/cs-go/config"
	"github.com/go-mumu/cs-go/dal"
	"github.com/go-mumu/cs-go/handler"
	"github.com/go-mumu/cs-go/mysql"
	"github.com/go-mumu/cs-go/proto/pb"
	"github.com/go-mumu/cs-go/server"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// App 全局应用程序
type App struct {
	DefMysql *mysql.DefMysql
	Config   *config.Config
	Dao      *dal.Dao
	Server   *server.Server
	Handlers *handler.Handlers
}

func (a *App) Run(s *server.Server, h *handler.Handlers) error {
	s.SetGrpcAddr(config.C.Rpc.GrpcAddr)
	s.SetHttpAddr(config.C.Rpc.HttpAddr)

	s.SetGrpcHandlerTimeout(config.C.Rpc.GrpcHandlerTimeout)

	s.SetHttpReadTimeout(config.C.Rpc.HttpReadTimeout)
	s.SetHttpWriteTimeout(config.C.Rpc.HttpWriteTimeout)

	s.SetGrpcIdleTimeout(config.C.Rpc.GrpcIdleTimeout)
	s.SetHttpIdleTimeout(config.C.Rpc.HttpIdleTimeout)

	s.SetMaxBodySize(config.C.Rpc.MaxBodySize)

	s.SetGrpcRegister(func(s *grpc.Server) {
		pb.RegisterUserServiceServer(s, h.UserServiceHandler)
	})

	s.SetHttpRegister(func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
		return server.HttpRegisterFunc(ctx, mux, endpoint, opts,
			[]server.HttpRegister{
				pb.RegisterUserServiceHandlerFromEndpoint,
			}...,
		)
	})

	return s.Run()
}
