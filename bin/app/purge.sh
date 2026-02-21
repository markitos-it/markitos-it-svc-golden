#!/bin/bash
set -e

echo "🧹 Cleaning build artifacts..."

echo "🧹 Cleaning Docker containers, images, and volumes..."
docker rm markitos-it-svc-acmes markitos-it-svc-acmes-postgres 2>/dev/null || true
docker image rm postgres:17-alpine markitos-it-svc-golden-template-markitos-it-svc-acmes:latest 2>/dev/null || true
docker volume rm markitos-it-svc-golden-template_markitos-it-svc-acmes-postgres_data 2>/dev/null || true
echo "✅ Removed Docker containers, images, and volumes"

echo "🧹 Cleaning generated protobuf files..."
rm -f proto/acme.pb.go
rm -f proto/acme_grpc.pb.go
echo "✅ Removed generated protobuf files"

echo "✨ Clean complete"