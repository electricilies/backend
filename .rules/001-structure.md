# Project Structure

## Tech Stack

- **Framework:** Go 1.23+, Gin, Google Wire
- **Database:** PostgreSQL 16, sqlc
- **Infrastructure:** S3 (MinIO), Keycloak, Redis
- **Testing:** mockery, testify, testcontainers

## Directory Structure

```
.
├── cmd/                    # Application entrypoint
├── internal/
│   ├── application/       # Use cases, orchestration, params
│   ├── client/            # External adapters (DB, Redis, S3, etc.)
│   ├── delivery/          # HTTP handlers (Gin)
│   ├── di/                # Dependency injection (Wire)
│   ├── domain/            # Interfaces, models, errors, enums
│   │                        # *_repository_mock.go (generated)
│   ├── infrastructure/    # Repository implementations
│   └── service/           # Business logic implementations
├── database/
│   ├── schema.sql         # Main DB schema
│   ├── queries/           # sqlc query files
│   └── temporary-table/   # sqlc temp tables (not in schema)
├── migration/             # Database migrations
├── build/                 # Dockerfiles
├── docker/                # Docker Compose
├── terraform/             # IaC (Keycloak, MinIO)
├── keycloak/              # Realm exports
├── pkg/logger/            # Logging utilities
├── docs/                  # API docs, Swagger
└── .rules/                # Project rules (this)
```

## Layer Organization

**Domain Layer:** `internal/domain`
- Interfaces (repositories, services)
- Domain models (entities)
- Domain errors
- Enums and types
- ✅ No implementations, pure contracts

**Service Layer:** `internal/service`
- Domain service implementations
- Factory methods
- ✅ Only dependency: `*validator.Validate`

**Application Layer:** `internal/application`
- Use case implementations
- Orchestration logic
- Parameter objects (DTOs)
- ✅ Dependencies: repos + services + adapters

**Repository Layer:** `internal/infrastructure/repository`
- Repository implementations
- sqlc integration
- Error mapping

**Delivery Layer:** `internal/delivery`
- HTTP handlers (Gin)
- Request/response mapping
- Error to HTTP status mapping

## Key Files

- `.mockery.yml` - Mock generation config
- `sqlc.yaml` - sqlc configuration
- `atlas.hcl` - Migration tool config
- `.golangci.yaml` - Linter configuration
- `wire.go` - Dependency injection

## Error Flow

Infrastructure → Domain → Service → Application → Delivery (HTTP)

PostgreSQL errors mapped to domain errors, then to HTTP status codes.
