//go:build wireinject

/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-07-26
 * File: injector.go
 * Desc: 注入器
 */

package inject

import (
	"github.com/go-mumu/cs-go/config"
	"github.com/go-mumu/cs-go/dal"
	"github.com/go-mumu/cs-go/handler"
	"github.com/go-mumu/cs-go/mysql"
	"github.com/go-mumu/cs-go/server"
	"github.com/google/wire"
)

// InitApp 注入全局应用程序
func InitApp() (*App, func(), error) {
	panic(
		wire.Build(
			wire.Struct(new(App), "*"),
			mysql.InitDef,
			config.Init,
			dal.DaoProviderSet,
			handler.HandlersProviderSet,
			server.NewServer,
		),
	)
}
