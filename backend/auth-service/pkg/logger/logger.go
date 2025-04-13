package logger

import (
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/internal/service"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/pkg/headers"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LoggerInterface interface {
	Error(ctx *gin.Context, action string, err error, resource string, extra map[string]interface{})
	Info(ctx *gin.Context, action string, resource string, extra map[string]interface{})
}

type Logger struct {
	zap              *zap.Logger
	auditLogsService service.AuditLogsServiceInterface
	headersHandler   headers.HeadersManagerInterface
}

func NewLogger(zapLogger *zap.Logger, auditLogsService service.AuditLogsServiceInterface, headersHandler headers.HeadersManagerInterface) LoggerInterface {
	return &Logger{
		zap:              zapLogger,
		auditLogsService: auditLogsService,
		headersHandler:   headersHandler,
	}
}

func (logger *Logger) Error(ctx *gin.Context, action string, err error, resource string, extra map[string]interface{}) {
	logger.zap.Error(action, zap.Error(err))
	headers := logger.headersHandler.Extract(ctx)
	auditLog := logger.auditLogsService.NewAuditLogFromRequest(headers, action, resource, err.Error(), extra)
	res, err := logger.auditLogsService.CreateAuditLogs(&auditLog)
	if err != nil {
		logger.zap.Error("Failed to create audit log", zap.Error(err))
		return
	}
	logger.zap.Info("Audit log created", zap.String("id", res.Id), zap.String("action", action), zap.String("resource", resource))
}

func (logger *Logger) Info(ctx *gin.Context, action string, resource string, extra map[string]interface{}) {
	logger.zap.Info(action)
	headers := logger.headersHandler.Extract(ctx)
	auditLog := logger.auditLogsService.NewAuditLogFromRequest(headers, action, resource, "", extra)
	res, err := logger.auditLogsService.CreateAuditLogs(&auditLog)
	if err != nil {
		logger.zap.Error("Failed to create audit log", zap.Error(err))
		return
	}
	logger.zap.Info("Audit log created", zap.String("id", res.Id), zap.String("action", action), zap.String("resource", resource))
}
