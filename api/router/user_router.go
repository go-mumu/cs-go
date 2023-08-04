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
	"github.com/go-mumu/cs-go/api/container"
	"github.com/go-mumu/cs-go/api/middleware"
)

func UserRouter(api *gin.RouterGroup, app *container.App) {

	user := api.Group("/user", middleware.AuthToken(app.RedisClient))
	{
		user.POST("isVip", app.Controller.UserController.IsVip)
	}
}
