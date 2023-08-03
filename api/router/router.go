/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-07-27
 * File: router.go
 * Desc: global router
 */

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func Router(router *gin.Engine, redisClient *redis.Client) {
	api := router.Group("/api")

	UserRouter(api, redisClient)
}
