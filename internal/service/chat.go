package service

import (
	"amethis-backend/internal/config"
	"amethis-backend/internal/dto"
	"amethis-backend/pkg/ai/ollama"
	"context"
	"fmt"
)

type ChatService interface {
	StreamChat(ctx context.Context, request dto.ChatRequest) (<-chan ollama.ResponseBody, error)
}

type chatService struct {
	providers     config.ModelConfigs
	ollamaService map[string]ollama.OllamaService
}

func NewChatService(providers config.ModelConfigs) ChatService {
	ollamaService := make(map[string]ollama.OllamaService)
	for model, cfg := range providers.Ollama {
		ollamaService[model] = ollama.NewOllamaService(cfg)
	}

	return &chatService{
		providers:     providers,
		ollamaService: ollamaService,
	}
}

func (c *chatService) StreamChat(ctx context.Context, request dto.ChatRequest) (<-chan ollama.ResponseBody, error) {
	requestBody := ollama.RequestBody{
		Model:    c.providers.Ollama[request.Model].Model,
		Messages: request.Messages,
		Stream:   true,
	}

	var responseChan <-chan ollama.ResponseBody
	var err error

	switch c.providers.DefaultProvider {
	case "ollama":
		responseChan, err = c.ollamaService[request.Model].StreamChat(requestBody)
	case "openai":
	case "anthropic":
	default:
		return nil, fmt.Errorf("unsupported model provider: %s", c.providers.DefaultProvider)
	}

	if err != nil {
		return nil, err
	}

	return responseChan, nil
}
