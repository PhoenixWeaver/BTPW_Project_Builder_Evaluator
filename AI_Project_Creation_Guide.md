# 🤖 AI Project Creation Guide - Step-by-Step Instructions

## 📋 Overview
This guide provides step-by-step instructions for creating new Go projects using the AI advice functions and existing project analysis. It combines the AI advisor diagrams (`AIAd_diagrams.go`) with the existing project analysis (`Existing_diagrams.go`) to provide comprehensive guidance.

## 🎯 Prerequisites
- Go installed (version 1.19+)
- Docker installed (for database)
- PostgreSQL (via Docker)
- Basic understanding of Go, HTTP, and databases

## 📊 Available Analysis Tools

### 1. AI Advisor Functions (`AIAd_diagrams.go`)
- **Purpose**: Provides AI-generated guidance for project structure
- **Functions**: 
  - `AIAd_WriteAllStructureDiagrams()` - Generates AI advice diagrams
  - `AIAd_WriteFunctionFlowAnalysis()` - Function flow analysis
  - `AIAd_WriteDevelopmentSequence()` - Development sequence guidance

### 2. Existing Project Analysis (`Existing_diagrams.go`)
- **Purpose**: Analyzes your current project state
- **Functions**:
  - `Existing_scanProject()` - Scans project for functions
  - `Existing_WriteFunctionDependencyDiagram()` - Function dependencies
  - `Existing_generateUpdatedReports()` - Current project reports

### 3. Function Order Diagrams (`BTFlowcharts_functions.go`)
- **Purpose**: Shows creation and execution order
- **Functions**:
  - `BTFlowcharts_WriteFunctionCreationOrder()` - Development sequence
  - `BTFlowcharts_WriteFunctionExecutionOrder()` - Runtime sequence

## 🚀 Step-by-Step Project Creation Process

### Phase 1: Project Foundation (42m 33s)
**Goal**: Set up basic project structure and HTTP server

#### Step 1: Initialize Project
```bash
# Create project directory
mkdir my-go-api
cd my-go-api

# Initialize Go module
go mod init my-go-api

# Create basic structure
mkdir -p internal/{app,api,store,middleware}
mkdir -p migrations
mkdir -p cmd
```

#### Step 2: Create Main Application
```go
// cmd/main.go
package main

import (
    "flag"
    "log"
    "net/http"
    "time"
)

func main() {
    port := flag.String("port", "8080", "Server port")
    flag.Parse()

    server := &http.Server{
        Addr:         ":" + *port,
        ReadTimeout:  10 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  time.Minute,
    }

    log.Printf("Server starting on port %s", *port)
    log.Fatal(server.ListenAndServe())
}
```

#### Step 3: Add Chi Router
```bash
go get github.com/go-chi/chi/v5
```

```go
// internal/app/app.go
package app

import (
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

type Application struct {
    Router *chi.Mux
}

func NewApplication() *Application {
    r := chi.NewRouter()
    
    // Middleware
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    
    // Routes
    r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("OK"))
    })
    
    return &Application{Router: r}
}
```

### Phase 2: Data Layer (1h 35s)
**Goal**: Set up database connection and data models

#### Step 4: Docker Database Setup
```yaml
# docker-compose.yml
version: '3.8'
services:
  postgres:
    image: postgres:15
    environment:
      POSTGRES_DB: myapp
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data

volumes:
  postgres_data:
```

```bash
# Start database
docker-compose up -d
```

#### Step 5: Database Connection
```bash
go get github.com/jackc/pgx/v5
go get github.com/jackc/pgx/v5/pgxpool
```

```go
// internal/database/database.go
package database

import (
    "context"
    "fmt"
    "github.com/jackc/pgx/v5/pgxpool"
)

func OpenDatabase(dsn string) (*pgxpool.Pool, error) {
    config, err := pgxpool.ParseConfig(dsn)
    if err != nil {
        return nil, fmt.Errorf("failed to parse config: %w", err)
    }
    
    pool, err := pgxpool.NewWithConfig(context.Background(), config)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to database: %w", err)
    }
    
    return pool, nil
}
```

#### Step 6: Migrations with Goose
```bash
go get github.com/pressly/goose/v3
```

```go
// internal/database/migrate.go
package database

import (
    "database/sql"
    "github.com/pressly/goose/v3"
    _ "github.com/jackc/pgx/v5/stdlib"
)

func Migrate(db *sql.DB, dir string) error {
    if err := goose.SetDialect("postgres"); err != nil {
        return err
    }
    
    return goose.Up(db, dir)
}
```

#### Step 7: Create First Migration
```sql
-- migrations/00001_create_items.sql
CREATE TABLE IF NOT EXISTS items (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Phase 3: CRUD Operations (1h 24m 15s)
**Goal**: Implement complete CRUD functionality

#### Step 8: Data Models
```go
// internal/store/models.go
package store

import "time"

type Item struct {
    ID          int       `json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
```

#### Step 9: Store Interface
```go
// internal/store/store.go
package store

import "context"

type Store interface {
    CreateItem(ctx context.Context, item *Item) error
    GetItemByID(ctx context.Context, id int) (*Item, error)
    UpdateItem(ctx context.Context, item *Item) error
    DeleteItem(ctx context.Context, id int) error
    ListItems(ctx context.Context) ([]*Item, error)
}
```

#### Step 10: Store Implementation
```go
// internal/store/postgres.go
package store

import (
    "context"
    "github.com/jackc/pgx/v5/pgxpool"
)

type PostgresStore struct {
    db *pgxpool.Pool
}

func NewPostgresStore(db *pgxpool.Pool) Store {
    return &PostgresStore{db: db}
}

func (s *PostgresStore) CreateItem(ctx context.Context, item *Item) error {
    query := `INSERT INTO items (name, description) VALUES ($1, $2) RETURNING id, created_at, updated_at`
    return s.db.QueryRow(ctx, query, item.Name, item.Description).Scan(&item.ID, &item.CreatedAt, &item.UpdatedAt)
}

func (s *PostgresStore) GetItemByID(ctx context.Context, id int) (*Item, error) {
    item := &Item{}
    query := `SELECT id, name, description, created_at, updated_at FROM items WHERE id = $1`
    err := s.db.QueryRow(ctx, query, id).Scan(&item.ID, &item.Name, &item.Description, &item.CreatedAt, &item.UpdatedAt)
    if err != nil {
        return nil, err
    }
    return item, nil
}

// ... implement other methods
```

#### Step 11: API Handlers
```go
// internal/api/handlers.go
package api

import (
    "encoding/json"
    "net/http"
    "strconv"
    "github.com/go-chi/chi/v5"
    "my-go-api/internal/store"
)

type Handler struct {
    store store.Store
}

func NewHandler(store store.Store) *Handler {
    return &Handler{store: store}
}

func (h *Handler) CreateItem(w http.ResponseWriter, r *http.Request) {
    var item store.Item
    if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }
    
    if err := h.store.CreateItem(r.Context(), &item); err != nil {
        http.Error(w, "Failed to create item", http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(item)
}

func (h *Handler) GetItem(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }
    
    item, err := h.store.GetItemByID(r.Context(), id)
    if err != nil {
        http.Error(w, "Item not found", http.StatusNotFound)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(item)
}

// ... implement other handlers
```

### Phase 4: Testing (38m 20s)
**Goal**: Implement comprehensive testing

#### Step 12: Test Setup
```go
// internal/store/store_test.go
package store

import (
    "context"
    "testing"
    "github.com/jackc/pgx/v5/pgxpool"
)

func setupTestDB(t *testing.T) *pgxpool.Pool {
    // Setup test database connection
    // Use separate test database
}

func TestCreateItem(t *testing.T) {
    db := setupTestDB(t)
    store := NewPostgresStore(db)
    
    item := &Item{
        Name:        "Test Item",
        Description: "Test Description",
    }
    
    err := store.CreateItem(context.Background(), item)
    if err != nil {
        t.Fatalf("Failed to create item: %v", err)
    }
    
    if item.ID == 0 {
        t.Error("Expected ID to be set")
    }
}
```

### Phase 5: Authentication (1h 20m 4s)
**Goal**: Add user authentication and authorization

#### Step 13: User Model
```go
// internal/store/user.go
package store

import (
    "time"
    "golang.org/x/crypto/bcrypt"
)

type User struct {
    ID        int       `json:"id"`
    Email     string    `json:"email"`
    Password  string    `json:"-"` // Don't include in JSON
    CreatedAt time.Time `json:"created_at"`
}

func (u *User) HashPassword() error {
    hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    u.Password = string(hashed)
    return nil
}

func (u *User) CheckPassword(password string) bool {
    return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}
```

#### Step 14: JWT Tokens
```bash
go get github.com/golang-jwt/jwt/v5
```

```go
// internal/auth/jwt.go
package auth

import (
    "time"
    "github.com/golang-jwt/jwt/v5"
)

type Claims struct {
    UserID int `json:"user_id"`
    jwt.RegisteredClaims
}

func GenerateToken(userID int, secret string) (string, error) {
    claims := Claims{
        UserID: userID,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(secret))
}
```

### Phase 6: Middleware (58m 44s)
**Goal**: Add authentication middleware and route protection

#### Step 15: Authentication Middleware
```go
// internal/middleware/auth.go
package middleware

import (
    "net/http"
    "strings"
    "github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(secret string) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            authHeader := r.Header.Get("Authorization")
            if authHeader == "" {
                http.Error(w, "Authorization header required", http.StatusUnauthorized)
                return
            }
            
            tokenString := strings.TrimPrefix(authHeader, "Bearer ")
            token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
                return []byte(secret), nil
            })
            
            if err != nil || !token.Valid {
                http.Error(w, "Invalid token", http.StatusUnauthorized)
                return
            }
            
            // Add user ID to context
            claims := token.Claims.(jwt.MapClaims)
            userID := int(claims["user_id"].(float64))
            ctx := context.WithValue(r.Context(), "user_id", userID)
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}
```

## 🔧 Using the Analysis Tools

### Generate AI Advice
```bash
# Run the interactive mode
go run -tags flowcharts BTProjectDiagrams.go -interactive -out BTFlowcharts

# Choose option 4: AI Advisor Diagrams
```

### Analyze Current Project
```bash
# Choose option 8: Existing Diagrams
# This will scan your current project and generate:
# - Function inventory
# - Development sequence
# - Project status report
# - Function dependencies
```

### Generate Function Order Diagrams
```bash
# Choose option 11: Function Creation & Execution Order Diagrams
# This will generate:
# - Function creation order (development sequence)
# - Function execution order (runtime sequence)
```

## 📊 Understanding the Output

### AI Advisor Diagrams
- **Development Sequence**: Shows optimal order to build functions
- **Execution Flow**: Shows how functions interact at runtime
- **Function Dependencies**: Shows which functions depend on others

### Existing Project Analysis
- **Function Inventory**: Complete list of all functions in your project
- **Dynamic Development Sequence**: Based on your actual project structure
- **Project Status Report**: Statistics about your current project

### Function Order Diagrams
- **Creation Order**: Shows the sequence functions were/will be created
- **Execution Order**: Shows the sequence functions execute at runtime

## 🎯 Best Practices

1. **Start with Foundation**: Always begin with database and basic server setup
2. **Build Incrementally**: Add one feature at a time and test thoroughly
3. **Use the Analysis Tools**: Regularly generate diagrams to understand your project
4. **Follow the Phases**: Don't skip phases - each builds on the previous
5. **Test Early and Often**: Write tests as you build, not after
6. **Document as You Go**: Use the generated diagrams as documentation

## 🚨 Common Issues and Solutions

### HTML Resolution Issues
- **Problem**: Charts not visible in HTML
- **Solution**: Use the fixed `Existing_WriteFunctionDependencyDiagram` function
- **Prevention**: Avoid complex HTML styling in Mermaid diagrams

### Function Dependencies
- **Problem**: Functions calling each other in wrong order
- **Solution**: Use the function order diagrams to understand dependencies
- **Prevention**: Build functions in the correct sequence

### Database Connection Issues
- **Problem**: Can't connect to database
- **Solution**: Ensure Docker is running and database is accessible
- **Prevention**: Test database connection early in development

## 📈 Next Steps

1. **Run the Analysis**: Generate all diagrams to understand your project
2. **Follow the Guide**: Use this step-by-step process for new projects
3. **Customize**: Adapt the examples to your specific use case
4. **Iterate**: Use the analysis tools to continuously improve your project structure

## 🔗 Related Files

- `AIAd_diagrams.go` - AI advisor functions
- `Existing_diagrams.go` - Project analysis functions
- `BTFlowcharts_functions.go` - Function order diagrams
- `BTProjectDiagrams.go` - Main orchestrator
- `IntructorProjectBuilderModel_Go1.txt` - Instructor's teaching progression

---

**Remember**: This guide combines AI advice with real project analysis to provide the most comprehensive development guidance possible. Use the generated diagrams to understand both the ideal structure and your current project state.
