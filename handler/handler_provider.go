/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-07-27
 * File: handler_provider.go
 * Desc: handler progider
 */

package handler

import (
	"github.com/google/wire"
)

var HandlersProviderSet = wire.NewSet(
	wire.NewSet(NewUserServiceHandler),
	wire.NewSet(NewHandlers),
)

type Handlers struct {
	UserServiceHandler *UserServiceHandler
}

func NewHandlers(
	userServiceHandler *UserServiceHandler,
) *Handlers {
	return &Handlers{
		UserServiceHandler: userServiceHandler,
	}
}
