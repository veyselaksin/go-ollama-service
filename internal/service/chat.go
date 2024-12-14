package service

import (
	"amethis-backend/internal/dto"
	"amethis-backend/pkg/ollama"
	"context"
)

type ChatService interface {
	StreamChat(ctx context.Context, request dto.ChatRequest) (<-chan ollama.ResponseBody, error)
}

type chatService struct {
	ollama ollama.OllamaService
}

func NewChatService(ollama ollama.OllamaService) ChatService {
	return &chatService{
		ollama: ollama,
	}
}

func (c *chatService) StreamChat(ctx context.Context, request dto.ChatRequest) (<-chan ollama.ResponseBody, error) {
	requestBody := ollama.RequestBody{
		Model:    request.Model,
		Messages: request.Messages,
		Stream:   request.Stream,
	}

	responseChan, err := c.ollama.StreamChat(requestBody)
	if err != nil {
		return nil, err
	}

	return responseChan, nil
}
