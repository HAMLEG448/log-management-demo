# SaaS Deployment Guide

## Overview

The SaaS version of the Demo Log Management System is deployed on
Google Cloud Compute Engine and is accessible through HTTPS.

Current deployment:

```text
Google Cloud Compute Engine
Ubuntu 22.04 LTS
Docker Compose
Nginx
Let's Encrypt
Certbot
```

Current public URL:

```text
https://35.247.183.116
```

---

## Cloud Environment

| Setting | Value |
|---|---|
| Cloud Provider | Google Cloud |
| Compute Service | Compute Engine |
| Region | asia-southeast1 |
| Operating System | Ubuntu 22.04 LTS |
| vCPU | 4 |
| Memory | 8 GB |
| Disk | 40 GB |
| Deployment | Docker Compose |

The VM uses a static public IPv4 address.

---

## Public Ports

| Port | Protocol | Purpose |
|---|---|---|
| 80 | TCP | HTTP redirect and ACME challenge |
| 443 | TCP | HTTPS Web UI and API |
| 514 | UDP | Syslog ingestion |
| 22 | TCP | SSH administration |

PostgreSQL is not exposed publicly.

Normal application traffic reaches the backend through Nginx.

---

## 1. Create the Google Cloud VM

Create a Compute Engine VM using:

```text
Ubuntu 22.04 LTS
4 vCPU
8 GB RAM
40 GB disk
```

Reserve and attach a static public IPv4 address.

Enable HTTP and HTTPS traffic.

---

## 2. Configure Firewall Rules

Allow inbound traffic for:

```text
TCP 80
TCP 443
UDP 514
```

UDP 514 is used for external Syslog ingestion.

For production use, the source range for UDP 514 should be restricted
to trusted senders whenever possible.

---

## 3. Install Docker

Install Docker Engine and the Docker Compose plugin.

Verify:

```bash
docker --version
```

```bash
docker compose version
```

---

## 4. Clone the Repository

```bash
git clone https://github.com/HAMLEG448/log-management-demo.git
cd log-management-demo
```

---

## 5. Configure Environment Variables

Create:

```bash
cp .env.example .env
```

Example:

```env
POSTGRES_PASSWORD=change-this-to-a-strong-password
```

Create:

```bash
cp backend/.env.example backend/.env
```

Configure:

```env
JWT_SECRET=change-this-to-a-long-random-secret

ADMIN_USERNAME=admin
ADMIN_PASSWORD=change-this-admin-password

VIEWER_USERNAME=viewerA
VIEWER_PASSWORD=change-this-viewer-password

LOG_RETENTION_DAYS=7
RETENTION_CHECK_INTERVAL_MINUTES=60
```

Real secrets must not be committed to Git.

---

## 6. Start the Initial Deployment

Run:

```bash
docker compose up -d --build
```

Check:

```bash
docker compose ps
```

Verify:

```text
http://<PUBLIC-IP>/api/health
```

Expected response:

```json
{
  "status": "ok"
}
```

---

## 7. Configure Cloud Nginx

The SaaS deployment uses a cloud-specific Nginx configuration.

The configuration provides:

```text
HTTP port 80
HTTPS port 443
ACME challenge path
HTTP to HTTPS redirect
/api reverse proxy
```

The ACME challenge path is:

```text
/.well-known/acme-challenge/
```

---

## 8. Install Certbot

Install Certbot on the VM.

Verify:

```bash
certbot --version
```

---

## 9. Request a Let's Encrypt Certificate

The current deployment uses a Let's Encrypt short-lived certificate
for the static public IP.

Example:

```bash
sudo certbot certonly \
  --preferred-profile shortlived \
  --webroot \
  --webroot-path "$PWD/certbot/www" \
  --ip-address 35.247.183.116 \
  --cert-name log-management-ip
```

Certificate files:

```text
/etc/letsencrypt/live/log-management-ip/fullchain.pem
/etc/letsencrypt/live/log-management-ip/privkey.pem
```

The private key must never be committed to Git.

---

## 10. Start the HTTPS Deployment

The cloud deployment uses the base Compose file and a cloud override:

```bash
docker compose \
  -f docker-compose.yml \
  -f docker-compose.cloud.yml \
  up -d --build
```

Check:

```bash
docker compose \
  -f docker-compose.yml \
  -f docker-compose.cloud.yml \
  ps
```

---

## 11. Verify HTTPS

HTTPS:

```bash
curl -I https://35.247.183.116
```

Expected:

```text
HTTP/1.1 200 OK
```

HTTP redirect:

```bash
curl -I http://35.247.183.116
```

Expected:

```text
HTTP/1.1 301 Moved Permanently
```

Health endpoint:

```bash
curl https://35.247.183.116/api/health
```

Expected:

```json
{
  "status": "ok"
}
```

---

## 12. Verify REST Ingestion

```bash
curl -X POST https://35.247.183.116/api/ingest \
  -H "Content-Type: application/json" \
  -d '{
    "tenant":"demoA",
    "source":"api",
    "event_type":"cloud_test",
    "user":"cloud-user",
    "ip":"203.0.113.50",
    "reason":"google_cloud_test"
  }'
```

The event should appear on the Logs page.

---

## 13. Verify AWS CloudTrail File Batch

```bash
curl -X POST https://35.247.183.116/api/ingest/aws-file \
  -F "file=@samples/aws_cloudtrail.json"
```

---

## 14. Verify Active Directory Ingestion

```bash
curl -X POST https://35.247.183.116/api/ingest/ad \
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

## 15. Verify External Syslog

Send a UDP message to:

```text
35.247.183.116:514
```

Example:

```text
vendor=demo product=ngfw action=deny src=10.0.1.10 dst=8.8.8.8 spt=5353 dpt=53 proto=udp
```

Flow:

```text
External Device
      |
      | UDP 514
      v
Google Cloud Firewall
      |
      v
VM
      |
      v
Docker
      |
      | UDP 5514
      v
Go Syslog Listener
      |
      v
PostgreSQL
```

---

## 16. Verify Alert Rule

Send three failed login events from the same IP within five minutes.

Confirm that:

```text
Repeated Failed Login
```

appears on the Alerts page.

---

## 17. Verify Tenant Isolation

Login as a Viewer assigned to:

```text
demoA
```

The Viewer must not be able to retrieve:

```text
demoB
```

logs even when query parameters are manually modified.

---

## 18. Verify Retention

Configuration:

```env
LOG_RETENTION_DAYS=7
RETENTION_CHECK_INTERVAL_MINUTES=60
```

Cloud verification procedure:

```text
1. Insert a log older than seven days
2. Confirm the log was ingested
3. Wait for retention cleanup
4. Confirm one expired log was deleted
5. Query PostgreSQL
6. Confirm the expired row no longer exists
```

---

## Certificate Renewal

Certbot handles certificate renewal.

Because Nginx runs inside Docker, Nginx must be reloaded after
certificate renewal.

Example deploy hook:

```bash
#!/bin/bash

cd /home/<USER>/log-management-demo || exit 1

/usr/bin/docker compose \
  -f docker-compose.yml \
  -f docker-compose.cloud.yml \
  exec -T frontend nginx -s reload
```

The actual VM username should replace:

```text
<USER>
```

---

## Starting the VM Again

After starting the Google Cloud VM, verify:

```bash
docker compose \
  -f docker-compose.yml \
  -f docker-compose.cloud.yml \
  ps
```

Then:

```bash
curl https://35.247.183.116/api/health
```

Containers use:

```text
restart: unless-stopped
```

so they normally start again automatically after the VM boots.

---

## Stopping the VM

The VM can be stopped from:

```text
Google Cloud Console
→ Compute Engine
→ VM instances
→ log-management-vm
→ Stop
```

Do not delete the VM unless the cloud deployment is no longer needed.

---

## Security Notes

Implemented controls include:

```text
HTTPS / TLS
JWT authentication
bcrypt password hashing
Admin / Viewer roles
Backend tenant isolation
Google Cloud firewall
PostgreSQL not exposed publicly
Environment secrets excluded from Git
Strong live credentials
```

Recommended production hardening:

```text
Login rate limiting
Authentication for ingestion endpoints
Restricted Syslog source ranges
Google Secret Manager
Audit logging
Database least-privilege account
```

---

## SaaS Verification Checklist

```text
[ ] VM is running
[ ] Docker services are running
[ ] PostgreSQL is healthy
[ ] HTTPS works
[ ] HTTP redirects to HTTPS
[ ] Health endpoint works
[ ] Admin login works
[ ] Viewer login works
[ ] REST ingestion works
[ ] AWS File Batch works
[ ] Active Directory ingestion works
[ ] External Syslog works
[ ] Search works
[ ] Dashboard works
[ ] Alert rule works
[ ] Tenant isolation works
[ ] Retention worker works
```