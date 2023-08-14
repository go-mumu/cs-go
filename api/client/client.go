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
	"github.com/go-mumu/cs-go/library/log"
	"github.com/go-mumu/cs-go/proto/pb"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strconv"
)

type ServiceClient struct {
}

func (c *ServiceClient) UserClient() pb.UserServiceClient {
	return pb.NewUserServiceClient(c.dial())
}

func (c *ServiceClient) dial() *grpc.ClientConn {
	service, err := c.discoverConsul()
	if err != nil {
		return nil
	}

	log.Cli.Info("consul service info", "info", service.Service)

	conn, err := grpc.Dial(
		service.Service.Address+":"+strconv.Itoa(service.Service.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil
	}

	return conn
}

func (c *ServiceClient) discoverConsul() (*api.ServiceEntry, error) {
	consulClient, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return nil, err
	}

	serviceEntry, _, err := consulClient.Health().Service(config.V.GetString("service.service_name"), "", true, &api.QueryOptions{})

	service := &api.ServiceEntry{}
	if len(serviceEntry) > 0 {
		service = serviceEntry[0]
	}

	return service, nil
}
