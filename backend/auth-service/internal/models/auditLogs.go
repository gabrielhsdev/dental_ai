package models

import "encoding/json"

type AuditLogs struct {
	Id               string          `json:"id"`
	RequestId        string          `json:"requestId"`
	RequestIp        string          `json:"requestIp"`
	RequestTimestamp string          `json:"requestTimestamp"`
	UserId           string          `json:"userId,omitempty"`
	Action           string          `json:"action"`
	Resource         string          `json:"resource"`
	Extra            json.RawMessage `json:"extra"`
	CreatedAt        string          `json:"createdAt"`
}
