/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-07-26
 * File: dao.go
 * Desc:
 */

package provider

import (
	"github.com/go-mumu/cs-go/dal/dao"
	"github.com/google/wire"
)

var DaoProviderSet = wire.NewSet(
	wire.Struct(new(Dao), "*"),
	dao.NewWxuserDao,
)

type Dao struct {
	WxuserDao *dao.WxuserDao
}
