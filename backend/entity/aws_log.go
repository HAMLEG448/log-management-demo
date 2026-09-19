package entity

type AWSCloud struct {
	Service   string `json:"service"`
	AccountID string `json:"account_id"`
	Region    string `json:"region"`
}

type AWSCloudTrailInput struct {
	Tenant    string                 `json:"tenant"`
	Source    string                 `json:"source"`
	Cloud     AWSCloud               `json:"cloud"`
	EventType string                 `json:"event_type"`
	User      string                 `json:"user"`
	Timestamp string                 `json:"@timestamp"`
	Raw       map[string]interface{} `json:"raw"`
}
