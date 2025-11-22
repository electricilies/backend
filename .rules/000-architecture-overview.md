# Architecture Overview

## Core Principles

This project follows **Domain-Driven Design (DDD)** with clear separation of concerns:

1. **Domain Layer** - Pure business logic, no infrastructure
2. **Service Layer** - Business logic implementations, **NO repositories or adapters**
3. **Application Layer** - Orchestration, injects multiple repositories and adapters
4. **Repository Layer** - Data persistence
5. **Delivery Layer** - HTTP handlers

## Layer Responsibilities

### 🎯 Domain Layer (`internal/domain`)

**Contains:**

- Domain models (entities, value objects)
- Repository interfaces
- Service interfaces
- Domain errors
- Enums and common types

**Rules:**

- ✅ Pure Go structs with validation tags
- ✅ Interfaces only, no implementations
- ✅ No infrastructure dependencies
- ✅ Define all domain errors here

**Example:**

```go
// Domain model
type Product struct {
    ID          uuid.UUID  `json:"id" validate:"required"`
    Name        string     `json:"name" validate:"required,gte=3"`
}

// Repository interface (minimal CRUD)
type ProductRepository interface {
    List(ctx context.Context, ...) (*[]Product, error)
    Count(ctx context.Context, ...) (*int, error)
    Get(ctx context.Context, id uuid.UUID) (*Product, error)
    Save(ctx context.Context, entity Product) error
}

// Service interface (business logic)
type ProductService interface {
    Create(name string, description string) (*Product, error)
}
```

---

### ⚙️ Service Layer (`internal/service`)

**Contains:**

- Domain service implementations
- Factory methods for entities
- Business logic and validation

**Dependencies:**

- ✅ **ONLY** `*validator.Validate`
- ❌ **NO** repositories
- ❌ **NO** external adapters (S3, Redis, etc.)

**Why no repos?** Services contain pure business logic that should be testable without any infrastructure.

**Example:**

```go
type Product struct {
    validate *validator.Validate  // ONLY dependency
}

func ProvideProduct(validate *validator.Validate) *Product {
    return &Product{validate: validate}
}

func (p *Product) Create(name string, description string) (*domain.Product, error) {
    // Generate ID
    id, err := uuid.NewV7()
    if err != nil {
        return nil, multierror.Append(domain.ErrInternal, err)
    }

    // Create entity
    product := &domain.Product{
        ID:          id,
        Name:        name,
        Description: description,
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
    }

    // Validate (pure business logic)
    if err := p.validate.Struct(product); err != nil {
        return nil, multierror.Append(domain.ErrInvalid, err)
    }

    return product, nil
}
```

---

### 🔄 Application Layer (`internal/application`)

**Contains:**

- Use case implementations
- Workflow orchestration
- Parameter objects (DTOs)

**Dependencies:**

- ✅ Multiple repositories (for different aggregates)
- ✅ Domain services
- ✅ External adapters (S3, Redis, Keycloak, etc.)

**Why multiple repos?** Application layer coordinates between different aggregates and external systems.

**Example:**

```go
type ProductImpl struct {
    // Multiple repositories
    productRepo   domain.ProductRepository
    categoryRepo  domain.CategoryRepository
    attributeRepo domain.AttributeRepository

    // Domain service (pure logic)
    productService domain.ProductService

    // External adapters
    s3Client    *s3.Client
    redisClient *redis.Client
}

func ProvideProduct(
    productRepo domain.ProductRepository,
    categoryRepo domain.CategoryRepository,
    attributeRepo domain.AttributeRepository,
    productService domain.ProductService,
    s3Client *s3.Client,
    redisClient *redis.Client,
) *ProductImpl {
    return &ProductImpl{
        productRepo:    productRepo,
        categoryRepo:   categoryRepo,
        attributeRepo:  attributeRepo,
        productService: productService,
        s3Client:       s3Client,
        redisClient:    redisClient,
    }
}

func (p *ProductImpl) Create(ctx context.Context, param CreateProductParam) (*domain.Product, error) {
    // 1. Fetch from multiple repos
    category, err := p.categoryRepo.Get(ctx, param.Data.CategoryID)
    if err != nil {
        return nil, err
    }

    attributes, err := p.attributeRepo.List(ctx, param.Data.AttributeIDs, ...)
    if err != nil {
        return nil, err
    }

    // 2. Call domain service (pure logic)
    product, err := p.productService.Create(param.Data.Name, param.Data.Description)
    if err != nil {
        return nil, err
    }

    // 3. Use adapters
    _, err = p.s3Client.PutObject(ctx, ...)
    if err != nil {
        return nil, err
    }

    // 4. Persist
    err = p.productRepo.Save(ctx, *product)
    if err != nil {
        return nil, err
    }

    // 5. Post-operations
    p.redisClient.Del(ctx, "products:list")

    return product, nil
}
```

---

### 💾 Repository Layer (`internal/infrastructure/repository`)

**Contains:**

- Domain repository implementations
- sqlc integration
- Error mapping

**Dependencies:**

- ✅ Database connection pool (`*pgxpool.Pool`)
- ✅ sqlc generated queries

**Example:**

```go
type ProductRepository struct {
    db      *pgxpool.Pool
    queries *postgres.Queries
}

func (r *ProductRepository) Save(ctx context.Context, entity domain.Product) error {
    err := r.queries.UpsertProduct(ctx, postgres.UpsertProductParams{
        ID:   entity.ID,
        Name: entity.Name,
        // ...
    })

    if err != nil {
        return mapError(err)  // Map to domain errors
    }

    return nil
}
```

---

## Data Flow Examples

### Create Product Flow

```
HTTP Request
    ↓
[Delivery] Parse request, validate params
    ↓
[Application] ProductImpl.Create()
    │
    ├→ CategoryRepo.Get()        ← Fetch related data
    ├→ AttributeRepo.List()      ← Fetch related data
    │
    ├→ ProductService.Create()   ← Pure business logic
    │   ├─ uuid.NewV7()
    │   ├─ validate.Struct()
    │   └─ return Product
    │
    ├→ S3Client.PutObject()      ← Upload images
    ├→ ProductRepo.Save()        ← Persist
    └→ RedisClient.Del()         ← Invalidate cache
    ↓
[Delivery] Return response
```

### Why This Separation?

**Service Layer (Pure Logic):**

```go
// ✅ Can test without database, S3, Redis
func TestProductService_Create(t *testing.T) {
    validate := validator.New()
    service := ProvideProduct(validate)

    product, err := service.Create("Laptop", "Gaming laptop")

    assert.NoError(t, err)
    assert.NotNil(t, product)
}
```

**Application Layer (Orchestration):**

```go
// ✅ Test with mocked dependencies
func TestProductImpl_Create(t *testing.T) {
    mockProductRepo := domain.NewMockProductRepository(t)
    mockCategoryRepo := domain.NewMockCategoryRepository(t)
    mockProductService := domain.NewMockProductService(t)
    // ... setup mocks

    app := ProvideProduct(mockProductRepo, mockCategoryRepo, ..., mockProductService, ...)
    result, err := app.Create(ctx, param)
    // ... assertions
}
```

---

## Testing Strategy

### Service Layer Tests

- **Type:** Unit tests
- **Dependencies:** Only validator (no mocks needed)
- **Pattern:** Table-driven tests
- **Focus:** Business logic correctness

### Application Layer Tests

- **Type:** Unit tests
- **Dependencies:** Mock all repos, services, adapters
- **Pattern:** Table-driven tests with mock setup
- **Focus:** Orchestration logic

### Repository Layer Tests

- **Type:** Integration tests
- **Dependencies:** Test containers (real PostgreSQL)
- **Pattern:** Setup DB → test query → assert
- **Focus:** Data persistence correctness

---

## Key Takeaways

| Layer           | Has Repos?           | Has Adapters?      | Has Services?        | Purpose               |
| --------------- | -------------------- | ------------------ | -------------------- | --------------------- |
| **Domain**      | ❌ (interfaces only) | ❌                 | ❌ (interfaces only) | Define contracts      |
| **Service**     | ❌ NO                | ❌ NO              | ✅ Implements        | Pure business logic   |
| **Application** | ✅ Multiple          | ✅ S3, Redis, etc. | ✅ Injects           | Orchestrate workflows |
| **Repository**  | ✅ Implements        | ❌                 | ❌                   | Data persistence      |

**Remember:**

- Service = Pure logic (testable without infrastructure)
- Application = Coordinator (brings everything together)
- Repository = Data access (maps to domain)
