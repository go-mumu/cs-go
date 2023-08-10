/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-08-10
 * File: yaml_config.go
 * Desc: yaml config
 */

package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/go-mumu/cs-go/library/common/flags"
	"github.com/go-mumu/cs-go/library/log"
	"github.com/spf13/viper"
)

var V *viper.Viper

func init() {
	// 控制台打印信息
	log.Cli.Info("config path", "path", flags.ConfPath)

	v := viper.New()

	// 设置路径、扩展名
	v.SetConfigFile("./config/local.yaml")
	v.SetConfigType("yaml")

	// 读取配置
	if err := v.ReadInConfig(); err != nil {
		log.Cli.Error("read config fail", "error", err)
	}

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config path is change:", e.Name, e.String())
	})

	v.WatchConfig()

	V = v
}
