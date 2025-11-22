# Domain Layer Rules

## Overview

The domain layer (`internal/domain`) is the core of the application and contains business logic, entities, and interfaces. It should have no dependencies on infrastructure or external frameworks.

## Structure

### Domain Models

Domain models are pure Go structs representing business entities with validation tags:

- **Struct tags:**
  - `json` - JSON serialization field names (camelCase)
  - `binding` - Gin binding validation rules
  - `validate` - Go Playground validator rules

**Example:**

```go
type Product struct {
    ID          uuid.UUID   `json:"id"          binding:"required"  validate:"required"`
    Name        string      `json:"name"        binding:"required"  validate:"required,gte=3,lte=200"`
    Description string      `json:"description" binding:"required"  validate:"required,gte=10"`
    Price       int64       `json:"price"       binding:"required"  validate:"required,gt=0"`
    CreatedAt   time.Time   `json:"createdAt"   binding:"required"  validate:"required"`
    UpdatedAt   time.Time   `json:"updatedAt"   binding:"required"  validate:"required,gtefield=CreatedAt"`
    DeletedAt   *time.Time  `json:"deletedAt"   validate:"omitempty,gtefield=CreatedAt"`
}
```

**Conventions:**

- Use `uuid.UUID` for all IDs (generate with `uuid.NewV7()`)
- Use `time.Time` for timestamps
- Use pointers for optional/nullable fields
- Use pointers for slice fields (`*[]Type`) to distinguish between nil and empty
- Money/currency values are stored as `int64` (cents/smallest unit)
- Soft delete with `DeletedAt *time.Time`

### Repository Interfaces

Repositories define minimal CRUD operations for aggregate roots:

**Required methods:**

- `List()` - Retrieve multiple entities with filters, sorting, pagination
- `Count()` - Count entities matching filters
- `Get()` - Retrieve a single entity by ID
- `Save()` - Upsert entity (insert or update)

**Conventions:**

- First parameter is always `context.Context`
- Return `error` as last return value
- Use pointer parameters for optional filters (`*string`, `*[]uuid.UUID`)
- Use custom enums for special parameters (e.g., `DeletedParam`)
- Repositories should NOT contain business logic

**Example:**

```go
type ProductRepository interface {
    List(
        ctx context.Context,
        ids *[]uuid.UUID,
        search *string,
        deleted DeletedParam,
        limit int,
        offset int,
    ) (*[]Product, error)

    Count(ctx context.Context, ids *[]uuid.UUID, deleted DeletedParam) (*int, error)
    Get(ctx context.Context, id uuid.UUID) (*Product, error)
    Save(ctx context.Context, entity Product) error
}
```

### Service Interfaces

Domain services contain business logic that doesn't naturally fit in a single entity:

**Conventions:**

- Factory methods for creating new entities (e.g., `Create()`, `CreateOption()`)
- Business logic operations (e.g., `AddValues()`, `UpdateQuantity()`)
- Validation is performed using `go-playground/validator`
- Return newly created entities as pointers
- Generate UUIDs using `uuid.NewV7()`

**Example:**

```go
type ProductService interface {
    Create(name string, description string, category Category) (*Product, error)
    CreateVariant(sku string, price int64, quantity int) (*ProductVariant, error)
}
```

### Domain Errors

All domain errors are defined in `internal/domain/error.go`:

```go
var (
    ErrConflict     = errors.New("conflict")
    ErrExists       = errors.New("already exists")
    ErrForbidden    = errors.New("forbidden")
    ErrInternal     = errors.New("internal error")
    ErrInvalid      = errors.New("invalid data")
    ErrNotFound     = errors.New("not found")
    ErrServiceError = errors.New("service error")
    ErrTimeout      = errors.New("timeout")
    ErrUnavailable  = errors.New("service unavailable")
    ErrUnknown      = errors.New("unknown error")
)
```

**Error Wrapping:**
Use `github.com/hashicorp/go-multierror` for wrapping:

```go
return nil, multierror.Append(domain.ErrInternal, err)
```

### Enums and Common Types

Define shared enums and types in `internal/domain/common.go`:

```go
type DeletedParam string

const (
    DeletedExclude DeletedParam = "exclude"
    DeletedOnly    DeletedParam = "only"
    DeletedAll     DeletedParam = "all"
)
```

## Testing

### Mock Generation

Repository mocks are auto-generated using mockery with testify:

**Configuration:** `.mockery.yml`

**Generated files:** `internal/domain/*_repository_mock.go`

**Usage in tests:**

```go
mockRepo := domain.NewMockProductRepository(t)
mockRepo.EXPECT().Get(ctx, productID).Return(&product, nil)
```

## Rules Summary

1. ✅ Domain models are pure structs with validation tags
2. ✅ Use `uuid.UUID` for IDs, `time.Time` for timestamps
3. ✅ Use pointers for optional fields and slices
4. ✅ Repositories have only List, Count, Get, Save methods
5. ✅ Services contain business logic and factory methods
6. ✅ All domain errors defined in `error.go`
7. ✅ Wrap errors with `multierror.Append()`
8. ✅ No infrastructure dependencies in domain layer
9. ✅ Generate mocks with mockery for testing
