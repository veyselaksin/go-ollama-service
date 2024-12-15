package ollama

import (
	"amethis-backend/pkg/ai"
)

// RequestBody represents the structure of requests to Ollama API
type RequestBody struct {
	Model    string         `json:"model"`    // Model name
	Messages []Message      `json:"messages"` // Chat messages
	Stream   bool           `json:"stream"`   // Enable streaming responses
	Options  map[string]any `json:"options"`  // Model options
}

// ResponseBody represents the structure of responses from Ollama API
type ResponseBody struct {
	Model     string `json:"model"`      // Model name
	CreatedAt string `json:"created_at"` // Response timestamp
	Message   struct {
		Role    string `json:"role"`    // Message role (system/user/assistant)
		Content string `json:"content"` // Message content
	} `json:"message"`
	Done               bool   `json:"done"` // Response completion flag
	DoneReason         string `json:"done_reason,omitempty"`
	TotalDuration      int64  `json:"total_duration,omitempty"`
	LoadDuration       int64  `json:"load_duration,omitempty"`
	PromptEvalCount    int    `json:"prompt_eval_count,omitempty"`
	PromptEvalDuration int64  `json:"prompt_eval_duration,omitempty"`
	EvalCount          int    `json:"eval_count,omitempty"`
	EvalDuration       int64  `json:"eval_duration,omitempty"`
}

// Message represents a chat message
type Message struct {
	Role    ai.Role `json:"role"`    // Message role (system/user/assistant)
	Content string  `json:"content"` // Message content
}
