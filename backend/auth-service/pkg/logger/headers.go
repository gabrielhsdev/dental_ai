package logger

import (
	"net"
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

func (logger *Logger) extractHeaders(context *gin.Context) HeadersInterface {
	parsedTime := extractTime(context)
	xRealIpParsed := extractIp(context)
	userId := extractUserId(context)

	headers := HeadersInterface{
		XRequestId:        context.GetHeader("X-Request-Id"),
		XRealIp:           xRealIpParsed.String(),
		XCurrentTimestamp: parsedTime,
		Authorization:     context.GetHeader("Authorization"),
		UserId:            userId,
	}

	return headers
}

func extractTime(context *gin.Context) time.Time {
	parsedTime := time.Now()
	if timestamp := context.GetHeader("X-Current-Timestamp"); timestamp != "" {
		if t, err := time.Parse(time.RFC3339Nano, timestamp); err == nil {
			parsedTime = t
		} else {
			parsedTime = time.Now()
		}
	}
	return parsedTime
}

func extractIp(context *gin.Context) net.IP {
	xRealIp := context.GetHeader("X-Real-IP")
	xRealIpParsed := net.ParseIP(xRealIp)
	if xRealIpParsed != nil {
		xRealIpParsed = xRealIpParsed.To4()
	}
	if xRealIpParsed == nil {
		xRealIpParsed = net.IPv4(0, 0, 0, 0)
	}
	return xRealIpParsed
}

func extractUserId(context *gin.Context) uuid.UUID {
	userIdParsed := uuid.Nil
	if userId := context.GetHeader("User-Id"); userId != "" {
		if id, err := uuid.Parse(userId); err == nil {
			userIdParsed = id
		}
	}
	return userIdParsed
}
