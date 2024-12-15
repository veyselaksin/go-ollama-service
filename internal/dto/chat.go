package dto

import "amethis-backend/pkg/ai/ollama"

type ChatRequest struct {
	Model    string           `json:"model"`
	Messages []ollama.Message `json:"messages"`
}

type ChatResponse struct {
	Content string `json:"content"`
	Done    bool   `json:"done"`
}
