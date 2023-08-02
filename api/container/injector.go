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
	"github.com/go-mumu/cs-go/config"
	localRedis "github.com/go-mumu/cs-go/redis"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

func Redis() *redis.Client {
	wire.Build(
		wire.NewSet(config.Init),
		wire.NewSet(localRedis.InitRedis),
	)

	return new(redis.Client)
}
