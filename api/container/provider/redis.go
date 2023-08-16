/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-08-16
 * File: redis.go
 * Desc: redis provider
 */

package provider

import (
	"github.com/go-mumu/cs-go/library/redis"
	"github.com/google/wire"
)

var RedisProviderSet = wire.NewSet(
	redis.Init,
)
