# Testing Rules

## Overview

This project uses a comprehensive testing strategy with:

- **Unit tests** for services and application logic
- **Integration tests** for repositories
- **Mock generation** via mockery
- **Test containers** for PostgreSQL, Redis, MinIO, Keycloak

## Test Organization

```
internal/
├── domain/
│   └── *_repository_mock.go   # Generated mocks
├── service/
│   └── *_test.go               # Service unit tests
├── application/
│   └── *_test.go               # Application unit tests
└── infrastructure/
    └── repository/
        └── *_test.go           # Repository integration tests
```

## Mock Generation

### Configuration

**File:** `.mockery.yml`

```yaml
with-expecter: true
dir: "internal/domain"
outpkg: "domain"
filename: "{{.InterfaceName | lower}}_mock.go"
packages:
  github.com/yourusername/backend/internal/domain:
    interfaces:
      ProductRepository:
      ProductService:
      CategoryRepository:
      # ... other interfaces
```

### Generate Mocks

```bash
mockery
```

### Using Mocks

```go
import (
    "testing"
    "backend/internal/domain"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

func TestProductImpl_Create(t *testing.T) {
    // Setup mocks
    mockRepo := domain.NewMockProductRepository(t)
    mockService := domain.NewMockProductService(t)

    // Define expectations
    product := &domain.Product{ID: uuid.New(), Name: "Test"}
    mockService.EXPECT().
        Create("Test Product", "Description", mock.Anything).
        Return(product, nil).
        Once()

    mockRepo.EXPECT().
        Save(mock.Anything, *product).
        Return(nil).
        Once()

    // Test
    app := application.ProvideProduct(mockRepo, mockService)
    result, err := app.Create(context.Background(), param)

    // Assert
    assert.NoError(t, err)
    assert.Equal(t, "Test", result.Name)
}
```

## Unit Tests

### Service Layer Tests

Test business logic without database:

```go
func TestProduct_Create(t *testing.T) {
    validate := validator.New()
    service := ProvideProduct(validate)

    category := domain.Category{
        ID:   uuid.New(),
        Name: "Electronics",
    }

    product, err := service.Create(
        "Gaming Laptop",
        "High-performance gaming laptop",
        category,
    )

    assert.NoError(t, err)
    assert.NotNil(t, product)
    assert.NotEqual(t, uuid.Nil, product.ID)
    assert.Equal(t, "Gaming Laptop", product.Name)
    assert.Equal(t, 0, product.ViewsCount)
}

func TestProduct_Create_ValidationError(t *testing.T) {
    validate := validator.New()
    service := ProvideProduct(validate)

    category := domain.Category{ID: uuid.New(), Name: "Electronics"}

    // Name too short (< 3 chars)
    product, err := service.Create("AB", "Description", category)

    assert.Error(t, err)
    assert.Nil(t, product)
    assert.ErrorIs(t, err, domain.ErrInvalid)
}
```

### Application Layer Tests

Test orchestration with mocks using table-driven tests:

```go
func TestProductImpl_Create(t *testing.T) {
    tests := []struct {
        name          string
        param         application.CreateProductParam
        setupMocks    func(*domain.MockProductRepository, *domain.MockProductService, *domain.MockCategoryRepository)
        wantErr       bool
        wantErrType   error
    }{
        {
            name: "successful creation",
            param: application.CreateProductParam{
                Data: application.CreateProductData{
                    Name:        "Laptop",
                    Description: "Gaming laptop",
                    CategoryID:  uuid.New(),
                },
            },
            setupMocks: func(mockProductRepo *domain.MockProductRepository, mockProductService *domain.MockProductService, mockCategoryRepo *domain.MockCategoryRepository) {
                category := &domain.Category{ID: uuid.New(), Name: "Electronics"}
                product := &domain.Product{
                    ID:          uuid.New(),
                    Name:        "Laptop",
                    Description: "Gaming laptop",
                    Category:    category,
                }

                mockCategoryRepo.EXPECT().
                    Get(mock.Anything, mock.Anything).
                    Return(category, nil)

                mockProductService.EXPECT().
                    Create("Laptop", "Gaming laptop", *category).
                    Return(product, nil)

                mockProductRepo.EXPECT().
                    Save(mock.Anything, *product).
                    Return(nil)
            },
            wantErr: false,
        },
        {
            name: "category not found",
            param: application.CreateProductParam{
                Data: application.CreateProductData{
                    Name:        "Laptop",
                    Description: "Gaming laptop",
                    CategoryID:  uuid.New(),
                },
            },
            setupMocks: func(mockProductRepo *domain.MockProductRepository, mockProductService *domain.MockProductService, mockCategoryRepo *domain.MockCategoryRepository) {
                mockCategoryRepo.EXPECT().
                    Get(mock.Anything, mock.Anything).
                    Return(nil, domain.ErrNotFound)
            },
            wantErr:     true,
            wantErrType: domain.ErrNotFound,
        },
        {
            name: "validation error from service",
            param: application.CreateProductParam{
                Data: application.CreateProductData{
                    Name:        "AB", // Too short
                    Description: "Gaming laptop",
                    CategoryID:  uuid.New(),
                },
            },
            setupMocks: func(mockProductRepo *domain.MockProductRepository, mockProductService *domain.MockProductService, mockCategoryRepo *domain.MockCategoryRepository) {
                category := &domain.Category{ID: uuid.New(), Name: "Electronics"}

                mockCategoryRepo.EXPECT().
                    Get(mock.Anything, mock.Anything).
                    Return(category, nil)

                mockProductService.EXPECT().
                    Create("AB", "Gaming laptop", *category).
                    Return(nil, domain.ErrInvalid)
            },
            wantErr:     true,
            wantErrType: domain.ErrInvalid,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Setup mocks
            mockProductRepo := domain.NewMockProductRepository(t)
            mockProductService := domain.NewMockProductService(t)
            mockCategoryRepo := domain.NewMockCategoryRepository(t)

            // Configure mock expectations
            tt.setupMocks(mockProductRepo, mockProductService, mockCategoryRepo)

            // Create application instance
            app := application.ProvideProduct(
                mockProductRepo,
                mockCategoryRepo,
                nil, // Other dependencies
                mockProductService,
                nil, // S3 client
                nil, // Redis client
            )

            // Execute
            result, err := app.Create(context.Background(), tt.param)

            // Assert
            if tt.wantErr {
                assert.Error(t, err)
                if tt.wantErrType != nil {
                    assert.ErrorIs(t, err, tt.wantErrType)
                }
                assert.Nil(t, result)
            } else {
                assert.NoError(t, err)
                assert.NotNil(t, result)
            }
        })
    }
}
```

## Integration Tests

### Test Containers Setup

```go
import (
    "testing"
    "context"
    "github.com/testcontainers/testcontainers-go/modules/postgres"
    "github.com/jackc/pgx/v5/pgxpool"
)

func setupTestDB(t *testing.T) (*pgxpool.Pool, func()) {
    ctx := context.Background()

    // Start PostgreSQL container
    pgContainer, err := postgres.RunContainer(ctx,
        testcontainers.WithImage("postgres:16-alpine"),
        postgres.WithDatabase("testdb"),
        postgres.WithUsername("test"),
        postgres.WithPassword("test"),
        postgres.WithInitScripts("../../database/schema.sql"),
    )
    require.NoError(t, err)

    // Get connection string
    connStr, err := pgContainer.ConnectionString(ctx)
    require.NoError(t, err)

    // Create pool
    pool, err := pgxpool.New(ctx, connStr)
    require.NoError(t, err)

    // Cleanup function
    cleanup := func() {
        pool.Close()
        pgContainer.Terminate(ctx)
    }

    return pool, cleanup
}
```

### Repository Integration Tests

```go
func TestProductRepository_Save(t *testing.T) {
    db, cleanup := setupTestDB(t)
    defer cleanup()

    repo := ProvideProductRepository(db)
    ctx := context.Background()

    product := domain.Product{
        ID:          uuid.New(),
        Name:        "Test Product",
        Description: "Test Description",
        Price:       10000,
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
    }

    // Save
    err := repo.Save(ctx, product)
    assert.NoError(t, err)

    // Retrieve
    retrieved, err := repo.Get(ctx, product.ID)
    assert.NoError(t, err)
    assert.Equal(t, product.Name, retrieved.Name)
}

func TestProductRepository_List(t *testing.T) {
    db, cleanup := setupTestDB(t)
    defer cleanup()

    repo := ProvideProductRepository(db)
    ctx := context.Background()

    // Seed data
    products := []domain.Product{
        {ID: uuid.New(), Name: "Product 1", Price: 1000},
        {ID: uuid.New(), Name: "Product 2", Price: 2000},
        {ID: uuid.New(), Name: "Product 3", Price: 3000},
    }

    for _, p := range products {
        err := repo.Save(ctx, p)
        require.NoError(t, err)
    }

    // Test list
    results, err := repo.List(ctx, nil, nil, domain.DeletedExclude, 10, 0)
    assert.NoError(t, err)
    assert.Len(t, *results, 3)
}
```

## Table-Driven Tests

Use table-driven tests for multiple scenarios:

```go
func TestProduct_Create(t *testing.T) {
    validate := validator.New()
    service := ProvideProduct(validate)

    tests := []struct {
        name        string
        productName string
        description string
        wantErr     bool
        errType     error
    }{
        {
            name:        "valid product",
            productName: "Valid Name",
            description: "Valid description with enough length",
            wantErr:     false,
        },
        {
            name:        "name too short",
            productName: "AB",
            description: "Valid description",
            wantErr:     true,
            errType:     domain.ErrInvalid,
        },
        {
            name:        "description too short",
            productName: "Valid Name",
            description: "Short",
            wantErr:     true,
            errType:     domain.ErrInvalid,
        },
    }

    category := domain.Category{ID: uuid.New(), Name: "Test"}

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            product, err := service.Create(tt.productName, tt.description, category)

            if tt.wantErr {
                assert.Error(t, err)
                if tt.errType != nil {
                    assert.ErrorIs(t, err, tt.errType)
                }
                assert.Nil(t, product)
            } else {
                assert.NoError(t, err)
                assert.NotNil(t, product)
            }
        })
    }
}
```

## Test Helpers

Create shared test utilities:

```go
// internal/test/helpers.go
package test

func CreateTestProduct(t *testing.T) domain.Product {
	t.Helper()
	return domain.Product{
		ID:          uuid.New(),
		Name:        "Test Product",
		Description: "Test Description",
		Price:       10000,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func CreateTestCategory(t *testing.T) domain.Category {
	t.Helper()
	return domain.Category{
		ID:   uuid.New(),
		Name: "Test Category",
	}
}
```

## Running Tests

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run with race detector
go test -race ./...

# Run specific package
go test ./internal/service/...

# Run specific test
go test -run TestProduct_Create ./internal/service

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## CI/CD Integration

**File:** `.github/workflows/main.yaml`

```yaml
name: Tests
on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v4
        with:
          go-version: "1.25.3"

      - name: Run tests
        run: go test -race -coverprofile=coverage.out ./...

      - name: Upload coverage
        uses: codecov/codecov-action@v3
        with:
          file: ./coverage.out
```

## Rules Summary

1. ✅ Generate mocks with mockery for all repository/service interfaces
2. ✅ Use testify for assertions and mocks
3. ✅ **Prefer table-driven tests** for multiple scenarios
4. ✅ Unit test services without database (only use validator)
5. ✅ Unit test application with mocked repositories and services
6. ✅ Application tests mock multiple repositories + adapters
7. ✅ Integration test repositories with test containers
8. ✅ Create test helpers for common setup
9. ✅ Run tests with `-race` flag
10. ✅ Aim for high coverage (>80%)
11. ✅ Use `t.Helper()` in test utility functions
12. ✅ Each table-driven test case has descriptive name
