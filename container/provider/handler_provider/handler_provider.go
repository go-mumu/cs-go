/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-07-27
 * File: handler_provider.go
 * Desc: handler progider
 */

package handler_provider

import (
	"github.com/go-mumu/cs-go/container/provider/dao_provider"
	"github.com/go-mumu/cs-go/handler"
)

type Handlers struct {
	UserServiceHandler *handler.UserServiceHandler
}

func NewHandlers(dao *dao_provider.Dao) *Handlers {
	return &Handlers{
		UserServiceHandler: handler.NewUserServiceHandler(dao),
	}
}
