package sevices

import (
	"fmt"
	"time"

	"log-management/backend/config"
	"log-management/backend/entity"
)

const (
	failedLoginThreshold = 3
	failedLoginWindow    = 5 * time.Minute
)

func CheckAlertRules(logEntry entity.Log) error {
	if logEntry.SrcIP == "" {
		return nil
	}

	if !isFailedLogin(logEntry) {
		return nil
	}

	windowStart := logEntry.Timestamp.Add(-failedLoginWindow)

	var failedCount int64

	err := config.DB.
		Model(&entity.Log{}).
		Where("src_ip = ?", logEntry.SrcIP).
		Where("timestamp >= ? AND timestamp <= ?", windowStart, logEntry.Timestamp).
		Where(
			"event_type IN ?",
			[]string{
				"app_login_failed",
				"LogonFailed",
			},
		).
		Count(&failedCount).Error

	if err != nil {
		return err
	}

	if failedCount < failedLoginThreshold {
		return nil
	}

	ruleName := "Repeated Failed Login"

	// ป้องกันการสร้าง Alert ซ้ำติด ๆ กัน
	var existingAlertCount int64

	err = config.DB.
		Model(&entity.Alert{}).
		Where("rule_name = ?", ruleName).
		Where("src_ip = ?", logEntry.SrcIP).
		Where("created_at >= ?", windowStart).
		Count(&existingAlertCount).Error

	if err != nil {
		return err
	}

	if existingAlertCount > 0 {
		return nil
	}

	alert := entity.Alert{
		Timestamp: logEntry.Timestamp,
		Tenant:    logEntry.Tenant,
		RuleName:  ruleName,
		Severity:  8,
		SrcIP:     logEntry.SrcIP,
		EventType: logEntry.EventType,

		Message: fmt.Sprintf(
			"Detected %d failed login attempts from IP %s within 5 minutes",
			failedCount,
			logEntry.SrcIP,
		),
	}

	return config.DB.Create(&alert).Error
}

func isFailedLogin(logEntry entity.Log) bool {
	return logEntry.EventType == "app_login_failed" ||
		logEntry.EventType == "LogonFailed"
}

func GetAlerts(tenant string) ([]entity.Alert, error) {
	var alerts []entity.Alert

	query := config.DB.Model(&entity.Alert{})

	if tenant != "" {
		query = query.Where("tenant = ?", tenant)
	}

	result := query.
		Order("timestamp DESC").
		Find(&alerts)

	if result.Error != nil {
		return nil, result.Error
	}

	return alerts, nil
}
