package sevices

import (
	"encoding/json"
	"time"

	"log-management/backend/entity"
)

func NormalizeAWSLog(input entity.AWSCloudTrailInput) (entity.Log, error) {
	timestamp := time.Now().UTC()

	if input.Timestamp != "" {
		parsedTime, err := time.Parse(time.RFC3339, input.Timestamp)
		if err != nil {
			return entity.Log{}, err
		}

		timestamp = parsedTime
	}

	rawJSON, err := json.Marshal(input.Raw)
	if err != nil {
		return entity.Log{}, err
	}

	source := input.Source
	if source == "" {
		source = "aws"
	}

	log := entity.Log{
		Timestamp: timestamp,
		Tenant:    input.Tenant,
		Source:    source,
		Vendor:    "AWS",
		Product:   "CloudTrail",
		EventType: input.EventType,
		User:      input.User,

		CloudAccountID: input.Cloud.AccountID,
		CloudRegion:    input.Cloud.Region,
		CloudService:   input.Cloud.Service,

		Raw: string(rawJSON),
	}

	return log, nil
}
