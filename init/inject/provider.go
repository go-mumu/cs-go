/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-07-26
 * File: provider.go
 * Desc: provider
 */

package inject

import (
	"github.com/go-mumu/cs-go/config"
	"github.com/go-mumu/cs-go/mysql"
)

// App 全局应用程序
type App struct {
	DefMysql *mysql.DefMysql
	Config   *config.Config
}
