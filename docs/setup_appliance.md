# Appliance Deployment Guide

## Overview

The Appliance deployment runs the complete Demo Log Management System
on a single machine or virtual machine using Docker Compose.

The deployment contains:

```text
Frontend / Nginx
Go Backend
PostgreSQL
```

## Recommended Environment

- Ubuntu Server 22.04 LTS or newer
- 4 vCPU
- 8 GB RAM
- 40 GB disk
- Docker Engine
- Docker Compose plugin
- Git

## Required Ports

| Port | Protocol | Purpose |
|---|---|---|
| 80 | TCP | Web application |
| 8080 | TCP | Backend API |
| 514 | UDP | Syslog ingestion |

PostgreSQL is available only inside the Docker network by default.

---

## 1. Clone the Repository

```bash
git clone https://github.com/HAMLEG448/log-management-demo.git
cd log-management-demo
```

---

## 2. Configure Environment Variables

Create the root environment file:

```bash
cp .env.example .env
```

Example:

```env
POSTGRES_PASSWORD=change-this-to-a-strong-password
```

Create the backend environment file:

```bash
cp backend/.env.example backend/.env
```

Example:

```env
DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=change-this-to-a-strong-password
DB_NAME=log_management

JWT_SECRET=change-this-to-a-long-random-secret

ADMIN_USERNAME=admin
ADMIN_PASSWORD=change-this-admin-password

VIEWER_USERNAME=viewerA
VIEWER_PASSWORD=change-this-viewer-password

LOG_RETENTION_DAYS=7
RETENTION_CHECK_INTERVAL_MINUTES=60
```

The database password must match the root `POSTGRES_PASSWORD`.

Real `.env` files must not be committed to Git.

---

## 3. Start the System

Run from the project root:

```bash
docker compose up -d --build
```

Check status:

```bash
docker compose ps
```

Expected services:

```text
log-management-db
log-management-backend
log-management-frontend
```

---

## 4. Verify the Application

Open:

```text
http://<SERVER-IP>
```

For a local machine:

```text
http://localhost
```

Health check:

```text
http://<SERVER-IP>/api/health
```

Expected response:

```json
{
  "status": "ok"
}
```

---

## 5. Verify REST API Ingestion

```bash
curl -X POST http://<SERVER-IP>/api/ingest \
  -H "Content-Type: application/json" \
  -d '{
    "tenant":"demoA",
    "source":"api",
    "event_type":"appliance_test",
    "user":"test-user",
    "ip":"203.0.113.10",
    "reason":"appliance_test"
  }'
```

The event should appear on the Logs page.

---

## 6. Verify Active Directory Ingestion

```bash
curl -X POST http://<SERVER-IP>/api/ingest/ad \
  -H "Content-Type: application/json" \
  -d '{
    "tenant":"demoA",
    "source":"ad",
    "event_id":4625,
    "event_type":"LogonFailed",
    "user":"demo\\eve",
    "host":"DC01",
    "ip":"203.0.113.77",
    "logon_type":3
  }'
```

---

## 7. Verify AWS CloudTrail File Batch

Sample file:

```text
samples/aws_cloudtrail.json
```

Upload:

```bash
curl -X POST http://<SERVER-IP>/api/ingest/aws-file \
  -F "file=@samples/aws_cloudtrail.json"
```

---

## 8. Verify Syslog Ingestion

The Docker host accepts Syslog on:

```text
UDP 514
```

Docker forwards it to:

```text
Backend UDP 5514
```

Example:

```text
vendor=demo product=ngfw action=deny src=10.0.1.10 dst=8.8.8.8 spt=5353 dpt=53 proto=udp
```

After sending the message, search:

```text
source = firewall
```

in the Logs page.

---

## 9. Authentication and Tenant Isolation

The system supports:

```text
Admin
Viewer
```

Admin can access multiple tenants.

Viewer accounts are restricted to their assigned tenant.

Tenant isolation is enforced by backend middleware rather than only
by the frontend.

---

## 10. Alert Verification

Current rule:

```text
Repeated Failed Login
```

The alert is created when at least three failed login events from
the same source IP occur within five minutes.

---

## 11. Log Retention

Default:

```env
LOG_RETENTION_DAYS=7
RETENTION_CHECK_INTERVAL_MINUTES=60
```

The retention worker removes expired logs when the backend starts
and then checks again at the configured interval.

---

## 12. View Logs

All containers:

```bash
docker compose logs
```

Backend:

```bash
docker compose logs backend
```

Follow backend logs:

```bash
docker compose logs -f backend
```

---

## 13. Stop the Appliance

```bash
docker compose down
```

This preserves the PostgreSQL volume.

Start again:

```bash
docker compose up -d
```

Do not use:

```bash
docker compose down -v
```

unless the stored database should also be deleted.

---

## Verification Checklist

```text
[ ] Containers are running
[ ] PostgreSQL is healthy
[ ] Web UI opens
[ ] Health endpoint works
[ ] Admin login works
[ ] Viewer login works
[ ] REST API ingestion works
[ ] AD ingestion works
[ ] AWS File Batch works
[ ] Syslog UDP 514 works
[ ] Log search works
[ ] Dashboard works
[ ] Alert rule works
[ ] Tenant isolation works
[ ] Retention worker works
```