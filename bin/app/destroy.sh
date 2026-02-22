#!/bin/bash
set -e

echo "🧹 Cleaning build artifacts..."

echo "🧹 Cleaning Docker containers, images, and volumes..."
docker rm markitos-it-svc-goldens markitos-it-svc-goldens-postgres 2>/dev/null || true
docker image rm postgres:17-alpine markitos-it-svc-golden-markitos-it-svc-goldens:latest 2>/dev/null || true
docker volume rm markitos-it-svc-golden_markitos-it-svc-goldens-postgres_data 2>/dev/null || true
echo "✅ Removed Docker containers, images, and volumes"

echo "🧹 Cleaning generated protobuf files..."
rm -f proto/*.pb.go
rm -f proto/*_grpc.pb.go
echo "✅ Removed generated protobuf files"

echo "🧹 Cleaning Go build cache..."
go clean -testcache
go clean -cache
go clean -modcache
echo "✅ Cleaned Go build cache"

echo "✨ Clean complete"