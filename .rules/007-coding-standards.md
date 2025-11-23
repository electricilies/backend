# Coding Standards

## Naming

**Packages:** `product`, `user` (singular, lowercase)  
**Types:** `Product`, `ProductService` (PascalCase)  
**Functions:** `Create`, `GetProduct` (PascalCase exported)  
**Variables:** `productID`, `userEmail` (camelCase)  
**Constants:** `MaxRetries`, `DefaultTimeout` (PascalCase)

## Import Groups

```go
import (
    // 1. Standard library
    "context"
    "errors"
    "time"
    
    // 2. External dependencies
    "github.com/google/uuid"
    "github.com/gin-gonic/gin"
    
    // 3. Internal packages
    "backend/internal/domain"
    "backend/internal/application"
)
```

## Error Handling

```go
// ✅ Always check errors
result, err := doSomething()
if err != nil {
    return nil, multierror.Append(domain.ErrInternal, err)
}

// ❌ Never ignore
result, _ := doSomething()  // BAD

// ✅ Defer cleanup with error check
defer func() {
    if err := file.Close(); err != nil {
        log.Error("failed to close", zap.Error(err))
    }
}()
```

## Context

```go
// ✅ First parameter
func (r *Repo) Get(ctx context.Context, id uuid.UUID) (*Product, error)

// ❌ Don't store in structs
type Service struct {
    ctx context.Context  // BAD
}
```

## Pointers vs Values

**Use pointers:**
- Large structs (>100 bytes)
- Optional fields (nil vs zero)
- Slices when nil ≠ empty: `*[]Type`

**Use values:**
- Small structs
- Immutable data
- UUIDs, timestamps

```go
type Product struct {
    ID       uuid.UUID    // Value
    Name     string       // Value
    Category *Category    // Pointer (optional)
    Variants *[]Variant   // Pointer (nil ≠ empty)
}
```

## Nil Checks

```go
// ✅ Early returns
if input == nil {
    return domain.ErrInvalid
}

// ✅ Check before dereference
if product.Category != nil {
    fmt.Println(product.Category.Name)
}
```

## Comments

```go
// ✅ Document exported symbols
// Product represents a purchasable item.
type Product struct {
    ID uuid.UUID
}

// Create creates a new product.
func (s *Service) Create(name string) (*Product, error)

// ✅ Explain complex logic
// Calculate trending score: views * 0.3 + purchases * 0.7
score := int64(float64(views)*0.3 + float64(purchases)*0.7)

// ❌ Don't state the obvious
counter++  // No comment needed
```

## Magic Numbers

```go
// ✅ Use constants
const (
    DefaultPageSize = 10
    MaxPageSize     = 100
)

// ❌ Avoid literals
limit := 10  // BAD
```

## JSON Tags

```go
type Product struct {
    ID        uuid.UUID  `json:"id"        binding:"required"  validate:"required"`
    Name      string     `json:"name"      binding:"required"  validate:"required,gte=3"`
    Price     int64      `json:"price"     binding:"required"  validate:"required,gt=0"`
}
```

**Conventions:**
- camelCase field names
- Align tags for readability

## Performance

### Preallocate Slices

```go
// ✅ Good
products := make([]Product, 0, len(ids))
for _, id := range ids {
    products = append(products, getProduct(id))
}

// ❌ Bad
var products []Product
for _, id := range ids {
    products = append(products, getProduct(id))
}
```

### String Concatenation

```go
// ✅ Use strings.Builder
var builder strings.Builder
for _, str := range strings {
    builder.WriteString(str)
}
result := builder.String()

// ❌ Avoid += in loops
var result string
for _, str := range strings {
    result += str  // BAD
}
```

### Defer in Loops

```go
// ✅ Extract function
func ProcessFiles(files []string) error {
    for _, filename := range files {
        if err := processFile(filename); err != nil {
            return err
        }
    }
    return nil
}

func processFile(filename string) error {
    f, _ := os.Open(filename)
    defer f.Close()  // Runs after each file
    // ...
}

// ❌ Defer in loop
for _, filename := range files {
    f, _ := os.Open(filename)
    defer f.Close()  // Won't run until function returns
}
```

## Linting

```bash
golangci-lint run
gofmt -w .
goimports -w .
```

**Key linters:**
- `errcheck` - Check error handling
- `gofmt` - Format code
- `govet` - Suspicious constructs
- `staticcheck` - Static analysis
- `unused` - Unused code

## Quick Rules

1. ✅ Follow standard Go naming
2. ✅ Group imports (stdlib, external, internal)
3. ✅ Always check and wrap errors
4. ✅ Context as first parameter
5. ✅ Pointers for large/optional/mutable
6. ✅ Document exported symbols
7. ✅ Constants for magic numbers
8. ✅ JSON tags in camelCase
9. ✅ Preallocate slices when size known
10. ✅ strings.Builder for concatenation
11. ✅ Run golangci-lint before commit
12. ✅ Soft limit: 100 chars, hard: 120 chars
