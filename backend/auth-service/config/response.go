package config

import (
	"time"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Status    int         `json:"status"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Error     string      `json:"error,omitempty"`
	Timestamp int64       `json:"timestamp"`
}

// SendResponse generates a standardized JSON response.
func SendResponse(context *gin.Context, status int, message string, data interface{}, err error) {
	resp := APIResponse{
		Status:    status,
		Message:   message,
		Timestamp: time.Now().Unix(),
	}

	if err != nil {
		resp.Error = err.Error()
	} else {
		resp.Data = data
	}

	context.JSON(status, resp)
}
