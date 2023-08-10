/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-08-01
 * File: redis.go
 * Desc: redis connect
 */

package redis

import (
	"github.com/go-mumu/cs-go/library/config"
	"github.com/go-mumu/cs-go/library/log"
	"github.com/redis/go-redis/v9"
	"time"
)

func Init() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:         config.V.GetString("redis.host"),
		DB:           config.V.GetInt("redis.database"),
		ReadTimeout:  time.Duration(config.V.GetInt("redis.read_timeout_ms")) * time.Millisecond,
		WriteTimeout: time.Duration(config.V.GetInt("redis.write_timeout_ms")) * time.Millisecond,
	})

	log.Cli.Info("init redis success.")

	return client
}
