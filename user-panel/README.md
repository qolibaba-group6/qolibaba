
# Final Production Ready Project

## Overview
This project is a modular microservice architecture featuring:
- Authentication and Authorization (JWT-based)
- Wallet and User Management
- API Gateway with Envoy
- Monitoring with Prometheus and Grafana
- Caching with Redis
- Queueing with Kafka

## Features
1. **Authentication and Authorization**: Secure user authentication with JWT.
2. **Wallet Service**: Manage user wallets and transactions.
3. **Monitoring**: Full-stack observability with Prometheus and Grafana.
4. **Queueing**: Event-driven architecture with Kafka.
5. **Caching**: Optimized performance with Redis.

## Getting Started

### Prerequisites
- Docker & Docker-Compose
- Go 1.17 or later
- PostgreSQL

### Running the Project
```bash
docker-compose up --build
```

### Testing
Run all unit tests:
```bash
go test ./... -v
```

### API Documentation
The API documentation can be found in `user-panel/api/user-panel.proto`.

## CI/CD
Coming soon: GitHub Actions for automated testing and deployment.
