// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package container

import (
	"github.com/go-mumu/cs-go/api/container/provider"
	"github.com/go-mumu/cs-go/api/controller"
	"github.com/go-mumu/cs-go/library/redis"
	redis2 "github.com/redis/go-redis/v9"
)

// Injectors from injector.go:

func InitApp() (*App, func(), error) {
	client := redis.Init()
	userController := controller.NewUserController()
	providerController := provider.NewController(userController)
	app := &App{
		RedisClient: client,
		Controller:  providerController,
	}
	return app, func() {
	}, nil
}

// injector.go:

type App struct {
	RedisClient *redis2.Client
	Controller  *provider.Controller
}
