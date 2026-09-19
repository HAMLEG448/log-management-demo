package entity

import "time"

type IngestRequest struct {
	Tenant    string `json:"tenant" binding:"required"`
	Source    string `json:"source" binding:"required"`
	EventType string `json:"event_type" binding:"required"`
	User      string `json:"user"`
	IP        string `json:"ip"`
	Reason    string `json:"reason"`
	Timestamp string `json:"@timestamp"`
}

type Log struct {
	ID uint `gorm:"primaryKey" json:"id"`

	Timestamp time.Time `json:"@timestamp"`

	Tenant    string `json:"tenant"`
	Source    string `json:"source"`
	Vendor    string `json:"vendor,omitempty"`
	Product   string `json:"product,omitempty"`
	EventType string `json:"event_type"`
	EventID   int    `json:"event_id,omitempty"`
	LogonType int    `json:"logon_type,omitempty"`

	Severity int    `json:"severity,omitempty"`
	Action   string `json:"action,omitempty"`

	SrcIP    string `json:"src_ip,omitempty"`
	SrcPort  int    `json:"src_port,omitempty"`
	DstIP    string `json:"dst_ip,omitempty"`
	DstPort  int    `json:"dst_port,omitempty"`
	Protocol string `json:"protocol,omitempty"`

	CloudAccountID string `json:"cloud_account_id,omitempty"`
	CloudRegion    string `json:"cloud_region,omitempty"`
	CloudService   string `json:"cloud_service,omitempty"`

	User   string `json:"user,omitempty"`
	Host   string `json:"host,omitempty"`
	Reason string `json:"reason,omitempty"`

	Raw string `json:"raw,omitempty"`

	CreatedAt time.Time `json:"created_at"`
}
