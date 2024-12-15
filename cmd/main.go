package main

import (
	"amethis-backend/cmd/api"
	"amethis-backend/docs"
	"amethis-backend/internal/config"
	"amethis-backend/internal/db/connection"
	"context"
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// Global variables
var (
	once      sync.Once
	conn      *gorm.DB
	redisConn *redis.Client
	appConfig *config.Config
)

// @title Amethis API
// @version 1.0
// @description This is a config for Amethis API.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /v1
func main() {
	// Initialize configurations and connections
	initializeApp()

	// Setup Fiber app
	app := setupFiberApp()

	// Start server
	go startServer(app)

	// Handle graceful shutdown
	if err := gracefulShutdown(app, 5*time.Second); err != nil {
		log.Error("Graceful shutdown error", err)
	}
}

func initializeApp() {
	once.Do(func() {
		// Initialize database connection
		conn = connection.PostgresSQLConnection(config.NewPostgresConfig())

		// Initialize Redis connection
		redisConn = connection.RedisConnection(config.NewRedisConfig())

		// Initialize app config
		appConfig = config.NewConfig()

		// Configure Swagger
		configureSwagger()
	})
}

func configureSwagger() {
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", appConfig.Server.Host, appConfig.Server.Port)
}

func setupFiberApp() *fiber.App {
	app := fiber.New(appConfig.Fiber)

	// Add middleware
	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(compress.New())
	app.Use(logger.New(logger.Config{
		TimeFormat: "2006-01-02T15:04:05.000Z",
		TimeZone:   "Europe/Istanbul",
	}))

	// Initialize routes
	api.InitializeRouters(app, conn, redisConn)

	return app
}

func startServer(app *fiber.App) {
	if err := app.Listen(":" + appConfig.Server.Port); err != nil {
		panic(err)
	}
}

func gracefulShutdown(app *fiber.App, timeout time.Duration) error {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	sig := <-sigChan
	log.Infof("Signal received: %v", sig)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Close database connection
	if db, err := conn.DB(); err == nil {
		if err := shutdownDatabase(ctx, db); err != nil {
			return err
		}
	}

	// Shutdown Fiber app
	return app.Shutdown()
}

func shutdownDatabase(ctx context.Context, db *sql.DB) error {
	ch := make(chan error, 1)
	go func() {
		ch <- db.Close()
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-ch:
		if err != nil {
			log.Error("Database close error", err)
			return err
		}
		return nil
	}
}
