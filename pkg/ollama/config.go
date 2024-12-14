package ollama

type OllamaConfig struct {
	BASE_URL      string
	MODEL         string
	OPTIONS       map[string]any
	SYSTEM_PROMPT string
}

var DefaultConfig = OllamaConfig{
	BASE_URL: "http://ollama:11434",
	MODEL:    "qwen2:7b",
	OPTIONS: map[string]any{
		"temperature": 0.7,
		// "max_tokens":  1000,
	},
	SYSTEM_PROMPT: "You are chat assistant! Give a short and concise answer!",
}

type RequestBody struct {
	Model    string         `json:"model"`
	Messages []Message      `json:"messages"`
	Stream   bool           `json:"stream"`
	Options  map[string]any `json:"options"`
}

type ResponseBody struct {
	Model     string `json:"model"`
	CreatedAt string `json:"created_at"`
	Message   struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"message"`
	Done               bool   `json:"done"`
	DoneReason         string `json:"done_reason,omitempty"`
	TotalDuration      int64  `json:"total_duration,omitempty"`
	LoadDuration       int64  `json:"load_duration,omitempty"`
	PromptEvalCount    int    `json:"prompt_eval_count,omitempty"`
	PromptEvalDuration int64  `json:"prompt_eval_duration,omitempty"`
	EvalCount          int    `json:"eval_count,omitempty"`
	EvalDuration       int64  `json:"eval_duration,omitempty"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Role string

const (
	SYSTEM Role = "system"
	USER   Role = "user"
	AID    Role = "assistant"
)

func (r Role) String() string {
	return string(r)
}
