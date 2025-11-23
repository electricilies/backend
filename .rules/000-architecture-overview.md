# DDD Architecture: Layer Responsibilities

## Five-Layer Structure

1. **Domain** - Interfaces, models, errors (no implementations)
2. **Service** - Pure business logic (ONLY validator dependency)
3. **Application** - Orchestration (repos + services + adapters)
4. **Repository** - Data persistence (PostgreSQL + sqlc)
5. **Delivery** - HTTP handlers (Gin)

## Critical Layer Rules

| Layer | Dependencies | Purpose | Testing |
|-------|-------------|---------|----------|
| **Domain** | None | Define contracts | N/A (interfaces) |
| **Service** | `*validator.Validate` ONLY | Pure business logic | Unit (no mocks) |
| **Application** | Multiple repos + services + adapters | Coordinate workflows | Unit (all mocked) |
| **Repository** | `*pgxpool.Pool` + sqlc | Persist data | Integration (test containers) |
| **Delivery** | Application interfaces | HTTP/REST API | E2E |

## Service vs Application

**Service (Pure Logic):**
```go
type Product struct {
    validate *validator.Validate  // ONLY dependency
}

func (s *Product) Create(name string) (*domain.Product, error) {
    id, _ := uuid.NewV7()
    product := &domain.Product{ID: id, Name: name}
    return product, s.validate.Struct(product)
}
```

**Application (Orchestration):**
```go
type ProductImpl struct {
    productRepo    domain.ProductRepository
    categoryRepo   domain.CategoryRepository  // Multiple repos
    productService domain.ProductService      // Inject service
    s3Client       *s3.Client                 // External adapters
    redisClient    *redis.Client
}

func (a *ProductImpl) Create(ctx context.Context, param CreateParam) (*domain.Product, error) {
    // 1. Fetch dependencies from multiple repos
    category, _ := a.categoryRepo.Get(ctx, param.CategoryID)
    
    // 2. Business logic via service
    product, _ := a.productService.Create(param.Name, param.Description)
    
    // 3. External operations
    a.s3Client.PutObject(ctx, ...)
    
    // 4. Persist
    a.productRepo.Save(ctx, *product)
    
    // 5. Cache invalidation
    a.redisClient.Del(ctx, "products:list")
    
    return product, nil
}
```

## Data Flow Pattern

```
[Delivery] → [Application] → [Service] → [Domain]
                    ↓             ↓
              [Repository]    [Validator]
                    ↓
               [Database]
```

**Application Orchestration Flow:**
1. Fetch dependencies (multiple repos)
2. Execute business logic (service)
3. External operations (S3, Redis, etc.)
4. Persist changes (repo)
5. Post-operations (cache, events)

## Key Principles

✅ **Service = Pure Logic** (testable without infrastructure)  
✅ **Application = Coordinator** (brings everything together)  
✅ **Repository = Data Access** (maps domain ↔ database)  
✅ **Domain = Contracts** (interfaces, no implementations)  

❌ Services NEVER have repos/adapters  
❌ Repositories NEVER contain business logic  
❌ Application NEVER duplicates service logic
