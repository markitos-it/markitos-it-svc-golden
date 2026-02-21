#!/bin/bash
set -e

echo "🔍 Checking for protoc compiler..."
if ! command -v protoc &> /dev/null; then
    echo "📦 Installing protoc..."
    sudo apt update && sudo apt install -y protobuf-compiler && protoc --version
fi

echo "🔍 Checking for protoc-gen-go plugin..."
if ! command -v protoc-gen-go &> /dev/null; then
    echo "📦 Installing protoc-gen-go..."
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
fi

echo "🔍 Checking for protoc-gen-go-grpc plugin..."
if ! command -v protoc-gen-go-grpc &> /dev/null; then
    echo "📦 Installing protoc-gen-go-grpc..."
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
fi
echo "✅ Protoc plugins are installed"




echo "🔄 Running protoc to generate Go code from proto files..."
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/acme.proto
echo "✅ Protobuf code generated"
