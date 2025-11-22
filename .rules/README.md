# Project Rules & Documentation

This directory contains comprehensive documentation for the backend project architecture, patterns, and coding standards.

## 📚 Documentation Index

### 0. [Architecture Overview](./000-architecture-overview.md) ⭐ START HERE

- **Purpose:** High-level architecture explanation and core principles
- **Topics:**
  - Layer responsibilities
  - Service layer: NO repos, NO adapters (pure logic)
  - Application layer: Multiple repos + adapters
  - Data flow examples
  - Testing strategy per layer
  - Key differences between layers

### 1. [Project Structure](./001-structure.md)

- **Purpose:** Overview of project organization and directory layout
- **Topics:**
  - Technology stack
  - Directory structure
  - File organization
  - Mock file locations

### 2. [Domain Layer](./002-domain-layer.md)

- **Purpose:** Domain-Driven Design patterns and conventions
- **Topics:**
  - Domain models with validation tags
  - Repository interfaces (List, Count, Get, Save)
  - Service interfaces (business logic)
  - Domain errors
  - Enums and common types
  - Mock generation with mockery

### 3. [Application Layer](./003-application-layer.md)

- **Purpose:** Use case orchestration and workflow patterns
- **Topics:**
  - Application interfaces and implementations
  - Parameter objects (DTOs)
  - Pagination patterns
  - Dependency injection with Google Wire
  - Error propagation
  - Workflow orchestration

### 4. [Service Layer](./004-service-layer.md)

- **Purpose:** Business logic and entity factory methods
- **Topics:**
  - Domain service implementation
  - Factory methods (Create, CreateVariant, etc.)
  - Validation with go-playground/validator
  - Error wrapping with multierror
  - UUID generation patterns
  - Business rule enforcement

### 5. [Repository Layer](./005-repository-layer.md)

- **Purpose:** Data persistence with PostgreSQL and sqlc
- **Topics:**
  - Repository implementation patterns
  - sqlc integration
  - Upsert operations (ON CONFLICT DO UPDATE)
  - Error mapping (PostgreSQL → domain errors)
  - Type mapping (sqlc → domain)
  - Transaction handling

### 6. [Testing](./006-testing.md)

- **Purpose:** Testing strategies and patterns
- **Topics:**
  - Mock generation with mockery
  - Unit tests (service, application layers)
  - Integration tests (repositories)
  - Test containers (PostgreSQL, Redis, MinIO, Keycloak)
  - Table-driven tests
  - Test helpers and utilities
  - CI/CD integration

### 7. [Coding Standards](./007-coding-standards.md)

- **Purpose:** Go style guide and best practices
- **Topics:**
  - Naming conventions
  - Code organization
  - Error handling patterns
  - Pointer vs value usage
  - Context usage
  - JSON tags and formatting
  - Performance tips
  - Linting with golangci-lint

### 8. [Database Conventions](./008-database-conventions.md)

- **Purpose:** Database schema design and SQL patterns
- **Topics:**
  - Schema design (tables, columns, types)
  - Indexes and constraints
  - Soft delete patterns
  - Migration management
  - sqlc query patterns
  - Triggers
  - Seeding data
  - Transaction handling

### 100. [Business Logic Rules](./100-business-logic.md)
- **Purpose:** Domain-specific business rules and workflows
- **Topics:**
  - Soft delete strategy (CartItem is hard delete exception)
  - Aggregate definitions and relationships
  - Product structure (variants, options, attributes)
  - Cart management (one cart per user, hard delete items)
  - Order workflow and inventory management
  - Order confirmation transaction (decrease inventory, clear cart)
  - Review and rating calculation
  - Validation rules per aggregate
  - Event-driven side effects
  - Concurrency handling and data consistency

## 🚀 Quick Start

If you're new to the project:

1. **Start with:** [000-architecture-overview.md](./000-architecture-overview.md) for core architecture concepts ⭐
2. **Then read:** [100-business-logic.md](./100-business-logic.md) for domain rules and workflows ⭐
3. **Understand:** [001-structure.md](./001-structure.md) for project organization
4. **Learn DDD:** [002-domain-layer.md](./002-domain-layer.md) for DDD fundamentals
5. **Services:** [004-service-layer.md](./004-service-layer.md) - services have NO repos!
6. **Application:** [003-application-layer.md](./003-application-layer.md) - apps inject multiple repos + adapters
7. **Reference:** Other documents as needed for specific tasks

## 🎯 Common Tasks

### Adding a New Entity

1. **Domain:** Define model, repository interface, service interface ([002-domain-layer.md](./002-domain-layer.md))
2. **Service:** Implement domain service with factory methods ([004-service-layer.md](./004-service-layer.md))
3. **Database:** Create migration, add sqlc queries ([008-database-conventions.md](./008-database-conventions.md))
4. **Repository:** Implement repository with sqlc ([005-repository-layer.md](./005-repository-layer.md))
5. **Application:** Create use cases and parameter objects ([003-application-layer.md](./003-application-layer.md))
6. **Testing:** Add unit and integration tests ([006-testing.md](./006-testing.md))
7. **Wire:** Add to dependency injection ([003-application-layer.md](./003-application-layer.md#dependency-injection))

### Writing Tests

1. **Generate mocks:** Run `mockery` ([006-testing.md](./006-testing.md#mock-generation))
2. **Unit tests:** Test services with validator ([006-testing.md](./006-testing.md#service-layer-tests))
3. **Integration tests:** Test repositories with test containers ([006-testing.md](./006-testing.md#integration-tests))
4. **Run tests:** `go test ./...` ([006-testing.md](./006-testing.md#running-tests))

### Adding Database Queries

1. **Create migration:** `migration/YYYYMMDDHHMMSS.sql` ([008-database-conventions.md](./008-database-conventions.md#migrations))
2. **Update schema:** `database/schema.sql` ([008-database-conventions.md](./008-database-conventions.md#schema-management))
3. **Write sqlc query:** `database/queries/<entity>.sql` ([008-database-conventions.md](./008-database-conventions.md#sqlc-queries))
4. **Generate code:** `sqlc generate` ([005-repository-layer.md](./005-repository-layer.md#sqlc-configuration))
5. **Use in repository:** Map to domain errors ([005-repository-layer.md](./005-repository-layer.md#error-mapping))

## 📖 Conventions at a Glance

### Architecture Layers (Top → Down)

```
Delivery (HTTP/gRPC)
        ↓
Application (Use Cases)
  - Injects multiple repositories
  - Injects external adapters (S3, Redis)
  - Orchestrates workflows
        ↓
Domain (Business Logic) ← Service (Pure Logic)
  - Interfaces only           - NO repositories
                               - NO adapters
                               - Only validator
        ↓
Repository (Data Access)
  - Implements domain interfaces
  - Maps errors
        ↓
Database (PostgreSQL)
```

### Error Flow (Bottom → Up)

```
PostgreSQL Error
        ↓
Repository: Map to Domain Error
        ↓
Service: Wrap with multierror
        ↓
Application: Propagate
        ↓
Delivery: Map to HTTP status
```

### Key Patterns

- **Repository:** List, Count, Get, Save (minimal CRUD)
- **Service:** Factory methods, validation, business logic (NO repos, NO adapters)
- **Application:** Orchestration, inject multiple repos + adapters, fetch deps → service → persist
- **Domain Errors:** Defined once, wrapped everywhere
- **Validation:** go-playground/validator in services only
- **DI:** Google Wire with provider functions
- **Testing:** mockery + testify + test containers + table-driven tests
- **Database:** sqlc + PostgreSQL + soft deletes

## 🛠️ Tools & Commands

**Quick Reference:** See [CHEATSHEET.md](./CHEATSHEET.md) for visual patterns and common code snippets.

```bash
# Generate Wire dependencies
wire gen ./internal/di

# Generate mocks
mockery

# Generate sqlc code
sqlc generate

# Run tests
go test ./...
go test -race -cover ./...

# Lint code
golangci-lint run

# Format code
gofmt -w .
goimports -w .

# Database migrations
atlas migrate apply --env local

# Run server
go run cmd/main.go
```

## 📝 Contribution Guidelines

When adding or modifying code:

1. ✅ Follow naming conventions ([007-coding-standards.md](./007-coding-standards.md))
2. ✅ Add validation tags to models ([002-domain-layer.md](./002-domain-layer.md))
3. ✅ Map errors appropriately ([005-repository-layer.md](./005-repository-layer.md#error-mapping))
4. ✅ Write tests for new features ([006-testing.md](./006-testing.md))
5. ✅ Document exported symbols ([007-coding-standards.md](./007-coding-standards.md#comments))
6. ✅ Run linter before commit ([007-coding-standards.md](./007-coding-standards.md#linting))
7. ✅ Update documentation if changing patterns

## 🔍 Need Help?

- **Architecture questions:** Start with [002-domain-layer.md](./002-domain-layer.md)
- **Code style questions:** Check [007-coding-standards.md](./007-coding-standards.md)
- **Database questions:** See [008-database-conventions.md](./008-database-conventions.md)
- **Testing questions:** Reference [006-testing.md](./006-testing.md)

---

**Last Updated:** 2025-11-22
**Maintainer:** Development Team
