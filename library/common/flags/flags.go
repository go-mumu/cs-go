/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-07-25
 * File: flags.go
 * Desc: 命令行解析
 */

package flags

import (
	"flag"
	jsoniter "github.com/json-iterator/go"
)

var ConsulMap map[string]string
var LogPath string

func init() {
	consul := flag.String("consul", `{"addr": "http://127.0.0.1:8500", "config_path": "config/local"}`, "")
	_ = jsoniter.UnmarshalFromString(*consul, &ConsulMap)

	flag.StringVar(&LogPath, "l", "./log.log", "-l set log path")

	flag.Parse()
}
