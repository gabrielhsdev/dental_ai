package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type AuditLogs struct {
	Id               string          `json:"id"`
	RequestId        string          `json:"requestId"`
	RequestIp        string          `json:"requestIp"`
	RequestTimestamp time.Time       `json:"requestTimestamp"`
	UserId           uuid.UUID       `json:"userId,omitempty"`
	Action           string          `json:"action"`
	Resource         string          `json:"resource"`
	Extra            json.RawMessage `json:"extra"`
	CreatedAt        time.Time       `json:"createdAt,omitempty"`
}
