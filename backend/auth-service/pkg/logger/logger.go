package logger

import (
	"encoding/json"

	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/internal/models"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

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

	extraBytes, _ := json.Marshal(map[string]any{
		"message": message,
		"extra":   extra,
	})

	return models.AuditLogs{
		RequestId:        headers.XRequestId,
		RequestIp:        headers.XRealIp,
		RequestTimestamp: headers.XCurrentTimestamp,
		UserId:           headers.UserId,
		Action:           action,
		Resource:         resource,
		Extra:            json.RawMessage(extraBytes),
	}
}
