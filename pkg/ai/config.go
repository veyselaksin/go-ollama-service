package ai

// Role represents the role of a participant in a conversation
type Role string

const (
	System    Role = "system"
	User      Role = "user"
	Assistant Role = "assistant"
)

type ModelProvider string

const (
	Ollama    ModelProvider = "ollama"
	OpenAI    ModelProvider = "openai"
	Anthropic ModelProvider = "anthropic"
)

const MODEL_PROVIDER_DEFAULT = Ollama

const (
	MODEL_QWEN2_7B          = "MODEL_QWEN2_7B"
	MODEL_LLAMA2_7B         = "MODEL_LLAMA2_7B"
	MODEL_GPT4_0            = "MODEL_GPT4_0"
	MODEL_CLAUDE_3_5_SONNET = "MODEL_CLAUDE_3_5_SONNET"
)
