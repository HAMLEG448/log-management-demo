#!/usr/bin/env bash

HOST="${1:-127.0.0.1}"
PORT="${2:-514}"

MESSAGE="vendor=demo product=ngfw action=deny src=10.0.1.10 dst=8.8.8.8 spt=5353 dpt=53 proto=udp"

python3 - "$HOST" "$PORT" "$MESSAGE" <<'PY'
import socket
import sys

host = sys.argv[1]
port = int(sys.argv[2])
message = sys.argv[3]

sock = socket.socket(
    socket.AF_INET,
    socket.SOCK_DGRAM,
)

sock.sendto(
    message.encode("utf-8"),
    (host, port),
)

sock.close()

print(
    f"Syslog sent to {host}:{port}"
)
PY