package config

import "amethis-backend/pkg/ai"

type BaseModelConfig struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Model       string `json:"model"`
}

type OllamaModelConfig struct {
	BaseModelConfig
	BaseURL      string         `json:"base_url"`
	SystemPrompt string         `json:"system_prompt"`
	Options      map[string]any `json:"options"`
}

type OpenAIModelConfig struct {
	BaseModelConfig
	APIKey    string         `json:"api_key"`
	BaseURL   string         `json:"base_url"`
	ModelType string         `json:"model_type"`
	Options   map[string]any `json:"options"`
}

type AnthropicModelConfig struct {
	BaseModelConfig
	APIKey    string         `json:"api_key"`
	BaseURL   string         `json:"base_url"`
	ModelType string         `json:"model_type"`
	Options   map[string]any `json:"options"`
}

type ModelConfigs struct {
	DefaultProvider string                          `json:"default_provider"`
	Ollama          map[string]OllamaModelConfig    `json:"ollama_models"`
	OpenAI          map[string]OpenAIModelConfig    `json:"openai_models"`
	Anthropic       map[string]AnthropicModelConfig `json:"anthropic_models"`
}

func NewModelConfigs() ModelConfigs {
	return ModelConfigs{
		DefaultProvider: "ollama",

		Ollama: map[string]OllamaModelConfig{
			ai.MODEL_QWEN2_7B: {
				BaseModelConfig: BaseModelConfig{
					Name:        "Qwen2 7B",
					Description: "Qwen2 7B base model",
					Model:       "qwen2:7b",
				},
				BaseURL: "http://ollama:11434",
				SystemPrompt: `You are an AI assistant that provides concise and accurate responses to user questions.
				Keep your answers brief unless explicitly asked for more detail. You must not provide answers or engage
				with any unethical, harmful, or illegal topics, including but not limited to violence, exploitation, or pedophilia.`,
				Options: map[string]any{
					"repeat_penalty": 1.1,
				},
			},
			ai.MODEL_LLAMA2_7B: {
				BaseModelConfig: BaseModelConfig{
					Name:        "LLaMA2 7B",
					Description: "Meta's LLaMA2 7B base model",
					Model:       "llama2:7b",
				},
				BaseURL:      "http://ollama:11434",
				SystemPrompt: "You are a helpful assistant!",
			},
		},

		OpenAI: map[string]OpenAIModelConfig{
			ai.MODEL_GPT4_0: {
				BaseModelConfig: BaseModelConfig{
					Name:        "GPT-4",
					Description: "Most capable GPT-4 model",
					Model:       "gpt-4-turbo-preview",
				},
				APIKey:    "${OPENAI_API_KEY}",
				BaseURL:   "https://api.openai.com/v1",
				ModelType: "gpt-4-turbo-preview",
			},
		},

		Anthropic: map[string]AnthropicModelConfig{
			ai.MODEL_CLAUDE_3_5_SONNET: {
				BaseModelConfig: BaseModelConfig{
					Name:        "Claude 3 Opus",
					Description: "Most capable Claude model",
					Model:       "claude-3-opus-20240229",
				},
				APIKey:    "${ANTHROPIC_API_KEY}",
				BaseURL:   "https://api.anthropic.com/v1",
				ModelType: "claude-3-sonnet-20240229",
			},
		},
	}
}
