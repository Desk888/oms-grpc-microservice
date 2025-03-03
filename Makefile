.PHONY: build rebuild up down logs ps proto

# Build all services
build:
	docker compose up --build

# Just rebuild without starting services
rebuild:
	docker compose build

# Run all services in the background
up:
	docker compose up -d

# Run all services in the foreground
run:
	docker compose up

# Stop all services
down:
	docker compose down

# Stop all services and remove volumes
clean:
	docker compose down -v

# Follow logs of all services
logs:
	docker compose logs -f

# Show status of services
ps:
	docker compose ps

# Generate protocol buffers
proto:
	cd common && make

# Run a specific service in development mode
dev-gateway:
	cd gateway && air

dev-orders:
	cd orders-service && air

dev-stock:
	cd stock-service && air

dev-payments:
	cd payments-service && air

dev-kitchen:
	cd kitchen-service && air 