# Architecture Cheatsheet

Quick reference for layer patterns and common operations.

> 📖 **See also:** [100-business-logic.md](./100-business-logic.md) for domain-specific business rules

## 🏗️ Layer Dependencies

```
┌─────────────────────────────────────────────────┐
│ DOMAIN (internal/domain)                        │
│ • Models, Interfaces, Errors                    │
│ • NO implementations, NO dependencies           │
└─────────────────────────────────────────────────┘
                       ↑
                       │ implements
            ┌──────────┴──────────┐
            │                     │
┌───────────────────────┐ ┌─────────────────────────────────────┐
│ SERVICE               │ │ REPOSITORY                          │
│ (internal/service)    │ │ (internal/infrastructure/repository)│
│                       │ │                                     │
│ Dependencies:         │ │ Dependencies:                       │
│ • validator ✅        │ │ • pgxpool.Pool ✅                   │
│ • NO repos ❌         │ │ • sqlc queries ✅                   │
│ • NO adapters ❌      │ │                                     │
└───────────────────────┘ └─────────────────────────────────────┘
            │                     │
            │                     │
            └──────────┬──────────┘
                       ↓ injects both
┌─────────────────────────────────────────────────┐
│ APPLICATION (internal/application)              │
│                                                 │
│ Dependencies:                                   │
│ • Multiple repos (product, category, etc.) ✅   │
│ • Domain services ✅                            │
│ • External adapters (S3, Redis) ✅              │
└─────────────────────────────────────────────────┘
                       ↑
                       │
┌─────────────────────────────────────────────────┐
│ DELIVERY (HTTP handlers)                        │
└─────────────────────────────────────────────────┘
```

## 📋 Quick Patterns

### Domain Model

```go
type Product struct {
    ID        uuid.UUID  `json:"id"        binding:"required" validate:"required"`
    Name      string     `json:"name"      binding:"required" validate:"required,gte=3,lte=200"`
    Price     int64      `json:"price"     binding:"required" validate:"required,gt=0"`
    CreatedAt time.Time  `json:"createdAt" binding:"required" validate:"required"`
    DeletedAt *time.Time `json:"deletedAt" validate:"omitempty,gtefield=CreatedAt"`
}
```

### Repository Interface

```go
type ProductRepository interface {
    List(ctx context.Context, filters..., limit, offset int) (*[]Product, error)
    Count(ctx context.Context, filters...) (*int, error)
    Get(ctx context.Context, id uuid.UUID) (*Product, error)
    Save(ctx context.Context, entity Product) error  // Upsert
}
```

### Service Implementation (NO REPOS!)

```go
type Product struct {
    validate *validator.Validate  // ONLY dependency
}

func ProvideProduct(validate *validator.Validate) *Product {
    return &Product{validate: validate}
}

func (p *Product) Create(name, description string) (*domain.Product, error) {
    id, _ := uuid.NewV7()
    product := &domain.Product{ID: id, Name: name, Description: description}
    
    if err := p.validate.Struct(product); err != nil {
        return nil, multierror.Append(domain.ErrInvalid, err)
    }
    
    return product, nil
}
```

### Application Implementation (MULTIPLE REPOS + ADAPTERS)

```go
type ProductImpl struct {
    productRepo    domain.ProductRepository   // ✅
    categoryRepo   domain.CategoryRepository  // ✅
    attributeRepo  domain.AttributeRepository // ✅
    productService domain.ProductService      // ✅
    s3Client       *s3.Client                 // ✅
    redisClient    *redis.Client              // ✅
}

func (p *ProductImpl) Create(ctx context.Context, param CreateProductParam) (*domain.Product, error) {
    // 1. Fetch from repos
    category, _ := p.categoryRepo.Get(ctx, param.CategoryID)
    
    // 2. Domain service (pure logic)
    product, _ := p.productService.Create(param.Name, param.Description)
    
    // 3. Use adapters
    p.s3Client.PutObject(ctx, ...)
    
    // 4. Persist
    p.productRepo.Save(ctx, *product)
    
    // 5. Cache invalidation
    p.redisClient.Del(ctx, "cache:key")
    
    return product, nil
}
```

### Repository Implementation

```go
type ProductRepository struct {
    db      *pgxpool.Pool
    queries *postgres.Queries
}

func (r *ProductRepository) Save(ctx context.Context, entity domain.Product) error {
    err := r.queries.UpsertProduct(ctx, postgres.UpsertProductParams{...})
    return mapError(err)  // Map to domain error
}

func mapError(err error) error {
    var pgErr *pgconn.PgError
    if errors.As(err, &pgErr) {
        switch pgErr.Code {
        case pgerrcode.UniqueViolation:
            return domain.ErrExists
        case pgerrcode.ForeignKeyViolation:
            return domain.ErrInvalid
        }
    }
    return domain.ErrInternal
}
```

## 🧪 Testing Patterns

### Service Test (NO MOCKS - Pure Logic)

```go
func TestProduct_Create(t *testing.T) {
    validate := validator.New()
    service := ProvideProduct(validate)  // Only validator!
    
    product, err := service.Create("Laptop", "Gaming laptop")
    
    assert.NoError(t, err)
    assert.NotNil(t, product)
}
```

### Application Test (Table-Driven with Mocks)

```go
func TestProductImpl_Create(t *testing.T) {
    tests := []struct {
        name       string
        setupMocks func(*MockProductRepo, *MockCategoryRepo, *MockProductService)
        wantErr    bool
    }{
        {
            name: "success",
            setupMocks: func(pr, cr, ps) {
                cr.EXPECT().Get(...).Return(&category, nil)
                ps.EXPECT().Create(...).Return(&product, nil)
                pr.EXPECT().Save(...).Return(nil)
            },
            wantErr: false,
        },
        {
            name: "category not found",
            setupMocks: func(pr, cr, ps) {
                cr.EXPECT().Get(...).Return(nil, domain.ErrNotFound)
            },
            wantErr: true,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            mockPR := NewMockProductRepository(t)
            mockCR := NewMockCategoryRepository(t)
            mockPS := NewMockProductService(t)
            
            tt.setupMocks(mockPR, mockCR, mockPS)
            
            app := ProvideProduct(mockPR, mockCR, mockPS, nil, nil)
            _, err := app.Create(ctx, param)
            
            if tt.wantErr {
                assert.Error(t, err)
            } else {
                assert.NoError(t, err)
            }
        })
    }
}
```

### Repository Test (Integration with Test Containers)

```go
func TestProductRepository_Save(t *testing.T) {
    db, cleanup := setupTestDB(t)
    defer cleanup()
    
    repo := ProvideProductRepository(db)
    product := domain.Product{ID: uuid.New(), Name: "Test"}
    
    err := repo.Save(context.Background(), product)
    assert.NoError(t, err)
    
    retrieved, _ := repo.Get(context.Background(), product.ID)
    assert.Equal(t, product.Name, retrieved.Name)
}
```

## 🗄️ Database Patterns

### sqlc Query (Upsert)

```sql
-- name: UpsertProduct :exec
INSERT INTO products (id, name, price, created_at, updated_at)
VALUES (
    sqlc.arg('id'),
    sqlc.arg('name'),
    sqlc.arg('price'),
    sqlc.arg('created_at'),
    sqlc.arg('updated_at')
)
ON CONFLICT (id) DO UPDATE SET
    name = EXCLUDED.name,
    price = EXCLUDED.price,
    updated_at = EXCLUDED.updated_at;
```

### sqlc Query (List with Filters)

```sql
-- name: ListProducts :many
SELECT * FROM products
WHERE
    (sqlc.narg('ids')::uuid[] IS NULL OR id = ANY(sqlc.narg('ids')::uuid[]))
    AND (sqlc.narg('search') IS NULL OR name ILIKE '%' || sqlc.narg('search') || '%')
    AND (
        CASE sqlc.arg('deleted')
            WHEN 'exclude' THEN deleted_at IS NULL
            WHEN 'only' THEN deleted_at IS NOT NULL
            ELSE TRUE
        END
    )
ORDER BY created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');
```

## 🔧 Wire (Dependency Injection)

```go
//go:build wireinject

func InitializeApp() (*App, error) {
    wire.Build(
        // Clients
        client.ProvideDatabase,
        client.ProvideS3,
        client.ProvideRedis,
        
        // Repositories
        repository.ProvideProduct,
        wire.Bind(new(domain.ProductRepository), new(*repository.ProductRepository)),
        
        // Services (only validator)
        service.ProvideProduct,
        wire.Bind(new(domain.ProductService), new(*service.Product)),
        
        // Application (inject repos + services + adapters)
        application.ProvideProduct,
        wire.Bind(new(application.Product), new(*application.ProductImpl)),
    )
    return &App{}, nil
}
```

## ⚠️ Common Mistakes

### ❌ DON'T: Add Repository to Service

```go
// ❌ WRONG - Service should NOT have repos
type Product struct {
    validate *validator.Validate
    repo     domain.ProductRepository  // ❌ NO!
}
```

### ✅ DO: Service Only Has Validator

```go
// ✅ CORRECT - Service has pure logic only
type Product struct {
    validate *validator.Validate  // Only this!
}
```

### ❌ DON'T: Add Adapters to Service

```go
// ❌ WRONG - Service should NOT have adapters
type Product struct {
    validate    *validator.Validate
    s3Client    *s3.Client     // ❌ NO!
    redisClient *redis.Client  // ❌ NO!
}
```

### ✅ DO: Application Has Multiple Repos + Adapters

```go
// ✅ CORRECT - Application orchestrates everything
type ProductImpl struct {
    productRepo    domain.ProductRepository   // ✅
    categoryRepo   domain.CategoryRepository  // ✅
    productService domain.ProductService      // ✅
    s3Client       *s3.Client                 // ✅
    redisClient    *redis.Client              // ✅
}
```

## 📊 Error Flow

```
PostgreSQL Error (23505 unique_violation)
    ↓
Repository: mapError() → domain.ErrExists
    ↓
Service: multierror.Append(domain.ErrExists, err)
    ↓
Application: return err (no wrapping)
    ↓
Delivery: Map domain.ErrExists → HTTP 409 Conflict
```

## 🎯 Validation Flow

```
HTTP Request
    ↓
Delivery: Gin binding tags validate request
    ↓
Application: Pass to service
    ↓
Service: validate.Struct(entity) → business rules
    ↓
Return to Application → Delivery → HTTP Response
```

## 📝 Naming Conventions

| Item | Pattern | Example |
|------|---------|---------|
| Domain model | PascalCase | `Product`, `Category` |
| Repository interface | `<Entity>Repository` | `ProductRepository` |
| Service interface | `<Entity>Service` | `ProductService` |
| Service impl | `<Entity>` | `type Product struct` |
| Application interface | `<Entity>` | `type Product interface` |
| Application impl | `<Entity>Impl` | `type ProductImpl struct` |
| Provider function | `Provide<Entity>` | `ProvideProduct(...)` |
| JSON fields | camelCase | `"productId"`, `"createdAt"` |
| SQL columns | snake_case | `product_id`, `created_at` |

## 🚀 Quick Commands

```bash
# Generate mocks
mockery

# Generate sqlc
sqlc generate

# Generate wire
wire gen ./internal/di

# Run tests
go test ./...                    # All tests
go test -race ./...              # With race detector
go test -cover ./...             # With coverage
go test ./internal/service/...   # Specific package

# Lint
golangci-lint run

# Format
gofmt -w .
goimports -w .
```

---

**Remember:** Service = Pure Logic | Application = Orchestrator
