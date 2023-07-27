/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-07-26
 * File: dao.go
 * Desc:
 */

package dal

import (
	"github.com/go-mumu/cs-go/dal/dao"
	"github.com/google/wire"
)

var DaoProviderSet = wire.NewSet(
	wire.NewSet(dao.NewWxuserDao),
	wire.NewSet(NewDao),
)

type Dao struct {
	WxuserDao *dao.WxuserDao
}

func NewDao(
	wxuserDao *dao.WxuserDao,
) *Dao {
	return &Dao{
		WxuserDao: wxuserDao,
	}
}
