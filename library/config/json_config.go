/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-08-10
 * File: json_config.go
 * Desc:
 */

package config

import (
	"github.com/go-mumu/cs-go/library/common/flags"
	"github.com/go-mumu/cs-go/library/log"
	"github.com/spf13/viper"
)

var (
	C *Config
)

type Config struct {
	DefMysql *MysqlConf    `json:"mysql_def" toml:"mysql_def" mapstructure:"mysql_def"`
	DefRedis *RedisConf    `json:"redis" toml:"redis" mapstructure:"redis"`
	Rpc      *RpcConf      `json:"rpc" toml:"rpc" mapstructure:"rpc"`
	Client   *ClientConf   `json:"client" toml:"client" mapstructure:"client"`
	Domain   *DomainConf   `json:"domain" toml:"domain" mapstructure:"domain"`
	Interest *InterestConf `json:"interest" toml:"interest" mapstructure:"interest"`
}

// MysqlConf Mysql配置
type MysqlConf struct {
	Username  string `json:"username" toml:"username" mapstructure:"username"`
	Password  string `json:"password" toml:"password" mapstructure:"password"`
	Protocol  string `json:"protocol" toml:"protocol" mapstructure:"protocol"`
	Address   string `json:"address" toml:"address" mapstructure:"address"`
	Port      int    `json:"port" toml:"port" mapstructure:"port"`
	Dbname    string `json:"dbname" toml:"dbname" mapstructure:"dbname"`
	Charset   string `json:"charset" toml:"charset" mapstructure:"charset"`
	ParseTime bool   `json:"parseTime" toml:"parseTime" mapstructure:"parseTime"`
	Loc       string `json:"loc" toml:"loc" mapstructure:"loc"`
}

// RedisConf redis config
type RedisConf struct {
	Host           string `json:"host" toml:"host" mapstructure:"host"`
	Password       string `json:"password" toml:"password" mapstructure:"password"`
	Database       int    `json:"database" toml:"database" mapstructure:"database"`
	PrefixKey      string `json:"prefix_key" toml:"prefix_key" mapstructure:"prefix_key"`
	ReadTimeoutMs  int    `json:"read_timeout_ms" toml:"read_timeout_ms" mapstructure:"read_timeout_ms"`
	WriteTimeoutMs int    `json:"write_timeout_ms" toml:"write_timeout_ms" mapstructure:"write_timeout_ms"`
}

// RpcConf 配置
type RpcConf struct {
	GrpcAddr           string `json:"grpc_addr" toml:"grpc_addr" mapstructure:"grpc_addr"`                                  // grpc server地址
	HttpAddr           string `json:"http_addr" toml:"http_addr" mapstructure:"http_addr"`                                  // http server地址
	GrpcHandlerTimeout int    `json:"grpc_handler_timeout" toml:"grpc_handler_timeout" mapstructure:"grpc_handler_timeout"` // Grpc Handler timeout(ms), default 5000
	HttpReadTimeout    int    `json:"http_read_timeout" toml:"http_read_timeout" mapstructure:"http_read_timeout"`          // Receive http request timeout(ms), including the body, default 5000
	HttpWriteTimeout   int    `json:"http_write_timeout" toml:"http_write_timeout" mapstructure:"http_write_timeout"`       // Keep-alive timeout(ms), default 60000
	HttpIdleTimeout    int    `json:"http_idle_timeout" toml:"http_idle_timeout" mapstructure:"http_idle_timeout"`          // Keep-alive timeout(ms), default 60000
	GrpcIdleTimeout    int    `json:"grpc_idle_timeout" toml:"grpc_idle_timeout" mapstructure:"grpc_idle_timeout"`          // grpc Keep-alive timeout(ms), default 60000
	MaxBodySize        int    `json:"max_body_size" toml:"max_body_size" mapstructure:"max_body_size"`                      // grpc 能处理的最大bodysize 20M
}

// ClientConf 配置
type ClientConf struct {
	ServiceAddr string `json:"service_addr" toml:"service_addr" mapstructure:"service_addr"`
}

// DomainConf 配置
type DomainConf struct {
	Center string `json:"center" toml:"center" mapstructure:"center"`
}

// InterestConf 权益配置
type InterestConf struct {
	Code string `json:"code" toml:"code" maostructure:"code"`
}

// Init 注入初始化配置文件
func Init() *Config {
	// 控制台打印信息
	log.Cli.Info("config path", "path", flags.ConfPath)

	// 设置路径、扩展名
	viper.SetConfigFile("./config/local.json")
	viper.SetConfigType("json")

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
