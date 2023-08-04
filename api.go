/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-07-28
 * File: api.go
 * Desc: main api
 */

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-mumu/cs-go/api/container"
)
import "github.com/go-mumu/cs-go/api/router"

func main() {
	app, cleanfunc, err := container.InitApp()
	defer cleanfunc()

	r := gin.Default()

	router.Router(r, app)

	if err = r.Run(":8888"); err != nil {
		return
	}
}
