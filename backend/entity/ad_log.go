package entity

type ADLogInput struct {
	Tenant    string `json:"tenant" binding:"required"`
	Source    string `json:"source"`
	EventID   int    `json:"event_id"`
	EventType string `json:"event_type"`
	User      string `json:"user"`
	Host      string `json:"host"`
	IP        string `json:"ip"`
	LogonType int    `json:"logon_type"`
	Timestamp string `json:"@timestamp"`
}
