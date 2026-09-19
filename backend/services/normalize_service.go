package sevices

import (
	"time"

	"log-management/backend/entity"
)

func NormalizeLog(request entity.IngestRequest) entity.Log {
	var timestamp time.Time

	if request.Timestamp == "" {
		timestamp = time.Now().UTC()
	} else {
		parsedTime, err := time.Parse(time.RFC3339, request.Timestamp)

		if err != nil {
			timestamp = time.Now().UTC()
		} else {
			timestamp = parsedTime
		}
	}

	log := entity.Log{
		Timestamp: timestamp,
		Tenant:    request.Tenant,
		Source:    request.Source,
		EventType: request.EventType,
		User:      request.User,
		SrcIP:     request.IP,
		Reason:    request.Reason,
	}

	return log
}
