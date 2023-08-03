/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-08-03
 * File: handler.go
 * Desc:
 */

package provider

import (
	"github.com/go-mumu/cs-go/service/handler"
	"github.com/google/wire"
)

var HandlerProviderSet = wire.NewSet(
	NewHandler,
	handler.NewUserServiceHandler,
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
