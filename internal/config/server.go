package config

import (
	"errors"
	"time"

	"amethis-backend/pkg/cresponse"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type ServerConfig struct {
	Host         string
	Port         string
	BodyLimit    int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func NewServerConfig() ServerConfig {
	return ServerConfig{
		Host:         getEnvOrDefault("APP_HOST", "0.0.0.0"),
		Port:         getEnvOrDefault("APP_PORT", "8000"),
		BodyLimit:    50 * 1024 * 1024, // 50MB
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
}

func NewFiberConfig(serverConfig ServerConfig) fiber.Config {
	return fiber.Config{
		AppName:      "Amethis",
		BodyLimit:    serverConfig.BodyLimit,
		ReadTimeout:  serverConfig.ReadTimeout,
		WriteTimeout: serverConfig.WriteTimeout,

		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			var code int = fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a fiber.*Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			log.Error("Error occurred: ", err)

			return cresponse.ErrorResponse(ctx, code, "Unexpected error occurred")
		},
	}
}

func GetLanguage(ctx *fiber.Ctx) string {
	return ctx.Get("Accept-Language", "en")
}
