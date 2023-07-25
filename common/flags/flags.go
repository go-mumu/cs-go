/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-07-25
 * File: flags.go
 * Desc:
 */

package flags

import "flag"

var ConfPath string
var LogPath string

func init() {
	flag.StringVar(&ConfPath, "c", "./config/config.toml", "-c set config path")
	flag.StringVar(&LogPath, "l", "./log.log", "-l set log path")
	flag.Parse()
}
