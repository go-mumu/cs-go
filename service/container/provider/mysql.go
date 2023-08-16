/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-08-16
 * File: mysql.go
 * Desc: mysql provider
 */

package provider

import (
	"github.com/go-mumu/cs-go/library/mysql"
	"github.com/google/wire"
)

var MysqlProviderSet = wire.NewSet(
	mysql.InitDef,
)
