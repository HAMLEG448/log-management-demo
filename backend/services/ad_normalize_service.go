package sevices

import (
	"encoding/json"
	"time"

	"log-management/backend/entity"
)

func NormalizeADLog(input entity.ADLogInput) (entity.Log, error) {
	timestamp := time.Now().UTC()

	if input.Timestamp != "" {
		parsedTime, err := time.Parse(time.RFC3339, input.Timestamp)
		if err != nil {
			return entity.Log{}, err
		}

		timestamp = parsedTime
	}

	source := input.Source
	if source == "" {
		source = "ad"
	}

	rawJSON, err := json.Marshal(input)
	if err != nil {
		return entity.Log{}, err
	}

	log := entity.Log{
		Timestamp: timestamp,
		Tenant:    input.Tenant,
		Source:    source,

		Vendor:  "Microsoft",
		Product: "Windows Security",

		EventType: input.EventType,
		EventID:   input.EventID,
		LogonType: input.LogonType,

		User:  input.User,
		Host:  input.Host,
		SrcIP: input.IP,

		Raw: string(rawJSON),
	}

	return log, nil
}
