# Coding Standards

## Go Style

Follow standard Go conventions and best practices:

### Naming Conventions

**Packages:**

- Use singular nouns: `product`, `user`, `order` (not `products`, `users`)
- Short, lowercase, no underscores: `domain`, `service`, `application`

**Types:**

- PascalCase for exported: `Product`, `ProductService`
- camelCase for unexported: `productImpl`, `mapper`
- Interface implementations: `ProductImpl`, `ProductRepository`

**Functions:**

- PascalCase for exported: `Create`, `GetProduct`
- camelCase for unexported: `mapError`, `validateInput`
- Provider functions: `ProvideProduct`, `ProvideDatabase`

**Variables:**

- camelCase for all: `productID`, `userEmail`
- Use short names in limited scope: `i`, `err`, `ctx`
- Descriptive names in wider scope: `productRepository`, `attributeService`

**Constants:**

- PascalCase for exported: `MaxRetries`, `DefaultTimeout`
- camelCase for unexported: `defaultLimit`, `maxPageSize`

### Code Organization

**Import Groups:**

```go
import (
    // Standard library
    "context"
    "errors"
    "time"

    // External dependencies
    "github.com/google/uuid"
    "github.com/gin-gonic/gin"

    // Internal packages
    "backend/internal/domain"
    "backend/internal/application"
)
```

**Order in files:**

1. Package declaration
2. Imports
3. Constants
4. Type definitions
5. Constructor/provider functions
6. Interface assertions (`var _ Interface = &Implementation{}`)
7. Methods (exported first, unexported after)
8. Helper functions

### Error Handling

**Always check errors:**

```go
// ✅ Good
result, err := doSomething()
if err != nil {
    return nil, err
}

// ❌ Bad
result, _ := doSomething()
```

**Wrap errors with context:**

```go
// ✅ Good
if err != nil {
    return nil, multierror.Append(domain.ErrInternal, err)
}

// ❌ Bad
if err != nil {
    return nil, err
}
```

**Don't ignore errors:**

```go
// ✅ Good
defer func() {
    if err := file.Close(); err != nil {
        log.Error("failed to close file", zap.Error(err))
    }
}()

// ❌ Bad
defer file.Close()
```

### Pointers vs Values

**Use pointers for:**

- Large structs (>100 bytes)
- Mutable state
- Optional fields (distinguishing nil from zero value)
- Slices when nil ≠ empty: `*[]Type`

**Use values for:**

- Small structs
- Immutable data
- Parameters that don't need modification

```go
// ✅ Good
type Product struct {
    ID        uuid.UUID    // Value (16 bytes)
    Name      string       // Value (small)
    Category  *Category    // Pointer (optional, mutable)
    Variants  *[]Variant   // Pointer (nil ≠ empty)
}

// ❌ Bad
type Product struct {
    ID        *uuid.UUID   // Unnecessary pointer
    Name      *string      // Unnecessary pointer
}
```

### Context Usage

**Always pass context as first parameter:**

```go
// ✅ Good
func (r *Repository) Get(ctx context.Context, id uuid.UUID) (*Product, error)

// ❌ Bad
func (r *Repository) Get(id uuid.UUID, ctx context.Context) (*Product, error)
```

**Don't store context in structs:**

```go
// ✅ Good
type Service struct {
    repo Repository
}

func (s *Service) DoWork(ctx context.Context) error {
    return s.repo.Get(ctx, id)
}

// ❌ Bad
type Service struct {
    ctx  context.Context
    repo Repository
}
```

### Nil Checks

**Check nil pointers before dereferencing:**

```go
// ✅ Good
if product.Category != nil {
    fmt.Println(product.Category.Name)
}

// ❌ Bad
fmt.Println(product.Category.Name) // panic if nil
```

**Use early returns:**

```go
// ✅ Good
func Process(input *Input) error {
    if input == nil {
        return domain.ErrInvalid
    }

    // continue processing
}

// ❌ Bad
func Process(input *Input) error {
    if input != nil {
        // deeply nested processing
    }
    return nil
}
```

## Code Quality

### Linting

Use golangci-lint with project configuration:

```bash
golangci-lint run
```

**File:** `.golangci.yaml`

Key enabled linters:

- `errcheck` - Check error handling
- `gofmt` - Code formatting
- `govet` - Suspicious constructs
- `staticcheck` - Static analysis
- `unused` - Unused code detection
- `gosimple` - Simplification suggestions
- `ineffassign` - Ineffectual assignments

### Formatting

**Use `gofmt` or `goimports`:**

```bash
gofmt -w .
goimports -w .
```

**Line length:**

- Soft limit: 100 characters
- Hard limit: 120 characters

**Function length:**

- Keep functions under 50 lines
- Extract helper functions if longer

### Comments

**Document exported symbols:**

```go
// Product represents a purchasable item in the catalog.
type Product struct {
    ID   uuid.UUID
    Name string
}

// Create creates a new product with the given parameters.
// It validates the input and returns an error if validation fails.
func (s *Service) Create(name string, description string) (*Product, error) {
    // implementation
}
```

**Don't comment obvious code:**

```go
// ❌ Bad
// Increment counter
counter++

// ✅ Good (no comment needed)
counter++
```

**Explain complex logic:**

```go
// ✅ Good
// Calculate trending score using weighted formula:
// score = views * 0.3 + purchases * 0.7
trendingScore := int64(float64(views)*0.3 + float64(purchases)*0.7)
```

### Magic Numbers

**Use constants:**

```go
// ✅ Good
const (
    DefaultPageSize = 10
    MaxPageSize     = 100
)

func List(page int) []Product {
    if page < 1 {
        page = 1
    }
    limit := DefaultPageSize
    // ...
}

// ❌ Bad
func List(page int) []Product {
    if page < 1 {
        page = 1
    }
    limit := 10
    // ...
}
```

## JSON and Tags

### Field Naming

Use camelCase in JSON tags:

```go
type Product struct {
    ID            uuid.UUID  `json:"id"`
    Name          string     `json:"name"`
    CreatedAt     time.Time  `json:"createdAt"`
    CategoryID    uuid.UUID  `json:"categoryId"`
}
```

### Tag Alignment

Align tags for readability:

```go
type Product struct {
    ID          uuid.UUID  `json:"id"          binding:"required"  validate:"required"`
    Name        string     `json:"name"        binding:"required"  validate:"required,gte=3"`
    Price       int64      `json:"price"       binding:"required"  validate:"required,gt=0"`
    Rating      float64    `json:"rating"      binding:"required"  validate:"gte=0,lte=5"`
}
```

## Dependency Injection

### Google Wire

**Provider functions:**

```go
// Return pointer to implementation
func ProvideProduct(repo Repository) *ProductImpl {
    return &ProductImpl{repo: repo}
}

// Bind interface to implementation in wire.go
wire.Bind(new(Product), new(*ProductImpl))
```

**Wire file:**

```go
//go:build wireinject
// +build wireinject

package di

func InitializeApp() (*App, error) {
    wire.Build(
        // Clients
        client.ProvideDatabase,
        client.ProvideRedis,

        // Repositories
        repository.ProvideProduct,
        wire.Bind(new(domain.ProductRepository), new(*repository.ProductRepository)),

        // Services
        service.ProvideProduct,
        wire.Bind(new(domain.ProductService), new(*service.Product)),

        // Application
        application.ProvideProduct,
        wire.Bind(new(application.Product), new(*application.ProductImpl)),
    )

    return &App{}, nil
}
```

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

### Avoid String Concatenation in Loops

```go
// ✅ Good
var builder strings.Builder
for _, str := range strings {
    builder.WriteString(str)
}
result := builder.String()

// ❌ Bad
var result string
for _, str := range strings {
    result += str
}
```

### Use Defer Carefully

```go
// ✅ Good - defer outside loop
func ProcessFiles(files []string) error {
    for _, filename := range files {
        if err := processFile(filename); err != nil {
            return err
        }
    }
    return nil
}

func processFile(filename string) error {
    f, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer f.Close()

    // process file
    return nil
}

// ❌ Bad - defer in loop
func ProcessFiles(files []string) error {
    for _, filename := range files {
        f, err := os.Open(filename)
        if err != nil {
            return err
        }
        defer f.Close() // won't run until function returns

        // process file
    }
    return nil
}
```

## Rules Summary

1. ✅ Follow standard Go naming conventions
2. ✅ Group imports (stdlib, external, internal)
3. ✅ Always check and wrap errors
4. ✅ Context as first parameter
5. ✅ Use pointers for large/optional/mutable data
6. ✅ Document all exported symbols
7. ✅ Use constants instead of magic numbers
8. ✅ JSON tags in camelCase
9. ✅ Align struct tags for readability
10. ✅ Run golangci-lint before committing
11. ✅ Preallocate slices when size known
12. ✅ Use strings.Builder for concatenation
