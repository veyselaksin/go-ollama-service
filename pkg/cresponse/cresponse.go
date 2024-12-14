package cresponse

import "github.com/gofiber/fiber/v2"

// BaseResponse represents the standard API response structure
type BaseResponse struct {
	Success bool        `json:"success"` // Indicates if the request was successful
	Message string      `json:"message"` // Human readable message about the response
	Data    interface{} `json:"data"`    // Optional payload
}

// Response is a generic response helper that allows explicit success state
func Response(ctx *fiber.Ctx, status int, success bool, data interface{}, msg string) error {
	return ctx.Status(status).JSON(BaseResponse{
		Success: success,
		Message: msg,
		Data:    data,
	})
}

// SuccessResponse sends a success response with optional data and message
func SuccessResponse(ctx *fiber.Ctx, status int, data interface{}, msg ...string) error {
	message := "Success"
	if len(msg) > 0 {
		message = msg[0]
	}

	return Response(ctx, status, true, data, message)
}

// ErrorResponse sends an error response with a required message and optional data
func ErrorResponse(ctx *fiber.Ctx, status int, msg string, data ...interface{}) error {
	var responseData interface{}
	if len(data) > 0 {
		responseData = data[0]
	}

	return Response(ctx, status, false, responseData, msg)
}

// RedirectResponse performs a temporary redirect to the specified URL
func RedirectResponse(ctx *fiber.Ctx, url string) error {
	return ctx.Redirect(url, fiber.StatusTemporaryRedirect)
}
