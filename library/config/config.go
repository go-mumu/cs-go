/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-08-10
 * File: yaml_config.go
 * Desc: yaml config
 */

package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/go-mumu/cs-go/library/common/flags"
	"github.com/go-mumu/cs-go/library/log"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

var V *viper.Viper

func init() {
	v := viper.New()

	v.SetConfigType("yaml")

	if err := v.AddRemoteProvider("consul", flags.ConsulAddr, flags.ConsulConfigPath); err != nil {
		log.Cli.Error("add consul remote provider failed", "error", err)
	}

	if err := v.ReadRemoteConfig(); err != nil {
		log.Cli.Error("read consul remote config failed", "error", err)
	}

	v.OnConfigChange(func(e fsnotify.Event) {
		log.Cli.Info("config path is change:", e.Name, e.String())
	})

	if err := v.WatchRemoteConfigOnChannel(); err != nil {
		log.Cli.Error("watch remote config error", "error", err)
	}

	log.Cli.Info("init config success.")

	V = v
}
