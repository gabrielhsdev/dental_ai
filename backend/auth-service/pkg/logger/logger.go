package logger

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/internal/models"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/internal/service"
	"go.uber.org/zap"
)

type LoggerInterface interface {
	Error(ctx context.Context, action string, err error, resource string, extra map[string]interface{})
	Info(ctx context.Context, action string, resource string, extra map[string]interface{})
	buildAuditLog(ctx context.Context, action string, resource string, message string, extra map[string]interface{}) models.AuditLogs
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

func (logger *Logger) Error(ctx context.Context, action string, err error, resource string, extra map[string]interface{}) {
	logger.zap.Error(action, zap.Error(err))

	auditLog := logger.buildAuditLog(ctx, action, resource, err.Error(), extra)
	res, err := logger.auditLogsService.CreateAuditLogs(&auditLog)
	if err != nil {
		logger.zap.Error("Failed to create audit log", zap.Error(err))
		return
	}
	logger.zap.Info("Audit log created", zap.String("id", res.Id), zap.String("action", action), zap.String("resource", resource))
}

func (logger *Logger) Info(ctx context.Context, action string, resource string, extra map[string]interface{}) {
	logger.zap.Info(action)
	auditLog := logger.buildAuditLog(ctx, action, resource, "", extra)
	res, err := logger.auditLogsService.CreateAuditLogs(&auditLog)
	if err != nil {
		logger.zap.Error("Failed to create audit log", zap.Error(err))
		return
	}
	logger.zap.Info("Audit log created", zap.String("id", res.Id), zap.String("action", action), zap.String("resource", resource))
}

func (logger *Logger) buildAuditLog(ctx context.Context, action string, resource string, message string, extra map[string]interface{}) models.AuditLogs {
	headers := extractHeaders(ctx)
	now := time.Now().UTC().Format(time.RFC3339)
	extraBytes, _ := json.Marshal(map[string]interface{}{
		"message": message,
		"extra":   extra,
	})

	return models.AuditLogs{
		RequestId:        headers["X-Request-Id"],
		RequestIp:        headers["X-Real-Ip"],
		RequestTimestamp: headers["X-Current-Timestamp"],
		UserId:           headers["User-Id"], // optional: add a custom middleware to set this
		Action:           action,
		Resource:         resource,
		Extra:            json.RawMessage(extraBytes),
		CreatedAt:        now,
	}
}

func extractHeaders(ctx context.Context) map[string]string {
	req, ok := ctx.Value("httpRequest").(*http.Request)
	if !ok {
		return map[string]string{}
	}

	return map[string]string{
		"X-Request-Id":        req.Header.Get("X-Request-Id"),
		"X-Real-Ip":           req.Header.Get("X-Real-IP"),
		"X-Current-Timestamp": req.Header.Get("X-Current-Timestamp"),
		"User-Id":             req.Header.Get("X-User-Id"), // optional if you're adding that later
	}
}
