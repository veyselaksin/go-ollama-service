package ollama

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type OllamaService interface {
	StreamChat(RequestBody) (<-chan ResponseBody, error)
}

type ollamaService struct {
	ollamaConfig OllamaConfig
}

func NewOllamaService(ollamaConfig OllamaConfig) *ollamaService {
	return &ollamaService{
		ollamaConfig: ollamaConfig,
	}
}

func (o *ollamaService) StreamChat(requestBody RequestBody) (<-chan ResponseBody, error) {
	responseChan := make(chan ResponseBody)

	go func() {
		defer close(responseChan)

		// Add default options if none provided
		if requestBody.Options == nil {
			requestBody.Options = o.ollamaConfig.OPTIONS
		}

		// Add system prompt to the top of messages
		requestBody.Messages = append([]Message{{Role: "system", Content: o.ollamaConfig.SYSTEM_PROMPT}}, requestBody.Messages...)

		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			responseChan <- ResponseBody{Done: false, DoneReason: fmt.Sprintf("Error marshalling request body: %v", err)}
			return
		}

		resp, err := http.Post(o.ollamaConfig.BASE_URL+"/api/chat", "application/json", bytes.NewBuffer(jsonBody))
		if err != nil {
			responseChan <- ResponseBody{Done: false, DoneReason: fmt.Sprintf("Error making request: %v", err)}
			return
		}
		defer resp.Body.Close()

		reader := bufio.NewReader(resp.Body)
		for {
			line, err := reader.ReadBytes('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				responseChan <- ResponseBody{Done: false, DoneReason: fmt.Sprintf("Error reading response: %v", err)}
				break
			}

			// Parse the JSON response
			var response struct {
				Message struct {
					Role    string `json:"role"`
					Content string `json:"content"`
				} `json:"message"`
			}

			if err := json.Unmarshal(line, &response); err != nil {
				continue // Skip invalid JSON
			}

			// Send only the content
			if response.Message.Content != "" {
				responseChan <- ResponseBody{Message: response.Message, Done: false}
			}
		}
	}()

	return responseChan, nil
}
