// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package container

import (
	"github.com/go-mumu/cs-go/config"
	"github.com/go-mumu/cs-go/container/provider"
	"github.com/go-mumu/cs-go/dal/dao"
	"github.com/go-mumu/cs-go/mysql"
	"github.com/go-mumu/cs-go/server"
)

// Injectors from injector.go:

// InitApp 注入全局应用程序
func InitApp() (*App, func(), error) {
	configConfig := config.Init()
	defMysql := mysql.InitDef(configConfig)
	wxuserDao := dao.NewWxuserDao(defMysql)
	providerDao := provider.NewDao(wxuserDao)
	serverServer := server.NewServer()
	app := &App{
		DefMysql: defMysql,
		Config:   configConfig,
		Dao:      providerDao,
		Server:   serverServer,
	}
	return app, func() {
	}, nil
}
