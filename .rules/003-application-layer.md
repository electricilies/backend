# Application Layer (`internal/application`)

## Purpose

Orchestrate workflows: coordinate repos + services + adapters

## File Structure

- `<entity>.go` - Interface definition
- `<entity>impl.go` - Implementation
- `<entity>requestdto.go` - Request DTOs
- `<entity>responsedto.go` - Response DTOs

## Implementation Pattern

```go
// Interface
type Product interface {
    Create(context.Context, CreateProductParam) (*domain.Product, error)
    List(context.Context, ListProductParam) (*Pagination[domain.Product], error)
    Get(context.Context, GetProductParam) (*domain.Product, error)
    Update(context.Context, UpdateProductParam) (*domain.Product, error)
    Delete(context.Context, DeleteProductParam) error
}

// Implementation
type ProductImpl struct {
    // Multiple repositories
    productRepo   domain.ProductRepository
    categoryRepo  domain.CategoryRepository

    // Domain service
    productService domain.ProductService

    // External adapters
    s3Client      *s3.Client
    redisClient   *redis.Client
}

func ProvideProduct(
    productRepo domain.ProductRepository,
    categoryRepo domain.CategoryRepository,
    productService domain.ProductService,
    s3Client *s3.Client,
    redisClient *redis.Client,
) *ProductImpl {
    return &ProductImpl{
        productRepo:    productRepo,
        categoryRepo:   categoryRepo,
        productService: productService,
        s3Client:       s3Client,
        redisClient:    redisClient,
    }
}

var _ Product = &ProductImpl{}  // Assert interface implementation
```

## Orchestration Workflow

```go
func (a *ProductImpl) Create(ctx context.Context, param CreateProductParam) (*domain.Product, error) {
    // 1. Fetch dependencies from multiple repos
    category, err := a.categoryRepo.Get(ctx, param.Data.CategoryID)
    if err != nil {
        return nil, err
    }

    // 2. Execute business logic via service
    product, err := a.productService.Create(
        param.Data.Name,
        param.Data.Description,
        *category,
    )
    if err != nil {
        return nil, err
    }

    // 3. External operations
    for _, img := range param.Data.Images {
        _, err = a.s3Client.PutObject(ctx, &s3.PutObjectInput{
            Bucket: aws.String("products"),
            Key:    aws.String(img.URL),
        })
        if err != nil {
            return nil, domain.ErrServiceError
        }
    }

    // 4. Persist
    err = a.productRepo.Save(ctx, *product)
    if err != nil {
        return nil, err
    }

    // 5. Post-operations (cache invalidation, events)
    a.redisClient.Del(ctx, "products:list")

    return product, nil
}
```

## Parameter Objects

```go
// Create parameters
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

// List parameters
type ListProductParam struct {
    PaginationParam
    CategoryIDs *[]uuid.UUID `binding:"omitempty"`
    MinPrice    *int64       `binding:"omitempty"`
    MaxPrice    *int64       `binding:"omitempty"`
    Search      *string      `binding:"omitempty"`
    Deleted     *string      `binding:"omitempty,oneof=exclude only all"`
}
```

## Pagination

```go
// common.go
type PaginationParam struct {
    Page  *int `form:"page"  binding:"omitempty,gte=1"`
    Limit *int `form:"limit" binding:"omitempty,gte=1,lte=100"`
}

type Pagination[T any] struct {
    Data       []T `json:"data"`
    Page       int `json:"page"`
    Limit      int `json:"limit"`
    TotalItems int `json:"totalItems"`
    TotalPages int `json:"totalPages"`
}

func newPagination[T any](data []T, totalItems, page, limit int) *Pagination[T] {
    return &Pagination[T]{
        Data:       data,
        Page:       page,
        Limit:      limit,
        TotalItems: totalItems,
        TotalPages: (totalItems + limit - 1) / limit,
    }
}
```

## List Implementation

```go
func (a *ProductImpl) List(ctx context.Context, param ListProductParam) (*Pagination[domain.Product], error) {
    products, err := a.productRepo.List(
        ctx,
        param.CategoryIDs,
        param.Search,
        domain.DeletedExclude,
        *param.Limit,
        (*param.Page-1) * *param.Limit,  // offset
    )
    if err != nil {
        return nil, err
    }

    count, err := a.productRepo.Count(ctx, param.CategoryIDs, domain.DeletedExclude)
    if err != nil {
        return nil, err
    }

    return newPagination(*products, *count, *param.Page, *param.Limit), nil
}
```

## Dependency Injection (Wire)

**File:** `internal/di/wire.go`

```go
//go:build wireinject

func InitializeApp() (*App, error) {
    wire.Build(
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

## Quick Rules

1. ✅ Three files: interface, impl, param
2. ✅ Provider: `Provide<Entity>(...) *<Entity>Impl`
3. ✅ First param: `context.Context`, second: typed param
4. ✅ Inject multiple repos + services + adapters
5. ✅ Return domain entities (no DTO conversion here)
6. ✅ DON'T wrap errors (already wrapped in service)
7. ✅ Workflow: fetch deps → service → adapters → persist → post-ops
8. ✅ Use `Pagination[T]` for list endpoints
9. ✅ Params use `binding` tags for Gin validation
10. ❌ No business logic duplication from service
