/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-08-01
 * File: redis.go
 * Desc: redis connect
 */

package redis

import (
	"github.com/go-mumu/cs-go/config"
	"github.com/redis/go-redis/v9"
	"time"
)

func InitRedis(conf *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:         conf.DefRedis.Host,
		DB:           conf.DefRedis.Database,
		ReadTimeout:  time.Duration(conf.DefRedis.ReadTimeoutMs) * time.Millisecond,
		WriteTimeout: time.Duration(conf.DefRedis.WriteTimeoutMs) * time.Millisecond,
	})
}
