#!/bin/bash

set -e

cd "$(dirname "$0")/../.."

VERSION=$1

if [ -z "$VERSION" ]; then
    echo "❌ Error: Version is required"
    echo "Usage: make app-delete-tag n.n.n"
    exit 1
fi

if ! [[ $VERSION =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    echo "❌ Error: Invalid semver format"
    echo "Version must be in format n.n.n (e.g., 1.2.3)"
    echo "Only numbers and dots are allowed"
    exit 1
fi

TAG="${VERSION}"

if ! git rev-parse "$TAG" >/dev/null 2>&1; then
    echo "❌ Error: Tag $TAG does not exist locally"
    exit 1
fi

echo "🗑️  Deleting git tag: ${TAG}"

git tag -d ${TAG}

echo "📤 Deleting tag from GitHub..."

git push origin :refs/tags/${TAG}

echo "✅ Tag ${TAG} deleted successfully!"
