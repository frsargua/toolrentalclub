# Backend Architecture - Domain-Driven Design

This document describes the Domain-Driven Design (DDD) architecture of the Tool Rental Club backend API.

## Overview

The backend follows a clean, layered architecture based on Domain-Driven Design principles, ensuring separation of concerns, maintainability, and testability.

## Directory Structure

```
backend/
├── cmd/
│   └── api/
│       └── main.go                 # Application entry point
├── domain/                         # Core business logic (no dependencies)
│   ├── auth/
│   │   ├── service.go             # Auth service interface
│   │   └── token.go               # Token entity
│   └── user/
│       ├── entity.go              # User entity
│       └── repository.go          # User repository interface
├── application/                    # Use cases and application services
│   ├── auth/
│   │   └── service.go             # Auth use cases
│   └── user/
│       └── service.go             # User use cases
├── infrastructure/                 # External dependencies and implementations
│   ├── firebase/
│   │   ├── auth_service.go        # Firebase auth implementation
│   │   └── config.go              # Firebase initialization
│   └── repository/
│       └── memory/
│           └── user_repository.go # In-memory user repository
├── interfaces/                     # HTTP interface layer
│   └── http/
│       ├── dto/                   # Data Transfer Objects
│       │   ├── auth.go
│       │   └── user.go
│       ├── handlers/              # HTTP handlers
│       │   ├── auth_handler.go
│       │   ├── health_handler.go
│       │   ├── response.go
│       │   └── user_handler.go
│       └── middleware/            # HTTP middleware
│           ├── auth.go
│           ├── cors.go
│           └── logging.go
└── pkg/                           # Shared utilities
    ├── config/
    │   └── config.go              # Configuration management
    └── server/
        └── server.go              # HTTP server utilities
```

## Layer Descriptions

### 1. Domain Layer (`domain/`)

The **domain layer** contains the core business logic and has **no external dependencies**. This is the heart of the application.

- **Entities**: Core business objects (e.g., `User`)
- **Repository Interfaces**: Define contracts for data access
- **Service Interfaces**: Define contracts for domain services (e.g., `auth.Service`)

**Rules:**

- No dependencies on other layers
- Pure business logic
- Framework-agnostic

**Examples:**

- `domain/user/entity.go` - User entity with business rules
- `domain/user/repository.go` - Repository interface for user data access
- `domain/auth/service.go` - Authentication service interface

### 2. Application Layer (`application/`)

The **application layer** orchestrates the flow of data and implements use cases. It coordinates domain objects and infrastructure services.

- **Use Cases**: Business operations (e.g., verify token and get/create user)
- **Application Services**: Coordinate multiple domain objects

**Rules:**

- Depends only on the domain layer
- Implements business workflows
- Coordinates repository and domain service calls

**Examples:**

- `application/auth/service.go` - Authentication use cases like token verification
- `application/user/service.go` - User-related use cases

### 3. Infrastructure Layer (`infrastructure/`)

The **infrastructure layer** provides concrete implementations of interfaces defined in the domain layer. This includes external services, databases, and third-party integrations.

- **Repository Implementations**: Concrete data access implementations
- **External Services**: Firebase, APIs, etc.
- **Configuration**: Service initialization

**Rules:**

- Implements domain interfaces
- Contains all external dependencies
- Can be easily swapped or mocked

**Examples:**

- `infrastructure/firebase/auth_service.go` - Firebase implementation of `auth.Service`
- `infrastructure/repository/memory/user_repository.go` - In-memory implementation of `user.Repository`

### 4. Interfaces Layer (`interfaces/`)

The **interfaces layer** handles external communication (HTTP, gRPC, etc.). It translates external requests into application/domain calls.

- **HTTP Handlers**: REST API endpoints
- **DTOs**: Data Transfer Objects for API requests/responses
- **Middleware**: Cross-cutting concerns (auth, logging, CORS)

**Rules:**

- Depends on application layer
- Translates HTTP requests to use case calls
- Handles serialization/deserialization

**Examples:**

- `interfaces/http/handlers/auth_handler.go` - HTTP handlers for authentication
- `interfaces/http/dto/auth.go` - Request/response DTOs
- `interfaces/http/middleware/auth.go` - Authentication middleware

### 5. Package Layer (`pkg/`)

The **pkg layer** contains shared utilities and helpers that can be used across all layers.

- **Configuration**: Application config management
- **Server**: HTTP server utilities
- **Common utilities**: Logging, error handling, etc.

**Examples:**

- `pkg/config/config.go` - Configuration loading
- `pkg/server/server.go` - HTTP server setup

## Dependency Flow

```
cmd/api (main.go)
    ↓
interfaces/http (handlers, middleware)
    ↓
application (use cases)
    ↓
domain (entities, interfaces)
    ↑
infrastructure (implementations)
```

**Key Principle:** Dependencies point inward. The domain layer is at the center and has no external dependencies.

## Key Benefits

1. **Testability**: Each layer can be tested independently
2. **Maintainability**: Clear separation of concerns
3. **Flexibility**: Easy to swap implementations (e.g., change from in-memory to PostgreSQL)
4. **Scalability**: Well-organized structure for growing applications
5. **Domain Focus**: Business logic is isolated and protected

## API Endpoints

### Public Routes

- `GET /api/health` - Health check

### Auth Routes

- `POST /api/auth/verify` - Verify Firebase token

### Protected Routes (require Bearer token)

- `GET /api/profile` - Get authenticated user profile

## Adding New Features

To add a new feature following DDD principles:

1. **Define domain entities and interfaces** in `domain/`
2. **Create application use cases** in `application/`
3. **Implement infrastructure services** in `infrastructure/`
4. **Add HTTP handlers and DTOs** in `interfaces/http/`
5. **Wire everything together** in `cmd/api/main.go`

## Example: Adding a Tool Rental Feature

1. Create `domain/tool/entity.go` and `domain/tool/repository.go`
2. Create `application/tool/service.go` with rental use cases
3. Implement `infrastructure/repository/memory/tool_repository.go`
4. Create `interfaces/http/handlers/tool_handler.go` and `dto/tool.go`
5. Register routes in `main.go`

## Testing Strategy

- **Domain Layer**: Pure unit tests (no mocks needed)
- **Application Layer**: Use case tests with mocked repositories
- **Infrastructure Layer**: Integration tests with real dependencies
- **Interfaces Layer**: HTTP integration tests

## Future Enhancements

- Add database persistence (PostgreSQL, MongoDB, etc.)
- Implement event sourcing for domain events
- Add CQRS pattern for read/write separation
- Implement domain-driven design patterns (Aggregates, Value Objects, etc.)
