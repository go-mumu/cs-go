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
	"github.com/go-mumu/cs-go/library/config"
	libRedis "github.com/go-mumu/cs-go/library/redis"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

type App struct {
	Config      *config.Config
	RedisClient *redis.Client
}

func InitApp() (*App, func(), error) {
	panic(
		wire.Build(
			wire.Struct(new(App), "*"),
			libRedis.InitRedis,
			config.Init,
		),
	)
}
