# Domain Layer (`internal/domain`)

## Purpose

Define contracts (interfaces), models, errors. **No implementations.**

## Domain Models

```go
type Product struct {
    ID          uuid.UUID   `json:"id"          binding:"required"  validate:"required"`
    Name        string      `json:"name"        binding:"required"  validate:"required,gte=3,lte=200"`
    Price       int64       `json:"price"       binding:"required"  validate:"required,gt=0"`
    CreatedAt   time.Time   `json:"createdAt"   binding:"required"  validate:"required"`
    UpdatedAt   time.Time   `json:"updatedAt"   binding:"required"  validate:"required,gtefield=CreatedAt"`
    DeletedAt   *time.Time  `json:"deletedAt"   validate:"omitempty,gtefield=CreatedAt"`
}
```

**Conventions:**
- ✅ `uuid.UUID` for IDs (generate with `uuid.NewV7()`)
- ✅ `time.Time` for timestamps
- ✅ Pointers for optional fields (`*time.Time`, `*string`)
- ✅ Pointers for slices (`*[]Type`) to distinguish nil vs empty
- ✅ Money as `int64` (cents/smallest unit)
- ✅ Soft delete with `DeletedAt *time.Time`

**Tags:**
- `json` - camelCase field names
- `binding` - Gin validation
- `validate` - go-playground/validator rules

## Repository Interfaces

**Minimal CRUD only:**

```go
type ProductRepository interface {
    List(ctx context.Context, ids *[]uuid.UUID, search *string, 
         deleted DeletedParam, limit int, offset int) (*[]Product, error)
    Count(ctx context.Context, ids *[]uuid.UUID, deleted DeletedParam) (*int, error)
    Get(ctx context.Context, id uuid.UUID) (*Product, error)
    Save(ctx context.Context, entity Product) error  // Upsert
}
```

**Rules:**
- ✅ First param is `context.Context`
- ✅ Pointers for optional filters (`*string`, `*[]uuid.UUID`)
- ✅ Return `error` as last value
- ✅ NO business logic in repos

## Service Interfaces

```go
type ProductService interface {
    Create(name string, description string) (*Product, error)
    AddVariant(product *Product, variant Variant) error
}
```

**Rules:**
- ✅ Factory methods (Create, CreateX)
- ✅ Business logic operations
- ✅ Return pointers for new entities

## Domain Errors

**File:** `internal/domain/error.go`

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

**Wrap with `multierror.Append`:**
```go
return nil, multierror.Append(domain.ErrInternal, err)
```

## Enums & Common Types

**File:** `internal/domain/common.go`

```go
type DeletedParam string

const (
    DeletedExclude DeletedParam = "exclude"
    DeletedOnly    DeletedParam = "only"
    DeletedAll     DeletedParam = "all"
)
```

## Mock Generation

**Config:** `.mockery.yml`

**Generate:** `mockery`

**Output:** `internal/domain/*_repository_mock.go`

**Usage:**
```go
mockRepo := domain.NewMockProductRepository(t)
mockRepo.EXPECT().Get(ctx, id).Return(&product, nil)
```

## Validation Tags

- `required` - Must be present
- `omitempty` - Optional
- `omitnil` - Pointer, validate if not nil
- `gt=0`, `gte=0`, `lte=100`, `lt=100` - Numeric constraints
- `dive` - Validate slice elements
- `gtefield=OtherField` - Field comparison

## Quick Rules

1. ✅ Pure structs, interfaces only
2. ✅ UUID for IDs, time.Time for timestamps
3. ✅ Pointers for optional/nullable/slices
4. ✅ Repos: List, Count, Get, Save (minimal CRUD)
5. ✅ Services: factory + business logic
6. ✅ All errors defined in `error.go`
7. ✅ Wrap errors with `multierror.Append()`
8. ❌ No implementations, no infrastructure deps
