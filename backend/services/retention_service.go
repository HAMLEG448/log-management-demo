package sevices

import (
	"log"
	"os"
	"strconv"
	"time"

	"log-management/backend/config"
	"log-management/backend/entity"
)

const (
	defaultRetentionDays         = 7
	defaultRetentionCheckMinutes = 60
)

// StartRetentionWorker starts the automatic log cleanup process.
//
// The worker:
//  1. Cleans expired logs immediately when the backend starts.
//  2. Checks for expired logs again based on the configured interval.
func StartRetentionWorker() {
	retentionDays := getPositiveIntEnv(
		"LOG_RETENTION_DAYS",
		defaultRetentionDays,
	)

	checkMinutes := getPositiveIntEnv(
		"RETENTION_CHECK_INTERVAL_MINUTES",
		defaultRetentionCheckMinutes,
	)

	// Run once immediately when the backend starts.
	runRetentionCleanup(retentionDays)

	// Continue running cleanup in the background.
	go func() {
		ticker := time.NewTicker(
			time.Duration(checkMinutes) * time.Minute,
		)
		defer ticker.Stop()

		for range ticker.C {
			runRetentionCleanup(retentionDays)
		}
	}()

	log.Printf(
		"[Retention] worker started: keep logs for %d days, check every %d minutes",
		retentionDays,
		checkMinutes,
	)
}

// CleanupExpiredLogs deletes logs older than the configured retention period.
//
// It returns the number of deleted rows.
func CleanupExpiredLogs(
	retentionDays int,
) (int64, error) {
	cutoff := time.Now().
		UTC().
		AddDate(0, 0, -retentionDays)

	result := config.DB.
		Where("timestamp < ?", cutoff).
		Delete(&entity.Log{})

	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

// runRetentionCleanup executes the cleanup and writes the result to the log.
func runRetentionCleanup(
	retentionDays int,
) {
	deleted, err := CleanupExpiredLogs(
		retentionDays,
	)

	if err != nil {
		log.Printf(
			"[Retention] cleanup failed: %v",
			err,
		)
		return
	}

	log.Printf(
		"[Retention] cleanup complete: deleted %d expired log(s)",
		deleted,
	)
}

// getPositiveIntEnv reads a positive integer from an environment variable.
// If the value is missing or invalid, it returns the fallback value.
func getPositiveIntEnv(
	name string,
	fallback int,
) int {
	value := os.Getenv(name)

	if value == "" {
		return fallback
	}

	number, err := strconv.Atoi(value)
	if err != nil || number <= 0 {
		log.Printf(
			"[Retention] invalid %s=%q, using default value %d",
			name,
			value,
			fallback,
		)

		return fallback
	}

	return number
}
