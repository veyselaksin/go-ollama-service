package config

import "github.com/gofiber/fiber/v2"

type Config struct {
	App    AppConfig
	Server ServerConfig
	Model  ModelConfigs
	Fiber  fiber.Config
}

func NewConfig() *Config {
	return &Config{
		App:    NewAppConfig(),
		Server: NewServerConfig(),
		Model:  NewModelConfigs(),
		Fiber:  NewFiberConfig(NewServerConfig()),
	}
}
