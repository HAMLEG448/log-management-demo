package sevices

import (
	"testing"
	"time"

	"log-management/backend/entity"
)

func TestNormalizeLogMapsFields(t *testing.T) {
	input := entity.IngestRequest{
		Tenant:    "demoA",
		Source:    "api",
		EventType: "app_login_failed",
		User:      "alice",
		IP:        "203.0.113.7",
		Reason:    "wrong_password",
		Timestamp: "2026-09-18T10:00:00Z",
	}

	result := NormalizeLog(input)

	if result.Tenant != input.Tenant {
		t.Errorf(
			"expected tenant %q, got %q",
			input.Tenant,
			result.Tenant,
		)
	}

	if result.Source != input.Source {
		t.Errorf(
			"expected source %q, got %q",
			input.Source,
			result.Source,
		)
	}

	if result.EventType != input.EventType {
		t.Errorf(
			"expected event type %q, got %q",
			input.EventType,
			result.EventType,
		)
	}

	if result.User != input.User {
		t.Errorf(
			"expected user %q, got %q",
			input.User,
			result.User,
		)
	}

	// API input uses IP,
	// normalized schema uses SrcIP.
	if result.SrcIP != input.IP {
		t.Errorf(
			"expected src_ip %q, got %q",
			input.IP,
			result.SrcIP,
		)
	}

	if result.Reason != input.Reason {
		t.Errorf(
			"expected reason %q, got %q",
			input.Reason,
			result.Reason,
		)
	}
}

func TestNormalizeLogParsesTimestamp(t *testing.T) {
	input := entity.IngestRequest{
		Tenant:    "demoA",
		Source:    "api",
		EventType: "test_event",
		Timestamp: "2026-09-18T10:30:00Z",
	}

	result := NormalizeLog(input)

	expected, err := time.Parse(
		time.RFC3339,
		input.Timestamp,
	)
	if err != nil {
		t.Fatalf(
			"failed to create expected timestamp: %v",
			err,
		)
	}

	if !result.Timestamp.Equal(expected) {
		t.Errorf(
			"expected timestamp %v, got %v",
			expected,
			result.Timestamp,
		)
	}
}

func TestNormalizeLogUsesCurrentTimeWhenTimestampMissing(
	t *testing.T,
) {
	before := time.Now().
		UTC().
		Add(-time.Second)

	input := entity.IngestRequest{
		Tenant:    "demoA",
		Source:    "api",
		EventType: "test_event",
	}

	result := NormalizeLog(input)

	after := time.Now().
		UTC().
		Add(time.Second)

	if result.Timestamp.Before(before) {
		t.Errorf(
			"timestamp %v is earlier than expected",
			result.Timestamp,
		)
	}

	if result.Timestamp.After(after) {
		t.Errorf(
			"timestamp %v is later than expected",
			result.Timestamp,
		)
	}
}
