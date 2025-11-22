# Application Layer Rules

## Overview

The application layer (`internal/application`) orchestrates business workflows, coordinates between domain services and repositories, and handles use case logic. It acts as a bridge between the delivery layer and domain layer.

## Structure

### Application Interfaces

Each aggregate root has:

1. **Interface file** (`<entity>.go`) - Defines use case methods
2. **Implementation file** (`<entity>impl.go`) - Implements the interface
3. **Parameter file** (`<entity>param.go`) - Defines request/response DTOs

**Naming Convention:**

- Interface: `type Product interface`
- Implementation: `type ProductImpl struct`
- Provider function: `func ProvideProduct(...) *ProductImpl`

### Interface Definition

Application interfaces define high-level use cases:

```go
package application

type Product interface {
	Create(context.Context, CreateProductParam) (*domain.Product, error)
	List(context.Context, ListProductParam) (*Pagination[domain.Product], error)
	Get(context.Context, GetProductParam) (*domain.Product, error)
	Update(context.Context, UpdateProductParam) (*Product, error)
	Delete(context.Context, DeleteProductParam) error
}
```

**Conventions:**

- First parameter is always `context.Context`
- Second parameter is a typed param struct (defined in `*param.go`)
- Return domain entities directly (no DTO mapping at this layer)
- Return `error` as last return value
- Use generic `Pagination[T]` for paginated results

### Implementation Structure

```go
package application

import (
	"backend/internal/domain"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/redis/go-redis/v9"
)

type ProductImpl struct {
	// Repositories for multiple aggregates
	productRepo   domain.ProductRepository
	categoryRepo  domain.CategoryRepository
	attributeRepo domain.AttributeRepository

	// Domain services
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

// Ensure implementation satisfies interface
var _ Product = &ProductImpl{}
```

**Conventions:**

- Dependencies injected via constructor (Google Wire)
- Use `Provide<Entity>` naming for provider functions
- Assert interface implementation with `var _ Interface = &Implementation{}`
- Inject multiple repositories (for related aggregates)
- Inject external adapters (S3, Redis, etc.) as needed
- Repository dependencies are interfaces
- Adapter dependencies can be concrete types (e.g., `*s3.Client`, `*redis.Client`)

### Implementation Methods

Application methods orchestrate workflows:

```go
func (p *ProductImpl) Create(ctx context.Context, param CreateProductParam) (*domain.Product, error) {
    // 1. Fetch dependencies from multiple repositories
    category, err := p.categoryRepo.Get(ctx, param.Data.CategoryID)
    if err != nil {
        return nil, err
    }

    var attributes *[]domain.Attribute
    if param.Data.AttributeValueIDs != nil {
        attributes, err = p.attributeRepo.List(ctx, param.Data.AttributeValueIDs, nil, 100, 0)
        if err != nil {
            return nil, err
        }
    }

    // 2. Create entity via domain service (pure business logic)
    product, err := p.productService.Create(
        param.Data.Name,
        param.Data.Description,
        *category,
    )
    if err != nil {
        return nil, err
    }

    // 3. Use external adapters if needed
    // Upload images to S3
    for _, img := range param.Data.Images {
        _, err = p.s3Client.PutObject(ctx, &s3.PutObjectInput{
            Bucket: aws.String("products"),
            Key:    aws.String(img.URL),
        })
        if err != nil {
            return nil, domain.ErrServiceError
        }
    }

    // 4. Persist entity
    err = p.productRepo.Save(ctx, *product)
    if err != nil {
        return nil, err
    }

    // 5. Invalidate cache
    p.redisClient.Del(ctx, "products:list")

    return product, nil
}
```

**Workflow Pattern:**

1. Extract and validate parameters
2. Fetch required dependencies from multiple repositories (related entities)
3. Execute domain service logic (pure business rules, no infrastructure)
4. Use external adapters as needed (S3 uploads, Redis caching, etc.)
5. Persist changes via appropriate repository
6. Perform post-save operations (cache invalidation, events, etc.)
7. Return result

**Error Handling:**

- Don't wrap errors here (already wrapped in domain/service layers)
- Let errors propagate up to delivery layer
- Domain errors will be mapped to HTTP errors in delivery

### Parameter Objects

Parameters are defined in `*param.go` files:

```go
package application

// Request parameters
type CreateProductParam struct {
	Data CreateProductData `binding:"required"`
}

type CreateProductData struct {
	Name        string      `json:"name"        binding:"required"`
	Description string      `json:"description" binding:"required"`
	CategoryID  uuid.UUID   `json:"categoryId"  binding:"required"`
	Price       int64       `json:"price"       binding:"required,gt=0"`
	Images      []ImageData `json:"images"      binding:"required,dive"`
}

type ImageData struct {
	URL   string `json:"url"   binding:"required,url"`
	Order int    `json:"order" binding:"required"`
}

// List/query parameters
type ListProductParam struct {
	PaginationParam
	CategoryIDs *[]uuid.UUID `binding:"omitempty"`
	MinPrice    *int64       `binding:"omitempty"`
	MaxPrice    *int64       `binding:"omitempty"`
	Search      *string      `binding:"omitempty"`
	Deleted     *string      `binding:"omitempty,oneof=exclude only all"`
}
```

**Conventions:**

- Param structs wrap data with clear naming (`Create*Param`, `List*Param`, etc.)
- Use `binding` tags for Gin validation
- Nested data structs for complex payloads
- Pointers for optional query parameters
- Use `PaginationParam` for paginated endpoints
- Use `dive` for slice validation

### Pagination

Common pagination helper:

```go
// internal/application/common.go
type PaginationParam struct {
    Page  *int `form:"page"  binding:"omitempty,gte=1"`
    Limit *int `form:"limit" binding:"omitempty,gte=1,lte=100"`
}

type Pagination[T any] struct {
    Data       []T   `json:"data"`
    Page       int   `json:"page"`
    Limit      int   `json:"limit"`
    TotalItems int   `json:"totalItems"`
    TotalPages int   `json:"totalPages"`
}

func newPagination[T any](data []T, totalItems int, page int, limit int) *Pagination[T] {
    totalPages := (totalItems + limit - 1) / limit
    return &Pagination[T]{
        Data:       data,
        Page:       page,
        Limit:      limit,
        TotalItems: totalItems,
        TotalPages: totalPages,
    }
}
```

### List Implementation Pattern

```go
func (a *AttributeImpl) List(ctx context.Context, param ListAttributesParam) (*Pagination[domain.Attribute], error) {
    // Fetch data
    attributes, err := a.attributeRepo.List(
        ctx,
        param.AttributeIDs,
        param.Search,
        param.Deleted,
        *param.Limit,
        *param.Page,
    )
    if err != nil {
        return nil, err
    }

    // Fetch count
    count, err := a.attributeRepo.Count(
        ctx,
        param.AttributeIDs,
        param.Deleted,
    )
    if err != nil {
        return nil, err
    }

    // Build pagination
    pagination := newPagination(*attributes, *count, *param.Page, *param.Limit)
    return pagination, nil
}
```

## Dependency Injection

Application implementations are wired via Google Wire:

**File:** `internal/di/wire.go`

```go
//go:build wireinject
// +build wireinject

func InitializeApplication() (*Application, error) {
    wire.Build(
        // Repositories
        wire.Bind(new(domain.ProductRepository), new(*postgres.ProductRepository)),

        // Services
        service.ProvideProduct,
        wire.Bind(new(domain.ProductService), new(*service.Product)),

        // Application
        application.ProvideProduct,
        wire.Bind(new(application.Product), new(*application.ProductImpl)),
    )
    return &Application{}, nil
}
```

## Rules Summary

1. ✅ Each entity has interface, implementation, and param files
2. ✅ Use `<Entity>Impl` for implementations
3. ✅ Provider functions named `Provide<Entity>`
4. ✅ First param is `context.Context`, second is typed param struct
5. ✅ Return domain entities directly (no DTO conversion)
6. ✅ Inject multiple repositories for related aggregates
7. ✅ Inject external adapters (S3, Redis) as needed
8. ✅ Orchestrate: fetch deps → service logic → adapters → persist → post-ops
9. ✅ Don't wrap errors (already wrapped in domain/service)
10. ✅ Use `Pagination[T]` for list endpoints
11. ✅ Params use `binding` tags for validation
12. ✅ Wire dependencies via Google Wire
