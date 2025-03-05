# Order Management System (OMS) - gRPC Microservices

This project is a microservices-based Order Management System using gRPC for service-to-service communication and a REST API gateway for client communication.

## Architecture

The system consists of the following services:

- **Gateway**: REST API that clients communicate with, which forwards requests to the appropriate microservices via gRPC
- **Orders Service**: Handles order creation and management
- **Stock Service**: Manages inventory and stock levels
- **Payments Service**: Processes payments for orders
- **Kitchen Service**: Manages food preparation and order fulfillment

## Technologies Used

- Go (Golang)
- gRPC/Protocol Buffers
- MongoDB
- Docker/Docker Compose

## Prerequisites

- Go 1.21+
- Docker and Docker Compose
- Make (optional, for using the Makefile commands)

## Getting Started

### Running with Docker Compose

To run all services with Docker Compose:

```bash
# Build all services
make build

# Run all services in the background
make up

# Or run in the foreground to see logs
make run
```

The API will be available at http://localhost:8080

### Development Mode

To run individual services in development mode with hot-reload:

```bash
# Run the gateway in development mode
make dev-gateway

# Run the orders service in development mode
make dev-orders

# Similarly for other services
make dev-stock
make dev-payments
make dev-kitchen
```

## API Endpoints

- **Create Order**: `POST /api/customers/{customerID}/orders`
  - Request body: Array of items with quantities
  - Response: Order details

## Project Structure

```
.
├── common/                  # Shared code and protobuf definitions
│   └── api/                 # Protocol buffer definitions
├── gateway/                 # API Gateway service
├── orders-service/          # Orders microservice
├── stock-service/           # Stock/inventory microservice
├── payments-service/        # Payments microservice
├── kitchen-service/         # Kitchen microservice
├── docker-compose.yml       # Docker Compose configuration
└── Makefile                 # Useful commands
```

## Environment Variables

Each service has its own .env file with service-specific configuration.
