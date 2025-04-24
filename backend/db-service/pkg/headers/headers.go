package headers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type HeadersInterface struct {
	XRequestId        string    `json:"X-Request-Id"`
	XRealIp           string    `json:"X-Real-IP"`
	XCurrentTimestamp time.Time `json:"X-Current-Timestamp"`
	Authorization     string    `json:"Authorization"`
	UserId            uuid.UUID `json:"User-Id"`
}

type HeadersManagerInterface interface {
	Extract(context *gin.Context) HeadersInterface
}

type HeadersManager struct{}

func NewHeadersManager() HeadersManagerInterface {
	return &HeadersManager{}
}

func (headersManager *HeadersManager) Extract(context *gin.Context) HeadersInterface {
	return HeadersInterface{
		XRequestId:        context.GetHeader("X-Request-Id"),
		XRealIp:           extractIP(context),
		XCurrentTimestamp: extractTime(context),
		Authorization:     context.GetHeader("Authorization"),
		UserId:            extractUserId(context),
	}
}
