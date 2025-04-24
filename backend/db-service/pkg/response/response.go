package response

import (
	"time"

	"github.com/gin-gonic/gin"
)

type ResponseStruct struct {
	Status    int         `json:"status"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Error     string      `json:"error,omitempty"`
	Timestamp int64       `json:"timestamp"`
}

type ResponseManagerInterface interface {
	Send(context *gin.Context, status int, mesage string, data interface{}, err error)
}

type ResponseManager struct{}

func NewResponseManager() ResponseManagerInterface {
	return &ResponseManager{}
}

func (responseManager *ResponseManager) Send(context *gin.Context, status int, mesage string, data interface{}, err error) {
	resp := ResponseStruct{
		Status:    status,
		Message:   mesage,
		Timestamp: time.Now().Unix(),
	}

	if err != nil {
		resp.Error = err.Error()
	} else {
		resp.Data = data
	}

	context.JSON(status, resp)
}
