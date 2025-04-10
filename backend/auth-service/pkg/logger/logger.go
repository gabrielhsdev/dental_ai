package logger

import (
	"encoding/json"
	"net"
	"time"

	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/internal/models"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type HeadersInterface struct {
	XRequestId        string    `json:"X-Request-Id"`
	XRealIp           string    `json:"X-Real-IP"`
	XCurrentTimestamp time.Time `json:"X-Current-Timestamp"`
	Authorization     string    `json:"Authorization"`
	UserId            uuid.UUID `json:"User-Id"`
}

type LoggerInterface interface {
	Error(ctx *gin.Context, action string, err error, resource string, extra map[string]interface{})
	Info(ctx *gin.Context, action string, resource string, extra map[string]interface{})
	buildAuditLog(ctx *gin.Context, action string, resource string, message string, extra map[string]interface{}) models.AuditLogs
	extractHeaders(c *gin.Context) HeadersInterface
}

type Logger struct {
	zap              *zap.Logger
	auditLogsService service.AuditLogsServiceInterface
}

func NewLogger(zapLogger *zap.Logger, auditLogsService service.AuditLogsServiceInterface) LoggerInterface {
	return &Logger{
		zap:              zapLogger,
		auditLogsService: auditLogsService,
	}
}

func (logger *Logger) Error(ctx *gin.Context, action string, err error, resource string, extra map[string]interface{}) {
	logger.zap.Error(action, zap.Error(err))
	auditLog := logger.buildAuditLog(ctx, action, resource, err.Error(), extra)
	res, err := logger.auditLogsService.CreateAuditLogs(&auditLog)
	if err != nil {
		logger.zap.Error("Failed to create audit log", zap.Error(err))
		return
	}
	logger.zap.Info("Audit log created", zap.String("id", res.Id), zap.String("action", action), zap.String("resource", resource))
}

func (logger *Logger) Info(ctx *gin.Context, action string, resource string, extra map[string]interface{}) {
	logger.zap.Info(action)
	auditLog := logger.buildAuditLog(ctx, action, resource, "", extra)
	res, err := logger.auditLogsService.CreateAuditLogs(&auditLog)
	if err != nil {
		logger.zap.Error("Failed to create audit log", zap.Error(err))
		return
	}
	logger.zap.Info("Audit log created", zap.String("id", res.Id), zap.String("action", action), zap.String("resource", resource))
}

func (logger *Logger) buildAuditLog(ctx *gin.Context, action string, resource string, message string, extra map[string]any) models.AuditLogs {
	headers := logger.extractHeaders(ctx)

	now := time.Now().UTC().Format(time.RFC3339)
	extraBytes, _ := json.Marshal(map[string]any{
		"message": message,
		"extra":   extra,
	})

	return models.AuditLogs{
		RequestId:        headers.XRequestId,
		RequestIp:        net.ParseIP(headers.XRealIp),
		RequestTimestamp: headers.XCurrentTimestamp,
		UserId:           headers.UserId,
		Action:           action,
		Resource:         resource,
		Extra:            json.RawMessage(extraBytes),
		CreatedAt:        now,
	}
}

func (logger *Logger) extractHeaders(c *gin.Context) HeadersInterface {
	parsedTime := time.Now()
	if timestamp := c.GetHeader("X-Current-Timestamp"); timestamp != "" {
		if t, err := time.Parse(time.RFC3339, timestamp); err == nil {
			parsedTime = t
		} else {
			parsedTime = time.Now()
		}
	}

	// TODO: Stopped here, split the logic into more functions, maybe a helper file for the headers stuff
	headers := HeadersInterface{
		XRequestId:        c.GetHeader("X-Request-Id"),
		XRealIp:           c.GetHeader("X-Real-IP"),
		XCurrentTimestamp: parsedTime,
		Authorization:     c.GetHeader("Authorization"),
		UserId:            uuid.Nil,
	}

	// Set UserId to be a valid UUID
	if userId := c.GetHeader("User-Id"); userId != "" {
		if id, err := uuid.Parse(userId); err == nil {
			headers.UserId = id
		}
	}

	return headers
}
