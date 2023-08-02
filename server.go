/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-07-25
 * File: main.go
 * Desc: main
 */

package main

import (
	"github.com/go-mumu/cs-go/container"
	"github.com/go-mumu/cs-go/log"
	"os"
)

func main() {
	// init app
	app, cleanfunc, err := container.InitApp()
	if err != nil {
		log.Cli.Info("init app failed")
		os.Exit(1)
	}

	defer cleanfunc()

	// gen default mysql models
	// dal.GenDefModels(app.DefMysql.DB)

	err = app.Run()
	if err != nil {
		return
	}
}
