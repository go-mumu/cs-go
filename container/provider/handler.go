/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-08-03
 * File: handler.go
 * Desc:
 */

package provider

import (
	"github.com/go-mumu/cs-go/handler"
	"github.com/google/wire"
)

var HandlerProviderSet = wire.NewSet(
	wire.NewSet(NewHandler),
	wire.NewSet(handler.NewUserServiceHandler),
)

type Handler struct {
	UserServiceHandler *handler.UserServiceHandler
}

func NewHandler(
	userSvcH *handler.UserServiceHandler,
) *Handler {
	return &Handler{
		UserServiceHandler: userSvcH,
	}
}