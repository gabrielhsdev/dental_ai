package headers

import (
	"net"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func extractTime(context *gin.Context) time.Time {
	timestamp := context.GetHeader("X-Current-Timestamp")
	if parsed, err := time.Parse(time.RFC3339Nano, timestamp); err == nil {
		return parsed
	}
	return time.Now()
}

func extractIP(context *gin.Context) string {
	ip := net.ParseIP(context.GetHeader("X-Real-IP"))
	if ip != nil {
		ip = ip.To4()
	}
	if ip == nil {
		ip = net.IPv4(0, 0, 0, 0)
	}
	return ip.String()
}

func extractUserId(context *gin.Context) uuid.UUID {
	id, err := uuid.Parse(context.GetHeader("User-Id"))
	if err != nil {
		return uuid.Nil
	}
	return id
}
