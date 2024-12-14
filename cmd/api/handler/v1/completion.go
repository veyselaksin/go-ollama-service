package handler

import (
	"amethis-backend/internal/dto"
	"amethis-backend/internal/service"
	"bufio"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type CompletionHandler struct {
	chatService service.ChatService
}

func NewCompletionHandler(chatService service.ChatService) *CompletionHandler {
	return &CompletionHandler{
		chatService: chatService,
	}
}

func (h *CompletionHandler) Completion(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")
	c.Set("Transfer-Encoding", "chunked")

	var request dto.ChatRequest
	if err := c.BodyParser(&request); err != nil {
		return err
	}

	responseChan, err := h.chatService.StreamChat(c.Context(), request)
	if err != nil {
		return err
	}

	c.Context().SetBodyStreamWriter(func(w *bufio.Writer) {
		for chunk := range responseChan {
			response := dto.ChatResponse{
				Content: chunk.Message.Content,
				Done:    chunk.Done,
			}
			jsonBytes, err := json.Marshal(response)
			if err != nil {
				fmt.Fprintf(w, "data: error encoding json: %v\n\n", err)
				w.Flush()
				return
			}
			fmt.Fprintf(w, "data: %s\n\n", jsonBytes)
			w.Flush()
		}
	})

	return nil
}
