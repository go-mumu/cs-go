// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package container

import (
	"github.com/go-mumu/cs-go/library/config"
	"github.com/go-mumu/cs-go/library/redis"
	redis2 "github.com/redis/go-redis/v9"
)

// Injectors from injector.go:

func InitApp() (*App, func(), error) {
	configConfig := config.Init()
	client := redis.InitRedis(configConfig)
	app := &App{
		Config:      configConfig,
		RedisClient: client,
	}
	return app, func() {
	}, nil
}

// injector.go:

type App struct {
	Config      *config.Config
	RedisClient *redis2.Client
}
