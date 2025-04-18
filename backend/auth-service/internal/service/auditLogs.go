package service

import (
	"encoding/json"

	"github.com/gabrielhsdev/dental_ai/backend/auth-service/internal/models"
	"github.com/gabrielhsdev/dental_ai/backend/auth-service/internal/repository"
	"github.com/gabrielhsdev/dental_ai/backend/auth-service/pkg/headers"
)

type AuditLogsServiceInterface interface {
	GetAuditLogsById(id int) (*models.AuditLogs, error)
	CreateAuditLogs(auditLog *models.AuditLogs) (*models.AuditLogs, error)
	NewAuditLogFromRequest(headers headers.HeadersInterface, action string, resource string, message string, extra map[string]any) models.AuditLogs
}

type AuditLogsService struct {
	Repository repository.AuditLogRepository
}

func NewAuditLogsService(repository repository.AuditLogRepository) AuditLogsServiceInterface {
	return &AuditLogsService{Repository: repository}
}

func (service *AuditLogsService) GetAuditLogsById(id int) (*models.AuditLogs, error) {
	return service.Repository.GetAuditLogsById(id)
}

func (service *AuditLogsService) CreateAuditLogs(auditLog *models.AuditLogs) (*models.AuditLogs, error) {
	return service.Repository.CreateAuditLog(auditLog)
}

func (service *AuditLogsService) NewAuditLogFromRequest(headers headers.HeadersInterface, action string, resource string, message string, extra map[string]any) models.AuditLogs {
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
