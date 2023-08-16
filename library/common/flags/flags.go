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

var (
	ConsulAddr       string
	ConsulConfigPath string
)

var (
	GRPCServiceIp   string
	GRPCServicePort int
	HTTPServiceIp   string
	HTTPServicePort int
)

func init() {
	// log
	flag.StringVar(&LogPath, "log-path", "", "-log-path")

	// consul
	flag.StringVar(&ConsulAddr, "consul-addr", "", "-consul-addr")
	flag.StringVar(&ConsulConfigPath, "consul-config-path", "", "-consul-config-path")

	// service
	flag.StringVar(&GRPCServiceIp, "grpc-service-ip", "", "-grpc-service-ip")
	flag.IntVar(&GRPCServicePort, "grpc-service-port", 0, "-grpc-service-port")
	flag.StringVar(&HTTPServiceIp, "http-service-ip", "", "-http-service-ip")
	flag.IntVar(&HTTPServicePort, "http-service-port", 0, "-http-service-port")

	flag.Parse()
}
