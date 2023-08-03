/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-07-26
 * File: wxuser_dao.go
 * Desc: 用户表
 */

package dao

import (
	"context"
	"github.com/go-mumu/cs-go/library/mysql"
	"github.com/go-mumu/cs-go/service/dal/model"
)

type WxuserDao struct {
	DB *mysql.DefMysql
}

func NewWxuserDao(db *mysql.DefMysql) *WxuserDao {
	return &WxuserDao{DB: db}
}

// GetUserInfoByMid 根据 mid 获取用户信息
func (dao *WxuserDao) GetUserInfoByMid(ctx context.Context, mid int64) *model.Wxuser {
	user := &model.Wxuser{}

	dao.DB.WithContext(ctx).Model(&model.Wxuser{}).Where("mid = ?", mid).Find(&user)

	return user
}
