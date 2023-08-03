/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-08-03
 * File: controller.go
 * Desc: controller provider
 */

package provider

import (
	"github.com/go-mumu/cs-go/api/controller"
	"github.com/google/wire"
)

var ControllerProviderSet = wire.NewSet(
	NewController,
	controller.NewUserController,
)

type Controller struct {
	UserController *controller.UserController
}

func NewController(
	UserController *controller.UserController,
) *Controller {
	return &Controller{
		UserController: UserController,
	}
}
