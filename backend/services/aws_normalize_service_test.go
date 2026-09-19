package sevices

import (
	"testing"
	"time"

	"log-management/backend/entity"
)

func TestNormalizeAWSLog(t *testing.T) {
	input := entity.AWSCloudTrailInput{
		Tenant: "demoB",
		Source: "aws",

		Cloud: entity.AWSCloud{
			Service:   "iam",
			AccountID: "123456789012",
			Region:    "ap-southeast-1",
		},

		EventType: "CreateUser",
		User:      "admin",

		Timestamp: "2026-09-18T11:00:00Z",

		Raw: map[string]interface{}{
			"eventName": "CreateUser",
			"requestParameters": map[string]interface{}{
				"userName": "temp-user",
			},
		},
	}

	result, err := NormalizeAWSLog(input)
	if err != nil {
		t.Fatalf(
			"NormalizeAWSLog returned error: %v",
			err,
		)
	}

	if result.Tenant != "demoB" {
		t.Errorf(
			"expected tenant demoB, got %q",
			result.Tenant,
		)
	}

	if result.Source != "aws" {
		t.Errorf(
			"expected source aws, got %q",
			result.Source,
		)
	}

	if result.Vendor != "AWS" {
		t.Errorf(
			"expected vendor AWS, got %q",
			result.Vendor,
		)
	}

	if result.Product != "CloudTrail" {
		t.Errorf(
			"expected product CloudTrail, got %q",
			result.Product,
		)
	}

	if result.EventType != "CreateUser" {
		t.Errorf(
			"expected event type CreateUser, got %q",
			result.EventType,
		)
	}

	if result.User != "admin" {
		t.Errorf(
			"expected user admin, got %q",
			result.User,
		)
	}

	if result.CloudService != "iam" {
		t.Errorf(
			"expected cloud service iam, got %q",
			result.CloudService,
		)
	}

	if result.CloudAccountID != "123456789012" {
		t.Errorf(
			"expected account ID 123456789012, got %q",
			result.CloudAccountID,
		)
	}

	if result.CloudRegion != "ap-southeast-1" {
		t.Errorf(
			"expected region ap-southeast-1, got %q",
			result.CloudRegion,
		)
	}

	expectedTime, err := time.Parse(
		time.RFC3339,
		input.Timestamp,
	)
	if err != nil {
		t.Fatalf(
			"failed to parse expected timestamp: %v",
			err,
		)
	}

	if !result.Timestamp.Equal(expectedTime) {
		t.Errorf(
			"expected timestamp %v, got %v",
			expectedTime,
			result.Timestamp,
		)
	}

	if result.Raw == "" {
		t.Error(
			"expected raw AWS log to be stored",
		)
	}
}
