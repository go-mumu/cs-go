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
)

var LogPath string
var ConsulAddr string
var ConsulConfigPath string
var NodeIp string
var NodePort int
var NodeId string
var HttpPort int

func init() {
	flag.StringVar(&LogPath, "log-path", "", "-log-path")

	flag.StringVar(&ConsulAddr, "consul-addr", "", "-consul-addr")
	flag.StringVar(&ConsulConfigPath, "consul-config-path", "", "-consul-config-path")

	flag.StringVar(&NodeIp, "node-ip", "", "-node-ip")
	flag.StringVar(&NodeId, "node-id", "", "-node-id")
	flag.IntVar(&NodePort, "node-port", 0, "-node-port")

	flag.IntVar(&HttpPort, "http-port", 0, "-http-port")

	flag.Parse()
}
