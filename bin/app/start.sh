#!/bin/bash

set -e

cd "$(dirname "$0")/../.."

make app-proto

export GRPC_PORT=${GRPC_PORT:-3000}
export DB_HOST=${DB_HOST:-localhost}
export DB_PORT=${DB_PORT:-55432}
export DB_USER=${DB_USER:-markitos-it-svc-acmes}
export DB_PASS=${DB_PASS:-markitos-it-svc-acmes}
export DB_NAME=${DB_NAME:-markitos-it-svc-acmes}

echo "🚀 Starting markitos-it-svc-acmes (Go)..."
echo "📡 GRPC_PORT: $GRPC_PORT"
echo "📦 DB_HOST..: $DB_HOST:$DB_PORT"
echo "📦 DB_USER..: $DB_USER"
echo "📦 DB_NAME..: $DB_NAME"
echo ""

docker compose up