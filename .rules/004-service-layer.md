# Service Layer (`internal/service`)

## Purpose

Pure business logic. **ONLY dependency: `*validator.Validate`**

## Implementation

```go
type Product struct {
    validate *validator.Validate  // ONLY dependency
}

func ProvideProduct(validate *validator.Validate) *Product {
    return &Product{validate: validate}
}

var _ domain.ProductService = &Product{}  // Assert interface
```

❌ **NO repositories**  
❌ **NO external adapters** (S3, Redis, etc.)

## Factory Methods

```go
func (s *Product) Create(
    name string,
    description string,
    category domain.Category,
) (*domain.Product, error) {
    // 1. Generate UUID v7 (time-ordered)
    id, err := uuid.NewV7()
    if err != nil {
        return nil, multierror.Append(domain.ErrInternal, err)
    }
    
    // 2. Build entity
    product := &domain.Product{
        ID:            id,
        Name:          name,
        Description:   description,
        Category:      &category,
        ViewsCount:    0,
        TotalPurchase: 0,
        Rating:        0.0,
        CreatedAt:     time.Now(),
        UpdatedAt:     time.Now(),
    }
    
    // 3. Validate
    if err := s.validate.Struct(product); err != nil {
        return nil, multierror.Append(domain.ErrInvalid, err)
    }
    
    return product, nil
}
```

## Business Logic Methods

```go
func (s *Product) AddVariant(
    product *domain.Product,
    variant domain.ProductVariant,
) error {
    // 1. Business rules
    if product.DeletedAt != nil {
        return domain.ErrInvalid
    }
    
    // 2. Validate input
    if err := s.validate.Struct(variant); err != nil {
        return multierror.Append(domain.ErrInvalid, err)
    }
    
    // 3. Modify entity
    if product.Variants == nil {
        product.Variants = &[]domain.ProductVariant{}
    }
    *product.Variants = append(*product.Variants, variant)
    
    // 4. Update timestamp
    product.UpdatedAt = time.Now()
    
    return nil
}
```

**Pattern:**
1. Validate business rules
2. Validate input
3. Modify entity state
4. Update timestamps

## Error Wrapping

```go
// Internal errors (UUID generation)
if err != nil {
    return nil, multierror.Append(domain.ErrInternal, err)
}

// Validation errors
if err := s.validate.Struct(entity); err != nil {
    return nil, multierror.Append(domain.ErrInvalid, err)
}

// Business rule violations
if product.Quantity < 0 {
    return multierror.Append(domain.ErrInvalid, errors.New("negative quantity"))
}
```

## Error Usage

| Error | When |
|-------|------|
| `ErrInvalid` | Validation/business rule failures |
| `ErrInternal` | UUID generation, unexpected errors |
| `ErrNotFound` | Repo layer (not service) |
| `ErrExists` | Duplicate detection |
| `ErrConflict` | Concurrent modifications |
| `ErrForbidden` | Authorization failures |

## Validation Tags

```go
type Product struct {
    Price     int64      `validate:"required,gt=0"`
    Rating    float64    `validate:"required,gte=0,lte=5"`
    UpdatedAt time.Time  `validate:"required,gtefield=CreatedAt"`
    Variants  *[]Variant `validate:"omitnil,gte=1,dive"`
}
```

- `required` - Must be present
- `omitempty` - Optional
- `omitnil` - Pointer, validate if not nil
- `gt=0`, `gte=0`, `lte=100` - Numeric constraints
- `dive` - Validate slice elements
- `gtefield=X` - Field comparison

## Testing

```go
func TestProduct_Create(t *testing.T) {
    validate := validator.New()
    service := ProvideProduct(validate)
    
    category := domain.Category{ID: uuid.New(), Name: "Electronics"}
    product, err := service.Create("Laptop", "Gaming laptop", category)
    
    assert.NoError(t, err)
    assert.NotNil(t, product)
    assert.NotEqual(t, uuid.Nil, product.ID)
}
```

✅ **No mocks needed** - only validator  
✅ **No database** - pure logic testing  

## Quick Rules

1. ✅ ONLY dependency: `*validator.Validate`
2. ❌ NO repositories or adapters
3. ✅ Provider: `Provide<Entity>(validate *validator.Validate)`
4. ✅ Generate IDs with `uuid.NewV7()`
5. ✅ Always validate with `s.validate.Struct()`
6. ✅ Wrap UUID errors with `ErrInternal`
7. ✅ Wrap validation errors with `ErrInvalid`
8. ✅ Update timestamps after modifications
9. ✅ Use `multierror.Append()` for wrapping
10. ✅ Assert interface: `var _ domain.XService = &X{}`
