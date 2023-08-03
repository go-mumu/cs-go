/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-07-26
 * File: register.go
 * Desc:
 */

package server

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type GrpcRegister func(*grpc.Server)

type HttpRegister func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error

func HttpRegisterFunc(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption, registerFunc ...HttpRegister) error {
	for _, f := range registerFunc {
		err := f(ctx, mux, endpoint, opts)
		if err != nil {
			return err
		}
	}

	return nil
}
