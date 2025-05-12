package repository

import (
	"github.com/gabrielhsdev/dental_ai/backend/db-service/internal/models"
	"github.com/gabrielhsdev/dental_ai/backend/db-service/pkg/database"
)

type AuditLogRepository interface {
	CreateAuditLog(auditLog *models.AuditLogs) (*models.AuditLogs, error)
}

type AuditLogRepositoryImplementation struct {
	DB database.Database
}

func NewAuditLogRepository(db database.Database) AuditLogRepository {
	return &AuditLogRepositoryImplementation{DB: db}
}

func (repository *AuditLogRepositoryImplementation) CreateAuditLog(auditLog *models.AuditLogs) (*models.AuditLogs, error) {
	query := `INSERT INTO audit_logs (requestId, requestIp, requestTimestamp, userId, action, resource, extra) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	err := repository.DB.QueryRow(query,
		auditLog.RequestId,
		auditLog.RequestIp,
		auditLog.RequestTimestamp,
		auditLog.UserId,
		auditLog.Action,
		auditLog.Resource,
		auditLog.Extra).Scan(&auditLog.Id)
	if err != nil {
		return nil, err
	}
	return auditLog, nil
}
