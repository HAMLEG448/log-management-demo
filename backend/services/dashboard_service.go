package sevices

import (
	"time"

	"log-management/backend/config"
	"log-management/backend/entity"

	"gorm.io/gorm"
)

type DashboardFilter struct {
	Tenant string
	Source string
	From   *time.Time
	To     *time.Time
}

type DashboardCountItem struct {
	Name  string `json:"name"`
	Count int64  `json:"count"`
}

type DashboardTimelineItem struct {
	Time  time.Time `json:"time"`
	Count int64     `json:"count"`
}

type DashboardSummary struct {
	TotalLogs     int64                   `json:"total_logs"`
	TopIPs        []DashboardCountItem    `json:"top_ips"`
	TopUsers      []DashboardCountItem    `json:"top_users"`
	TopEventTypes []DashboardCountItem    `json:"top_event_types"`
	Timeline      []DashboardTimelineItem `json:"timeline"`
}

func dashboardBaseQuery(filter DashboardFilter) *gorm.DB {
	query := config.DB.Model(&entity.Log{})

	if filter.Tenant != "" {
		query = query.Where("tenant = ?", filter.Tenant)
	}

	if filter.Source != "" {
		query = query.Where("source = ?", filter.Source)
	}

	if filter.From != nil {
		query = query.Where("timestamp >= ?", *filter.From)
	}

	if filter.To != nil {
		query = query.Where("timestamp <= ?", *filter.To)
	}

	return query
}

func GetDashboardSummary(filter DashboardFilter) (DashboardSummary, error) {
	var summary DashboardSummary

	// จำนวน Log ทั้งหมด
	if err := dashboardBaseQuery(filter).
		Count(&summary.TotalLogs).Error; err != nil {
		return summary, err
	}

	// Top IP
	if err := dashboardBaseQuery(filter).
		Select("src_ip AS name, COUNT(*) AS count").
		Where("src_ip IS NOT NULL AND src_ip <> ''").
		Group("src_ip").
		Order("count DESC").
		Limit(5).
		Scan(&summary.TopIPs).Error; err != nil {
		return summary, err
	}

	// Top User
	if err := dashboardBaseQuery(filter).
		Select(`"user" AS name, COUNT(*) AS count`).
		Where(`"user" IS NOT NULL AND "user" <> ''`).
		Group(`"user"`).
		Order("count DESC").
		Limit(5).
		Scan(&summary.TopUsers).Error; err != nil {
		return summary, err
	}

	// Top Event Type
	if err := dashboardBaseQuery(filter).
		Select("event_type AS name, COUNT(*) AS count").
		Where("event_type IS NOT NULL AND event_type <> ''").
		Group("event_type").
		Order("count DESC").
		Limit(5).
		Scan(&summary.TopEventTypes).Error; err != nil {
		return summary, err
	}

	// Timeline แบ่งเป็นรายชั่วโมง
	if err := dashboardBaseQuery(filter).
		Select("date_trunc('hour', timestamp) AS time, COUNT(*) AS count").
		Group("time").
		Order("time ASC").
		Scan(&summary.Timeline).Error; err != nil {
		return summary, err
	}

	return summary, nil
}
