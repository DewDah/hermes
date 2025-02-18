#!/bin/bash

# Exit on any error
set -e

echo "🚀 Initializing Hermes Project..."

# Step 1: Initialize Go Module
echo "📦 Creating Go module..."
go mod init github.com/DewDah/hermes

# Step 2: Install Dependencies
echo "📥 Installing dependencies..."
go get github.com/gin-gonic/gin
go get github.com/go-redis/redis/v8
go get github.com/joho/godotenv
go get github.com/swaggo/swag/cmd/swag
go get github.com/swaggo/gin-swagger
go get github.com/swaggo/files
go get github.com/prometheus/client_golang/prometheus
go get github.com/prometheus/client_golang/prometheus/promhttp
go get github.com/stretchr/testify
go get golang.org/x/net/context

# Step 3: Create Project Directory Structure
echo "📂 Creating project folders..."
mkdir -p config handlers redis tests utils docs

# Step 4: Create Placeholder Files
echo "📝 Creating empty files..."
touch main.go .env .gitignore
touch config/config.go
touch handlers/{auth.go,publish.go,subscribe.go,streams.go,metrics.go,logs.go,health.go}
touch redis/{client.go,streams.go,pubsub.go}
touch tests/{auth_test.go,publish_test.go,subscribe_test.go,streams_test.go}
touch docs/swagger.json docs/swagger.yaml

# Step 5: Initialize Swagger Docs
echo "📖 Initializing Swagger documentation..."
swag init

echo "✅ Hermes Project Setup Complete!"