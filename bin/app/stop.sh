#!/bin/bash

set -e

cd "$(dirname "$0")/../.."

echo "🛑 Stopping markitos-it-svc-acmes PostgreSQL..."
docker compose down -v --remove-orphans markitos-it-svc-acmes-postgres markitos-it-svc-acmes > /dev/null 2>&1 || true
echo "✅ markitos-it-svc-acmes PostgreSQL stopped."
echo
