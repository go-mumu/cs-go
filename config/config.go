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
	Rpc      *RpcConf   `toml:"rpc" mapstructure:"rpc"`
	DefRedis *RedisConf `toml:"redis" mapstructure:"redis"`
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

// RedisConf redis config
type RedisConf struct {
	Host           string `toml:"host" mapstructure:"host"`
	Password       string `toml:"password" mapstructure:"password"`
	Database       int    `toml:"database" mapstructure:"database"`
	PrefixKey      string `toml:"prefix_key" mapstructure:"prefix_key"`
	ReadTimeoutMs  int    `toml:"read_timeout_ms" mapstructure:"read_timeout_ms"`
	WriteTimeoutMs int    `toml:"write_timeout_ms" mapstructure:"write_timeout_ms"`
}

// RpcConf 配置
type RpcConf struct {
	GrpcAddr           string `toml:"grpc_addr" mapstructure:"grpc_addr"`                       // grpc server地址
	HttpAddr           string `toml:"http_addr" mapstructure:"http_addr"`                       // http server地址
	GrpcHandlerTimeout int    `toml:"grpc_handler_timeout" mapstructure:"grpc_handler_timeout"` // Grpc Handler timeout(ms), default 5000
	HttpReadTimeout    int    `toml:"http_read_timeout" mapstructure:"http_read_timeout"`       // Receive http request timeout(ms), including the body, default 5000
	HttpWriteTimeout   int    `toml:"http_write_timeout" mapstructure:"http_write_timeout"`     // Keep-alive timeout(ms), default 60000
	HttpIdleTimeout    int    `toml:"http_idle_timeout" mapstructure:"http_idle_timeout"`       // Keep-alive timeout(ms), default 60000
	GrpcIdleTimeout    int    `toml:"grpc_idle_timeout" mapstructure:"grpc_idle_timeout"`       // grpc Keep-alive timeout(ms), default 60000
	MaxBodySize        int    `toml:"max_body_size" mapstructure:"max_body_size"`               // grpc 能处理的最大bodysize 20M
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
