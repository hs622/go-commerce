package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Data    interface{}       `json:"data,omitempty"`
	Error   map[string]string `json:"error,omitempty"`
}

type APIError struct {
	StatusCode int `json:"statucCode"`
	Message    any `json:"errors"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("api error: %d", e.StatusCode)
}

func NewAPIError(statusCode int, err error) APIError {
	return APIError{
		StatusCode: statusCode,
		Message:    err.Error(),
	}
}

func InvalidRequestData(errors map[string]string) APIError {
	return APIError{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    errors,
	}
}

func InvalidJSON() APIError {
	return APIError{
		http.StatusBadRequest,
		fmt.Errorf("Invalid JSON request data"),
	}
}

type APIFunc func(ctx *gin.Context) error

func SuccessResponse(ctx *gin.Context, statusCode int, message string, data interface{}) {
	ctx.JSON(
		statusCode, Response{
			Success: true,
			Message: message,
			Data:    data,
		})
}

func ErrorResponse(ctx *gin.Context, statusCode int, message string, error map[string]string) {
	ctx.JSON(
		statusCode, Response{
			Success: false,
			Message: message,
			Error:   error,
		})
}
