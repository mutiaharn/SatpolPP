.PHONY: all dev build clean docker-up docker-down

all: build

# Run both Frontend & Backend concurrently for local development
dev:
	@echo "Starting Satpol PP Sulsel Monorepo in development mode..."
	@scripts/start_dev.bat

# Build production artifacts for both Frontend & Backend
build:
	@echo "Building Backend (Golang)..."
	@cd backend && go build -o api.exe ./cmd/api/main.go
	@echo "Building Frontend (SvelteKit)..."
	@cd frontend && npm run build
	@echo "Build complete."

# Run entire stack with Docker Compose
docker-up:
	@docker compose -f deployments/docker-compose.yml up --build -d

docker-down:
	@docker compose -f deployments/docker-compose.yml down

clean:
	@echo "Cleaning temporary files..."
	@rm -rf frontend/build frontend/.svelte-kit backend/api.exe
