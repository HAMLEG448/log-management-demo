package entity

import "time"

type Alert struct {
	ID uint `gorm:"primaryKey" json:"id"`

	Timestamp time.Time `json:"timestamp"`

	Tenant   string `json:"tenant"`
	RuleName string `json:"rule_name"`

	Severity int `json:"severity"`

	SrcIP     string `json:"src_ip"`
	EventType string `json:"event_type"`

	Message string `json:"message"`

	CreatedAt time.Time `json:"created_at"`
}
