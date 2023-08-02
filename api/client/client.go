/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-07-28
 * File: client.go
 * Desc:
 */

package client

import (
	"context"
	"github.com/go-mumu/cs-go/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Clients struct {
}

func (c *Clients) UserClient(ctx context.Context) pb.UserServiceClient {
	return pb.NewUserServiceClient(c.grpcDial(ctx))
}

func (c *Clients) grpcDial(ctx context.Context) *grpc.ClientConn {
	conn, err := grpc.DialContext(ctx, "127.0.0.1:8992", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil
	}

	return conn
}
