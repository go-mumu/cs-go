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
	"github.com/go-mumu/cs-go/api/container"
	"github.com/go-mumu/cs-go/api/middleware"
)

func Router(router *gin.Engine, app *container.App) {

	api := router.Group("/api", middleware.Trace())

	UserRouter(api, app)
}
