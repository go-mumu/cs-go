/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-08-02
 * File: app.go
 * Desc:
 */

package provider

import (
	"github.com/go-mumu/cs-go/library/config"
	"github.com/redis/go-redis/v9"
)

type App struct {
	Config      *config.Config
	RedisClient *redis.Client
}
