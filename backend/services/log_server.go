package sevices

import (
	"time"

	"log-management/backend/config"
	"log-management/backend/entity"
)

type LogFilter struct {
	Tenant    string
	Source    string
	EventType string
	User      string
	Search    string
	From      *time.Time
	To        *time.Time
}

func GetLogs(filter LogFilter) ([]entity.Log, error) {
	var logs []entity.Log

	query := config.DB.Model(&entity.Log{})

	if filter.Tenant != "" {
		query = query.Where("tenant = ?", filter.Tenant)
	}

	if filter.Source != "" {
		query = query.Where("source = ?", filter.Source)
	}

	if filter.EventType != "" {
		query = query.Where("event_type = ?", filter.EventType)
	}

	if filter.User != "" {
		query = query.Where("\"user\" = ?", filter.User)
	}

	if filter.From != nil {
		query = query.Where("timestamp >= ?", *filter.From)
	}

	if filter.To != nil {
		query = query.Where("timestamp <= ?", *filter.To)
	}

	if filter.Search != "" {
		keyword := "%" + filter.Search + "%"

		query = query.Where(
			`event_type ILIKE ?
			OR "user" ILIKE ?
			OR src_ip ILIKE ?
			OR reason ILIKE ?
			OR source ILIKE ?`,
			keyword,
			keyword,
			keyword,
			keyword,
			keyword,
		)
	}

	result := query.
		Order("timestamp DESC").
		Find(&logs)

	if result.Error != nil {
		return nil, result.Error
	}

	return logs, nil
}
