package main

import (
	"log"

	"log-management/backend/config"
	"log-management/backend/ingest"
	"log-management/backend/routers"
	sevices "log-management/backend/services"
)

func main() {
	// Connect to PostgreSQL and run migrations.
	config.ConnectDatabase()

	// Create demo Admin and Viewer accounts when they do not exist.
	sevices.SeedUsers()

	// Start automatic 7-day log retention.
	sevices.StartRetentionWorker()

	// Start UDP Syslog listener in the background.
	go ingest.StartSyslogListener()

	// Start HTTP API server.
	router := routers.SetupRouter()

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
