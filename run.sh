#!/usr/bin/env bash

set -e

COMMAND="${1:-start}"

case "$COMMAND" in
  start)
    echo "Starting Log Management System..."
    docker compose up -d --build
    docker compose ps
    ;;

  stop)
    echo "Stopping Log Management System..."
    docker compose down
    ;;

  restart)
    echo "Restarting Log Management System..."
    docker compose down
    docker compose up -d
    docker compose ps
    ;;

  status)
    docker compose ps
    ;;

  logs)
    docker compose logs -f
    ;;

  *)
    echo "Usage:"
    echo "  ./run.sh start"
    echo "  ./run.sh stop"
    echo "  ./run.sh restart"
    echo "  ./run.sh status"
    echo "  ./run.sh logs"
    exit 1
    ;;
esac