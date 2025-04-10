package repository

import (
	"fmt"

	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/internal/models"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/pkg/database"
)

type AuditLogRepository interface {
	GetAuditLogsById(id int) (*models.AuditLogs, error)
	CreateAuditLog(auditLog *models.AuditLogs) (*models.AuditLogs, error)
}

type AuditLogRepositoryImplementation struct {
	DB database.Database
}

func NewAuditLogRepository(db database.Database) AuditLogRepository {
	return &AuditLogRepositoryImplementation{DB: db}
}

func (repository *AuditLogRepositoryImplementation) GetAuditLogsById(id int) (*models.AuditLogs, error) {
	var auditLog models.AuditLogs
	query := `SELECT id, requestId, requestIp, requestTimestamp, userId, action, resource, extra FROM audit_logs WHERE id = $1`
	err := repository.DB.QueryRow(query, id).Scan(
		&auditLog.Id,
		&auditLog.RequestId,
		&auditLog.RequestIp,
		&auditLog.RequestTimestamp,
		&auditLog.UserId,
		&auditLog.Action,
		&auditLog.Resource,
		&auditLog.Extra)
	if err != nil {
		return nil, err
	}
	return &auditLog, nil
}

func (repository *AuditLogRepositoryImplementation) CreateAuditLog(auditLog *models.AuditLogs) (*models.AuditLogs, error) {
	query := `INSERT INTO audit_logs (requestId, requestIp, requestTimestamp, userId, action, resource, extra) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	// Lets log this request manually, values should be on their proper places
	// INSERT INTO audit_logs (requestId, requestIp, requestTimestamp, userId, action, resource, extra) VALUES (39172.22.0.11744281489.8791, 172.22.0.1, 1744281489.879, , Login, Login successful, {"extra":{"user":"admin@example.com"},"message":""})
	fmt.Printf("INSERT INTO audit_logs (requestId, requestIp, requestTimestamp, userId, action, resource, extra) VALUES (%s, %s, %s, %s, %s, %s, %s)\n",
		auditLog.RequestId,
		auditLog.RequestIp,
		auditLog.RequestTimestamp,
		auditLog.UserId,
		auditLog.Action,
		auditLog.Resource,
		auditLog.Extra)

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
