package api

import (
	"amethis-backend/cmd/api/handler/v1"
	"amethis-backend/internal/service"
	"amethis-backend/pkg/ollama"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// HealthCheck godoc
// @Summary Health Check API
// @Description Health Check for the API
// @Tags Health Check
// @Accept application/json
// @Produce application/json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func health(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "ok",
	})
}

func InitializeRouters(app *fiber.App, connection *gorm.DB, redis *redis.Client) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	completionHandler := handler.NewCompletionHandler(service.NewChatService(ollama.NewOllamaService(ollama.DefaultConfig)))

	v1.Post("/completion", completionHandler.Completion)

}
