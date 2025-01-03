# 🚀 Qolibaba API Project

## 🌟 Overview

Qolibaba API is a microservices-based backend solution, built with Go, that offers an extensible framework for handling user, admin, and route mapping services. The project is designed for high performance, scalability, and maintainability, leveraging gRPC, REST, and Docker.

---

## ✨ Features

- 🔑 **Admin Management**: APIs and gRPC services for administrative operations.
- 👥 **User Management**: APIs to manage user accounts and related functionalities.
- 🗺️ **Route Mapping**: Service for managing and querying route information.
- 🛢️ **PostgreSQL Integration**: Robust database support using GORM.
- 🔒 **JWT Authentication**: Secure API access with token-based authentication.
- 🐳 **Dockerized Deployment**: Easy setup with `docker-compose`.
- ⚙️ **Configurable Architecture**: Customizable environment and service parameters.

---

## 📂 Project Structure

```plaintext
api/
├── handlers/         # HTTP and gRPC handlers
│   ├── grpc/         # gRPC services
│   └── http/         # RESTful services
├── pb/               # Protobuf files and generated code
└── service/          # Core service logic

app/
├── admin/            # Admin-specific application logic
├── routemap/         # RouteMap-specific logic

config/               # Configuration-related utilities
internal/             # Internal domain logic
pkg/                  # Shared libraries and utilities
build/                # Dockerfile configurations
cmd/                  # Entry points for different services
```

---

## 🛠️ Installation

### Prerequisites

- 🐹 [Go](https://golang.org/) 1.23+
- 🐳 [Docker](https://www.docker.com/)
- 🧩 [Docker Compose](https://docs.docker.com/compose/)

### 🛒 Clone the Repository

```bash
git clone https://github.com/qolibaba-group6/qolibaba.git
cd qolibaba
```

### 📝 Environment Configuration

Copy the sample environment file and adjust parameters:

```bash
cp .env-sample .env
```

Update `.env` with your preferred values.

### 🗄️ Database Configuration

The project uses PostgreSQL. Edit `config.json` and the `.env` file to match your setup.

Sample configuration:

```json
{
    "db": {
        "host": "qolibaba-db",
        "port": 5432,
        "database": "qolibaba_db",
        "user": "app",
        "password": "password123",
        "schema": "public"
    }
}
```

---

## ▶️ Running the Application

### 🐳 With Docker Compose

Ensure Docker is running, then execute:

```bash
docker-compose up --build
```

### 🖥️ Without Docker

1. Start PostgreSQL and ensure it matches your `config.json` setup.
2. Build the Go binaries:

   ```bash
   go mod download
   go build ./cmd/main.go
   go build -o ./service-name ./cmd/service-name/main.go
   ```

3. Run the services:

   ```bash
   ./main
   ```
   run other services in new terminals:
    ```bash
    ./service-name
    ```

---

## 📖 API Documentation

### 🌐 REST API Endpoints

- **Admin**: `/api/v1/admin`
- **User**: `/api/v1/user`
- **RouteMap**: `/api/v1/routemap`

### 🔗 gRPC Services

Refer to the `.proto` files in the `api/pb` directory for available methods and message definitions.

---

## 👩‍💻 Development

### 🔨 Generating Protobuf Files

To regenerate gRPC and protobuf code:

```bash
make gen-proto
```

### ✅ Testing

Run unit tests with:

```bash
go test ./...
```

---

## 🚢 Deployment

Adjust the `docker-compose.yaml` file for production settings, such as:

- 🔧 Updating `HTTP_PORT`, `DB_PORT`, and `ROUTEMAP_PORT`.
- 🔒 Adding secure environment variables.

Deploy using:

```bash
docker-compose -f docker-compose.yaml up -d
```

---

## 🤝 Contributing

Feel free to submit issues, fork the repository, and create pull requests. Ensure your code passes all lint and test checks before submission.

---

## 📜 License

This project is licensed under the MIT License. See the `LICENSE` file for details.

---
