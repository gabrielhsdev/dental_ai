package logger

import (
	"github.com/gabrielhsdev/dental_ai/backend/db-service/internal/service"
	"github.com/gabrielhsdev/dental_ai/backend/db-service/pkg/headers"
	"github.com/gabrielhsdev/dental_ai/backend/db-service/pkg/resources"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// TODO: Currently we are typecasting our resources into a string, maybe we should refactor more the code so we dont need to do that

type LoggerInterface interface {
	Error(ctx *gin.Context, action string, err error, resource resources.ResourceType, extra map[string]interface{})
	Info(ctx *gin.Context, action string, resource resources.ResourceType, extra map[string]interface{})
}

type Logger struct {
	zap              *zap.Logger
	auditLogsService service.AuditLogsServiceInterface
	headersHandler   headers.HeadersManagerInterface
	resourceManager  resources.ResourceManagerInterface
}

func NewLogger(
	zapLogger *zap.Logger,
	auditLogsService service.AuditLogsServiceInterface,
	headersHandler headers.HeadersManagerInterface,
	resourceManager resources.ResourceManagerInterface,
) LoggerInterface {
	return &Logger{
		zap:              zapLogger,
		auditLogsService: auditLogsService,
		headersHandler:   headersHandler,
		resourceManager:  resourceManager,
	}
}

func (logger *Logger) Error(ctx *gin.Context, action string, err error, resource resources.ResourceType, extra map[string]interface{}) {
	if !logger.resourceManager.ValidateResource(resource) {
		logger.zap.Error("Invalid resource", zap.String("resource", string(resource)))
		return
	}

	errorMessage := "nil error" // This means that the error is nil
	if err != nil {
		errorMessage = err.Error()
	}

	logger.zap.Error(action, zap.Error(err))
	headers := logger.headersHandler.Extract(ctx)
	auditLog := logger.auditLogsService.NewAuditLogFromRequest(headers, action, string(resource), errorMessage, extra)
	res, err := logger.auditLogsService.CreateAuditLogs(&auditLog)
	if err != nil {
		logger.zap.Error("Failed to create audit log", zap.Error(err))
		return
	}
	logger.zap.Info("Audit log created", zap.String("id", res.Id), zap.String("action", action), zap.String("resource", string(resource)))
}

func (logger *Logger) Info(ctx *gin.Context, action string, resource resources.ResourceType, extra map[string]interface{}) {
	if !logger.resourceManager.ValidateResource(resource) {
		logger.zap.Error("Invalid resource", zap.String("resource", string(resource)))
		return
	}

	logger.zap.Info(action)
	headers := logger.headersHandler.Extract(ctx)
	auditLog := logger.auditLogsService.NewAuditLogFromRequest(headers, action, string(resource), "", extra)
	res, err := logger.auditLogsService.CreateAuditLogs(&auditLog)
	if err != nil {
		logger.zap.Error("Failed to create audit log", zap.Error(err))
		return
	}
	logger.zap.Info("Audit log created", zap.String("id", res.Id), zap.String("action", action), zap.String("resource", string(resource)))
}
