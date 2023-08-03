/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-08-03
 * File: user.go
 * Desc: user router
 */

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-mumu/cs-go/api/controller"
	"github.com/go-mumu/cs-go/api/middleware"
	"github.com/redis/go-redis/v9"
)

func UserRouter(api *gin.RouterGroup, redisClient *redis.Client) {
	user := api.Group("/user", middleware.AuthToken(redisClient))
	{
		user.POST("isVip", controller.NewUserController().IsVip)
	}
}
