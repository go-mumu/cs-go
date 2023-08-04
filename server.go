/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-07-25
 * File: main.go
 * Desc: main
 */

package main

import (
	"github.com/go-mumu/cs-go/library/log"
	"github.com/go-mumu/cs-go/service/container"
	"os"
)

func main() {
	// init app
	app, cleanfunc, err := container.InitApp()
	defer cleanfunc()

	if err != nil {
		log.Cli.Info("init app failed")
		os.Exit(1)
	}

	err = app.Run()
	if err != nil {
		return
	}
}
