# Tool Rental Club - Backend

A Golang REST API backend for Tool Rental Club with Firebase authentication integration.

## Features

- 🔐 Firebase Authentication integration
- 🚀 RESTful API with Gorilla Mux
- 🔒 JWT token verification middleware
- 📝 Request logging
- 🌐 CORS support
- 🏗️ Domain-Driven Design (DDD) architecture
- 🧩 Clean separation of concerns with layered architecture
- 🔄 Easy to extend and maintain

## Prerequisites

- Go 1.21 or higher
- Firebase project with Admin SDK credentials

## Setup

### 1. Install Dependencies

```bash
go mod download
```

### 2. Firebase Admin SDK Setup

1. Go to the [Firebase Console](https://console.firebase.google.com/)
2. Select your project
3. Go to Project Settings > Service Accounts
4. Click "Generate New Private Key"
5. Save the JSON file as `serviceAccountKey.json` in the backend directory (or any secure location)
6. Update the `.env` file with the path to your service account key

### 3. Environment Configuration

Copy `.env.example` to `.env`:

```bash
cp .env.example .env
```

Update the values in `.env`:

```env
PORT=8080
FIREBASE_SERVICE_ACCOUNT=./serviceAccountKey.json
```

### 4. Run the Server

Development mode:

```bash
go run cmd/api/main.go
```

Or build and run:

```bash
go build -o bin/server cmd/api/main.go
./bin/server
```

The server will start on `http://localhost:8080`

## API Endpoints

### Public Endpoints

- `GET /api/health` - Health check endpoint

### Authentication Endpoints

- `POST /api/auth/verify` - Verify Firebase ID token

  ```json
  Request:
  {
    "token": "firebase-id-token"
  }

  Response:
  {
    "success": true,
    "message": "Token verified successfully",
    "userId": "user-id",
    "email": "user@example.com"
  }
  ```

### Protected Endpoints

These endpoints require authentication (Bearer token in Authorization header):

- `GET /api/profile` - Get user profile

  ```
  Headers:
  Authorization: Bearer <firebase-id-token>

  Response:
  {
    "userId": "user-id",
    "email": "user@example.com",
    "message": "This is a protected route"
  }
  ```

## Project Structure

This backend follows **Domain-Driven Design (DDD)** principles with a clean, layered architecture:

```
backend/
├── cmd/
│   └── api/
│       └── main.go                    # Application entry point
├── domain/                            # Core business logic (no dependencies)
│   ├── auth/
│   │   ├── service.go                # Auth service interface
│   │   └── token.go                  # Token entity
│   └── user/
│       ├── entity.go                 # User entity
│       └── repository.go             # User repository interface
├── application/                       # Use cases and application services
│   ├── auth/
│   │   └── service.go                # Auth use cases
│   └── user/
│       └── service.go                # User use cases
├── infrastructure/                    # External dependencies and implementations
│   ├── firebase/
│   │   ├── auth_service.go           # Firebase auth implementation
│   │   └── config.go                 # Firebase initialization
│   └── repository/
│       └── memory/
│           └── user_repository.go    # In-memory user repository
├── interfaces/                        # HTTP interface layer
│   └── http/
│       ├── dto/                      # Data Transfer Objects
│       │   ├── auth.go
│       │   └── user.go
│       ├── handlers/                 # HTTP handlers
│       │   ├── auth_handler.go
│       │   ├── health_handler.go
│       │   ├── response.go
│       │   └── user_handler.go
│       └── middleware/               # HTTP middleware
│           ├── auth.go
│           ├── cors.go
│           └── logging.go
├── pkg/                              # Shared utilities
│   ├── config/
│   │   └── config.go                # Configuration management
│   └── server/
│       └── server.go                # HTTP server utilities
├── ARCHITECTURE.md                   # Detailed architecture documentation
├── go.mod                           # Go module definition
└── README.md                        # This file
```

**See [ARCHITECTURE.md](./ARCHITECTURE.md) for detailed documentation on the DDD structure.**

## Authentication Flow

1. User logs in on the frontend using Firebase Authentication
2. Frontend receives a Firebase ID token
3. Frontend sends the token to `/api/auth/verify`
4. Backend verifies the token with Firebase Admin SDK
5. Backend creates/updates user session (implement as needed)
6. For protected routes, frontend includes the token in Authorization header
7. Auth middleware verifies the token on each request

## Adding New Features

Following DDD principles, here's how to add a new feature (e.g., Tool Rental):

### 1. Define Domain Layer

```go
// domain/tool/entity.go
type Tool struct {
    ID          string
    Name        string
    Description string
    OwnerID     string
}

// domain/tool/repository.go
type Repository interface {
    FindByID(ctx context.Context, id string) (*Tool, error)
    Create(ctx context.Context, tool *Tool) error
}
```

### 2. Create Application Use Cases

```go
// application/tool/service.go
type UseCase struct {
    toolRepo tool.Repository
}

func (uc *UseCase) CreateTool(ctx context.Context, tool *tool.Tool) error {
    return uc.toolRepo.Create(ctx, tool)
}
```

### 3. Implement Infrastructure

```go
// infrastructure/repository/memory/tool_repository.go
type ToolRepository struct {
    // implementation
}
```

### 4. Add HTTP Interface

```go
// interfaces/http/handlers/tool_handler.go
type ToolHandler struct {
    toolUseCase *tool.UseCase
}

// interfaces/http/dto/tool.go
type CreateToolRequest struct {
    Name        string `json:"name"`
    Description string `json:"description"`
}
```

### 5. Wire in main.go

```go
// cmd/api/main.go
toolRepo := memory.NewToolRepository()
toolUseCase := toolApp.NewUseCase(toolRepo)
toolHandler := handlers.NewToolHandler(toolUseCase)

protectedRouter.HandleFunc("/tools", toolHandler.CreateTool).Methods("POST")
```

## Database Integration

Currently using an **in-memory repository** for users (see `infrastructure/repository/memory/`).

To add a real database:

### Option 1: PostgreSQL

1. Create `infrastructure/repository/postgres/user_repository.go`
2. Implement the `user.Repository` interface
3. Wire it in `main.go` instead of memory repository

### Option 2: Firestore

1. Create `infrastructure/repository/firestore/user_repository.go`
2. Use the existing Firebase app connection
3. Implement the `user.Repository` interface

### Option 3: MongoDB

1. Create `infrastructure/repository/mongo/user_repository.go`
2. Add MongoDB client initialization
3. Implement the `user.Repository` interface

**The beauty of DDD:** You can swap repository implementations without changing domain or application layers!

## Security Considerations

1. **CORS**: Update `CORSMiddleware` in production to only allow your frontend domain
2. **Environment Variables**: Never commit `.env` or service account keys
3. **HTTPS**: Use HTTPS in production (consider a reverse proxy like nginx)
4. **Rate Limiting**: Add rate limiting middleware to prevent abuse
5. **Input Validation**: Always validate and sanitize user input

## Testing

Create test files alongside your code:

```bash
go test ./...
```

## Deployment

### Option 1: Traditional Server

```bash
# Build
go build -o server cmd/api/main.go

# Run with environment variables
export PORT=8080
export FIREBASE_SERVICE_ACCOUNT=/path/to/key.json
./server
```

### Option 2: Docker

Create a `Dockerfile`:

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN go build -o server cmd/api/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]
```

### Option 3: Cloud Platforms

- **Google Cloud Run**: Native support for Go and Firebase
- **AWS Lambda**: Use with API Gateway
- **Heroku**: Deploy with a `Procfile`
- **DigitalOcean App Platform**: Simple Go deployment

## License

MIT
