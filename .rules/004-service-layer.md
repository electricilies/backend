# Service Layer Rules

## Overview

The service layer (`internal/service`) implements domain service interfaces defined in `internal/domain`. It contains **pure business logic** for entity creation, validation, and domain operations.

**Important:** Services do NOT have repositories. They are pure business logic without any infrastructure dependencies. Repository access happens in the application layer.

## Structure

Each domain service implementation:

- Implements a domain service interface
- Contains factory methods for creating entities
- Performs validation using `go-playground/validator`
- Wraps errors with domain errors
- **Contains NO repository dependencies** (pure business logic)
- **Contains NO external adapter dependencies** (S3, Redis, etc.)

## Service Implementation

```go
package service

import (
	"backend/internal/domain"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/hashicorp/go-multierror"
)

type Product struct {
	validate *validator.Validate
}

func ProvideProduct(validate *validator.Validate) *Product {
	return &Product{
		validate: validate,
	}
}

// Ensure implementation satisfies interface
var _ domain.ProductService = &Product{}
```

**Conventions:**

- Service struct matches domain entity name (e.g., `Product`, `Category`)
- Provider function: `Provide<Entity>(validate *validator.Validate) *<Entity>`
- Assert interface implementation: `var _ domain.<Entity>Service = &<Entity>{}`
- **Only dependency is validator** - no repositories or adapters
- Services are pure business logic

## Factory Methods

Factory methods create new domain entities with generated IDs and default values:

```go
func (p *Product) Create(
    name string,
    description string,
    category domain.Category,
) (*domain.Product, error) {
    // 1. Generate UUID v7 (time-ordered)
    id, err := uuid.NewV7()
    if err != nil {
        return nil, multierror.Append(domain.ErrInternal, err)
    }

    // 2. Create entity with required fields
    product := &domain.Product{
        ID:          id,
        Name:        name,
        Description: description,
        Category:    &category,
        ViewsCount:  0,
        TotalPurchase: 0,
        Rating:      0.0,
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
    }

    // 3. Validate entity
    if err := p.validate.Struct(product); err != nil {
        return nil, multierror.Append(domain.ErrInvalid, err)
    }

    return product, nil
}
```

**Conventions:**

- Generate UUIDs with `uuid.NewV7()` (time-ordered, better for DB indexes)
- Wrap UUID errors with `domain.ErrInternal`
- Set default values (counts to 0, timestamps to `time.Now()`)
- Always validate with `p.validate.Struct(entity)`
- Wrap validation errors with `domain.ErrInvalid`
- Return pointer to entity

## Business Logic Methods

Domain services also contain business logic that operates on entities:

```go
func (p *Product) AddVariant(
    product *domain.Product,
    variant domain.ProductVariant,
) error {
    // 1. Validate business rules
    if product.DeletedAt != nil {
        return domain.ErrInvalid
    }

    // 2. Validate variant
    if err := p.validate.Struct(variant); err != nil {
        return multierror.Append(domain.ErrInvalid, err)
    }

    // 3. Modify entity
    if product.Variants == nil {
        product.Variants = &[]domain.ProductVariant{}
    }
    *product.Variants = append(*product.Variants, variant)

    // 4. Update timestamps
    product.UpdatedAt = time.Now()

    return nil
}
```

**Pattern:**

1. Validate business rules
2. Validate input with `go-playground/validator`
3. Modify entity state
4. Update timestamps
5. Return error or nil

## Error Handling

### Wrapping Errors

Always wrap errors with domain errors:

```go
// Internal errors (UUID generation, unexpected failures)
if err != nil {
    return nil, multierror.Append(domain.ErrInternal, err)
}

// Validation errors
if err := p.validate.Struct(entity); err != nil {
    return nil, multierror.Append(domain.ErrInvalid, err)
}

// Business rule violations
if product.Quantity < 0 {
    return multierror.Append(domain.ErrInvalid, errors.New("quantity cannot be negative"))
}
```

### Domain Error Usage

| Domain Error   | Use Case                                            |
| -------------- | --------------------------------------------------- |
| `ErrInvalid`   | Validation failures, business rule violations       |
| `ErrInternal`  | Unexpected system errors (UUID gen, panic recovery) |
| `ErrNotFound`  | Entity not found (typically from repository)        |
| `ErrExists`    | Duplicate entity creation                           |
| `ErrConflict`  | Concurrent modification conflicts                   |
| `ErrForbidden` | Authorization failures                              |

## Validation

### Struct Validation

Use `go-playground/validator` for struct validation:

```go
if err := p.validate.Struct(product); err != nil {
    return nil, multierror.Append(domain.ErrInvalid, err)
}
```

### Field Validation

Entities use validation tags:

- `required` - Field must be present
- `omitempty` - Optional field
- `omitnil` - Pointer field, validate if not nil
- `gt=0`, `gte=0` - Numeric constraints
- `lte=100`, `lt=100` - Maximum values
- `dive` - Validate slice elements
- `gtefield=OtherField` - Field comparison

```go
type Product struct {
    Price     int64      `validate:"required,gt=0"`
    Rating    float64    `validate:"required,gte=0,lte=5"`
    UpdatedAt time.Time  `validate:"required,gtefield=CreatedAt"`
    Variants  *[]Variant `validate:"omitnil,gte=1,dive"`
}
```

## Testing

Service layer should be unit tested with mock validator:

```go
func TestProduct_Create(t *testing.T) {
    validate := validator.New()
    service := ProvideProduct(validate)

    category := domain.Category{ID: uuid.New(), Name: "Electronics"}
    product, err := service.Create("Laptop", "Gaming laptop", category)

    assert.NoError(t, err)
    assert.NotNil(t, product)
    assert.NotEqual(t, uuid.Nil, product.ID)
    assert.Equal(t, "Laptop", product.Name)
}
```

## Rules Summary

1. ✅ Service structs implement domain service interfaces
2. ✅ Provider function: `Provide<Entity>(validate *validator.Validate)`
3. ✅ **Services have NO repositories** - pure business logic only
4. ✅ **Services have NO external adapters** - no S3, Redis, etc.
5. ✅ Factory methods generate UUIDs with `uuid.NewV7()`
6. ✅ Always validate entities with `p.validate.Struct()`
7. ✅ Wrap UUID errors with `domain.ErrInternal`
8. ✅ Wrap validation errors with `domain.ErrInvalid`
9. ✅ Business logic methods modify entities in place
10. ✅ Update timestamps after modifications
11. ✅ Use `multierror.Append()` for error wrapping
12. ✅ Assert interface implementation with `var _`
