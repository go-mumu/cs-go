//go:build wireinject

/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-08-02
 * File: injector.go
 * Desc:
 */

package container

import (
	"github.com/go-mumu/cs-go/api/container/provider"
	"github.com/go-mumu/cs-go/library/config"
	"github.com/go-mumu/cs-go/library/redis"
	"github.com/google/wire"
)

func InitApp() (*provider.App, func(), error) {
	panic(
		wire.Build(
			wire.Struct(new(provider.App), "*"),
			redis.InitRedis,
			config.Init,
		),
	)
}
