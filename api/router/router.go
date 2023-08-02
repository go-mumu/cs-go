/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-07-27
 * File: router.go
 * Desc: 路由
 */

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-mumu/cs-go/api/controller"
	"github.com/go-mumu/cs-go/api/middleware"
)

func Router(router *gin.Engine) {
	api := router.Group("/api", middleware.Login())
	{
		api.POST("isVip", controller.NewUserController().IsVip)
	}
}
