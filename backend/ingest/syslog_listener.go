package ingest

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"log-management/backend/config"
	"log-management/backend/entity"
)

func StartSyslogListener() {
	addr, err := net.ResolveUDPAddr("udp", ":5514")
	if err != nil {
		fmt.Println("Failed to resolve syslog address:", err)
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Failed to start syslog listener:", err)
		return
	}

	defer conn.Close()

	fmt.Println("Syslog listener running on UDP :5514")

	buffer := make([]byte, 4096)

	for {
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Failed to read syslog:", err)
			continue
		}

		message := string(buffer[:n])

		fmt.Printf("Syslog from %s: %s\n", remoteAddr, message)

		log := parseFirewallSyslog(message)

		if err := config.DB.Create(&log).Error; err != nil {
			fmt.Println("Failed to save syslog:", err)
		}
	}
}

func parseFirewallSyslog(message string) entity.Log {
	log := entity.Log{
		Timestamp: time.Now().UTC(),
		Tenant:    "demoA",
		Source:    "firewall",
		EventType: "firewall_event",
		Raw:       message,
	}

	fields := strings.Fields(message)

	for _, field := range fields {
		parts := strings.SplitN(field, "=", 2)

		if len(parts) != 2 {
			continue
		}

		key := parts[0]
		value := parts[1]

		switch key {
		case "vendor":
			log.Vendor = value

		case "product":
			log.Product = value

		case "action":
			log.Action = value

		case "src":
			log.SrcIP = value

		case "dst":
			log.DstIP = value

		case "spt":
			port, _ := strconv.Atoi(value)
			log.SrcPort = port

		case "dpt":
			port, _ := strconv.Atoi(value)
			log.DstPort = port

		case "proto":
			log.Protocol = value
		}
	}

	return log
}
