package service

import (
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/internal/models"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/internal/repository"
)

type AuditLogsServiceInterface interface {
	GetAuditLogsById(id int) (*models.AuditLogs, error)
	CreateAuditLogs(auditLog *models.AuditLogs) (*models.AuditLogs, error)
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

/*
func (service *AuditLogsService) NewAuditLogFromRequest(, action string, resource string, message string, extra map[string]any) models.AuditLogs {
}
*/
