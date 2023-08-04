/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-07-28
 * File: client.go
 * Desc: service client
 */

package client

import (
	"github.com/go-mumu/cs-go/library/config"
	"github.com/go-mumu/cs-go/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Config *config.Config
}

func NewServiceClient(conf *config.Config) *ServiceClient {
	return &ServiceClient{
		Config: conf,
	}
}

func (c *ServiceClient) UserClient() pb.UserServiceClient {
	return pb.NewUserServiceClient(c.dial())
}

func (c *ServiceClient) dial() *grpc.ClientConn {
	conn, err := grpc.Dial(
		c.Config.Client.ServiceAddr+c.Config.Rpc.GrpcAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil
	}

	return conn
}
