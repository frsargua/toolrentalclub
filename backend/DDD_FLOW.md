# Domain-Driven Design Flow

## Request Flow Example: User Profile Request

```
┌─────────────────────────────────────────────────────────────────┐
│                          CLIENT REQUEST                          │
│                  GET /api/profile (with Bearer token)            │
└────────────────────────────┬────────────────────────────────────┘
                             │
                             ▼
┌─────────────────────────────────────────────────────────────────┐
│                     INTERFACES LAYER                             │
│                                                                   │
│  ┌──────────────────────────────────────────────────────────┐   │
│  │  Middleware: CORSMiddleware                              │   │
│  │  └─▶ LoggingMiddleware                                   │   │
│  │      └─▶ AuthMiddleware                                  │   │
│  │          - Extracts token from Authorization header      │   │
│  │          - Calls authUseCase.VerifyToken()               │   │
│  │          - Adds userID & email to context                │   │
│  └──────────────────────────────────────────────────────────┘   │
│                             │                                    │
│                             ▼                                    │
│  ┌──────────────────────────────────────────────────────────┐   │
│  │  Handler: UserHandler.GetProfile()                       │   │
│  │  - Gets userID from context                              │   │
│  │  - Calls userUseCase.GetUserByID(userID)                 │   │
│  │  - Converts User entity to UserProfileResponse DTO       │   │
│  │  - Returns JSON response                                 │   │
│  └──────────────────────────────────────────────────────────┘   │
└────────────────────────────┬────────────────────────────────────┘
                             │
                             ▼
┌─────────────────────────────────────────────────────────────────┐
│                    APPLICATION LAYER                             │
│                                                                   │
│  ┌──────────────────────────────────────────────────────────┐   │
│  │  UseCase: user.UseCase.GetUserByID()                     │   │
│  │  - Orchestrates business logic                           │   │
│  │  - Calls userRepo.FindByID(userID)                       │   │
│  │  - Returns domain User entity                            │   │
│  └──────────────────────────────────────────────────────────┘   │
└────────────────────────────┬────────────────────────────────────┘
                             │
                             ▼
┌─────────────────────────────────────────────────────────────────┐
│                      DOMAIN LAYER                                │
│                                                                   │
│  ┌──────────────────────────────────────────────────────────┐   │
│  │  Repository Interface: user.Repository                   │   │
│  │  - Defines contract: FindByID(id) → *User                │   │
│  │  - Framework-agnostic                                    │   │
│  │  - Pure business logic                                   │   │
│  └──────────────────────────────────────────────────────────┘   │
│                                                                   │
│  ┌──────────────────────────────────────────────────────────┐   │
│  │  Entity: User                                            │   │
│  │  - ID, Email, CreatedAt, UpdatedAt                       │   │
│  └──────────────────────────────────────────────────────────┘   │
└────────────────────────────┬────────────────────────────────────┘
                             │
                             ▼
┌─────────────────────────────────────────────────────────────────┐
│                  INFRASTRUCTURE LAYER                            │
│                                                                   │
│  ┌──────────────────────────────────────────────────────────┐   │
│  │  Implementation: memory.UserRepository                   │   │
│  │  - Implements user.Repository interface                  │   │
│  │  - In-memory storage with sync.RWMutex                   │   │
│  │  - Can be swapped with PostgreSQL, MongoDB, etc.         │   │
│  └──────────────────────────────────────────────────────────┘   │
└────────────────────────────┬────────────────────────────────────┘
                             │
                             ▼
                    Returns User entity
                             │
                             ▼
            [Flows back up through layers to client]
```

## Authentication Flow Example: Token Verification

```
┌─────────────────────────────────────────────────────────────────┐
│                          CLIENT REQUEST                          │
│              POST /api/auth/verify { "token": "..." }            │
└────────────────────────────┬────────────────────────────────────┘
                             │
                             ▼
┌─────────────────────────────────────────────────────────────────┐
│                     INTERFACES LAYER                             │
│  ┌──────────────────────────────────────────────────────────┐   │
│  │  Handler: AuthHandler.VerifyToken()                      │   │
│  │  - Parses VerifyTokenRequest DTO                         │   │
│  │  - Calls authUseCase.VerifyTokenAndGetUser()             │   │
│  │  - Returns VerifyTokenResponse DTO                       │   │
│  └──────────────────────────────────────────────────────────┘   │
└────────────────────────────┬────────────────────────────────────┘
                             │
                             ▼
┌─────────────────────────────────────────────────────────────────┐
│                    APPLICATION LAYER                             │
│  ┌──────────────────────────────────────────────────────────┐   │
│  │  UseCase: auth.UseCase.VerifyTokenAndGetUser()          │   │
│  │  1. Calls authService.VerifyToken()                      │   │
│  │  2. Calls userRepo.FindByID(token.UserID)                │   │
│  │  3. If user doesn't exist, creates new user              │   │
│  │  4. Returns token and user                               │   │
│  └──────────────────────────────────────────────────────────┘   │
└──────────────┬─────────────────────────────┬────────────────────┘
               │                             │
               ▼                             ▼
┌──────────────────────────┐  ┌─────────────────────────────────┐
│     DOMAIN LAYER         │  │     DOMAIN LAYER                │
│  auth.Service Interface  │  │  user.Repository Interface      │
└──────────────┬───────────┘  └─────────────┬───────────────────┘
               │                             │
               ▼                             ▼
┌──────────────────────────┐  ┌─────────────────────────────────┐
│  INFRASTRUCTURE LAYER    │  │  INFRASTRUCTURE LAYER           │
│  firebase.AuthService    │  │  memory.UserRepository          │
│  - Verifies with         │  │  - Finds or creates user        │
│    Firebase SDK          │  │                                 │
│  - Returns auth.Token    │  │                                 │
└──────────────────────────┘  └─────────────────────────────────┘
```

## Dependency Inversion Principle

```
┌─────────────────────────────────────────────────────────────────┐
│                        HIGH-LEVEL POLICY                         │
│                                                                   │
│                    APPLICATION LAYER                             │
│         (Depends on Domain Interfaces)                           │
│                          ▲  ▲                                    │
│                          │  │                                    │
│                          │  │ depends on                         │
│                          │  │                                    │
└──────────────────────────┼──┼───────────────────────────────────┘
                           │  │
┌──────────────────────────┼──┼───────────────────────────────────┐
│                          │  │                                    │
│                    DOMAIN LAYER                                  │
│            (Defines Interfaces - No Dependencies)                │
│                                                                   │
│    ┌─────────────────┐         ┌──────────────────┐            │
│    │ user.Repository │         │  auth.Service    │            │
│    │   (interface)   │         │   (interface)    │            │
│    └─────────────────┘         └──────────────────┘            │
│              ▲                          ▲                        │
│              │                          │                        │
│              │ implements               │ implements             │
│              │                          │                        │
└──────────────┼──────────────────────────┼────────────────────────┘
               │                          │
┌──────────────┼──────────────────────────┼────────────────────────┐
│              │                          │                        │
│        INFRASTRUCTURE LAYER                                      │
│    (Implements Domain Interfaces)                                │
│                                                                   │
│    ┌───────────────────┐      ┌────────────────────┐           │
│    │ memory.UserRepo   │      │ firebase.AuthSvc   │           │
│    │ (implementation)  │      │  (implementation)  │           │
│    └───────────────────┘      └────────────────────┘           │
│                                                                   │
│                        LOW-LEVEL DETAILS                         │
└─────────────────────────────────────────────────────────────────┘
```

## Key Benefits Illustrated

### 1. Testability

```
Test Application Layer:
  authUseCase := auth.NewUseCase(mockAuthService, mockUserRepo)
  ✓ No need for real Firebase
  ✓ No need for real database

Test Domain Layer:
  user := user.NewUser("123", "test@example.com")
  ✓ Pure unit tests
  ✓ No external dependencies
```

### 2. Flexibility

```
Want to switch from Memory to PostgreSQL?

  Before:
    userRepo := memory.NewUserRepository()

  After:
    userRepo := postgres.NewUserRepository(db)

  ✓ Only change one line in main.go
  ✓ Domain and Application layers unchanged
  ✓ All business logic preserved
```

### 3. Maintainability

```
Each layer has a clear responsibility:

  Domain:        What is a User? What can it do?
  Application:   How do we verify a user?
  Infrastructure: Where do we store users?
  Interfaces:    How do clients interact with users?
```

## Common Patterns in DDD

### Repository Pattern

```
Domain defines what:
  type Repository interface {
      FindByID(ctx, id) (*User, error)
  }

Infrastructure defines how:
  type PostgresRepository struct {
      db *sql.DB
  }

  func (r *PostgresRepository) FindByID(...) (*User, error) {
      // SQL query implementation
  }
```

### Use Case Pattern

```
Application coordinates:
  func (uc *UseCase) VerifyTokenAndGetUser(...) {
      token := authService.VerifyToken()  // Infrastructure
      user := userRepo.FindByID()         // Infrastructure
      // Business logic here
      return token, user
  }
```

### DTO Pattern

```
Interfaces translate:
  Domain Entity (internal):
    type User struct {
        ID, Email string
        CreatedAt time.Time
    }

  DTO (external API):
    type UserProfileResponse struct {
        UserID  string `json:"userId"`
        Email   string `json:"email"`
    }
```
