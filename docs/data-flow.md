# Data Flow

## Overview

Logs can enter the system through multiple ingestion methods.

Supported demo sources:

1. REST API
2. Firewall / Network Syslog
3. Active Directory / Windows
4. AWS CloudTrail File Batch

The system normalizes different source formats into a common log schema before storing them in PostgreSQL.

---

## REST API Flow

```mermaid
flowchart LR
    SOURCE[Application]
    API[POST /ingest]
    NORMALIZE[Normalize Service]
    DB[(PostgreSQL)]
    ALERT[Alert Engine]

    SOURCE -->|HTTP JSON| API
    API --> NORMALIZE
    NORMALIZE --> DB
    NORMALIZE --> ALERT
```

Example input:

```json
{
  "tenant": "demoA",
  "source": "api",
  "event_type": "app_login_failed",
  "user": "alice",
  "ip": "203.0.113.7",
  "reason": "wrong_password"
}
```

The field:

```text
ip
```

is normalized into:

```text
src_ip
```

before storage.

---

## Firewall Syslog Flow

```mermaid
flowchart LR
    FIREWALL[Firewall]
    HOST[Appliance UDP 514]
    DOCKER[Docker Port Mapping]
    LISTENER[Go Syslog Listener UDP 5514]
    PARSER[Syslog Parser]
    DB[(PostgreSQL)]

    FIREWALL -->|UDP 514| HOST
    HOST --> DOCKER
    DOCKER -->|UDP 5514| LISTENER
    LISTENER --> PARSER
    PARSER --> DB
```

Example Syslog:

```text
vendor=demo product=ngfw action=deny src=10.0.1.10 dst=8.8.8.8 spt=5353 dpt=53 proto=udp
```

The parser extracts fields such as:

```text
vendor
product
action
src_ip
dst_ip
src_port
dst_port
protocol
```

---

## Active Directory Flow

```mermaid
flowchart LR
    AD[Active Directory]
    ENDPOINT[POST /ingest/ad]
    NORMALIZE[AD Normalization]
    DB[(PostgreSQL)]
    ALERT[Alert Engine]

    AD -->|HTTP JSON| ENDPOINT
    ENDPOINT --> NORMALIZE
    NORMALIZE --> DB
    NORMALIZE --> ALERT
```

Example event:

```text
Event ID: 4625
Event Type: LogonFailed
```

---

## AWS CloudTrail Flow

```mermaid
flowchart LR
    FILE[AWS CloudTrail JSON]
    UPLOAD[POST /ingest/aws-file]
    PARSER[JSON Parser]
    NORMALIZE[AWS Normalization]
    DB[(PostgreSQL)]

    FILE -->|Multipart File| UPLOAD
    UPLOAD --> PARSER
    PARSER --> NORMALIZE
    NORMALIZE --> DB
```

The AWS normalizer maps fields such as:

```text
cloud.service
cloud.account_id
cloud.region
event_type
user
```

into the common schema.

---

## Search Flow

```mermaid
flowchart LR
    USER[User]
    UI[React Logs Page]
    API[GET /logs]
    AUTH[JWT Middleware]
    TENANT[Tenant Middleware]
    QUERY[Log Query Service]
    DB[(PostgreSQL)]

    USER --> UI
    UI --> API
    API --> AUTH
    AUTH --> TENANT
    TENANT --> QUERY
    QUERY --> DB
    DB --> QUERY
    QUERY --> UI
```

Supported filters include:

```text
tenant
source
event_type
user
search
from
to
```

---

## Dashboard Flow

```mermaid
flowchart LR
    UI[Dashboard]
    API[GET /dashboard]
    TENANT[Tenant Filter]
    DB[(PostgreSQL)]
    RESULT[Aggregated Results]

    UI --> API
    API --> TENANT
    TENANT --> DB
    DB --> RESULT
    RESULT --> UI
```

Dashboard aggregation includes:

- Total Logs
- Timeline
- Top IP
- Top Users
- Top Event Types

---

## Alert Flow

```mermaid
flowchart LR
    LOG[Normalized Log]
    RULE[Alert Rule Engine]
    QUERY[5 Minute Window Query]
    ALERT[Create Alert]
    DB[(PostgreSQL)]

    LOG --> RULE
    RULE --> QUERY
    QUERY --> DB
    QUERY -->|Threshold reached| ALERT
    ALERT --> DB
```

Current demo rule:

```text
Repeated Failed Login
```

An alert is generated when at least three failed login events are detected from the same source IP within five minutes.

---

## Retention Flow

```mermaid
flowchart LR
    WORKER[Retention Worker]
    CUTOFF[Calculate Cutoff]
    DB[(PostgreSQL)]

    WORKER --> CUTOFF
    CUTOFF -->|timestamp older than 7 days| DB
```

Default retention period:

```text
7 days
```

The value can be configured using:

```text
LOG_RETENTION_DAYS
```