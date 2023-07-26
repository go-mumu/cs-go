/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-07-25
 * File: config.go
 * Desc: 处理配置文件
 */

package config

import (
	"github.com/go-mumu/cs-go/common/flags"
	"github.com/go-mumu/cs-go/log"
	"github.com/spf13/viper"
)

var (
	C *Config
)

type Config struct {
	DefMysql *MysqlConf `toml:"mysql_def" mapstructure:"mysql_def"`
}

// MysqlConf Mysql配置
type MysqlConf struct {
	Username  string `toml:"username" mapstructure:"username"`
	Password  string `toml:"password" mapstructure:"password"`
	Protocol  string `toml:"protocol" mapstructure:"protocol"`
	Address   string `toml:"address" mapstructure:"address"`
	Port      int    `toml:"port" mapstructure:"port"`
	Dbname    string `toml:"dbname" mapstructure:"dbname"`
	Charset   string `toml:"charset" mapstructure:"charset"`
	ParseTime bool   `toml:"parseTime" mapstructure:"parseTime"`
	Loc       string `toml:"loc" mapstructure:"loc"`
}

// Init 注入初始化配置文件
func Init() *Config {
	// 控制台打印信息
	log.Cli.Info("config path", "path", flags.ConfPath)

	// 设置路径、扩展名
	viper.SetConfigFile(flags.ConfPath)
	viper.SetConfigType("toml")

	// 读取配置
	if err := viper.ReadInConfig(); err != nil {
		log.Cli.Error("read config fail", "error", err)
	}

	// 解析失败
	if err := viper.Unmarshal(&C); err != nil {
		log.Cli.Error("unmarshal fail", "error", err)
	}

	return C
}
