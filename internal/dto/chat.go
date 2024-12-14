package dto

import "amethis-backend/pkg/ollama"

type ChatRequest struct {
	Model    string           `json:"model"`
	Messages []ollama.Message `json:"messages"`
	Stream   bool             `json:"stream"`
}

type ChatResponse struct {
	Content string `json:"content"`
	Done    bool   `json:"done"`
}
