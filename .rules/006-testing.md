# Testing Strategy

## Overview

- **Service:** Unit tests (no mocks, only validator)
- **Application:** Unit tests (mock all deps)
- **Repository:** Integration tests (test containers)

## Mock Generation

**Config:** `.mockery.yml`

**Generate:** `mockery`

**Output:** `internal/domain/*_repository_mock.go`

```go
mockRepo := domain.NewMockProductRepository(t)
mockService := domain.NewMockProductService(t)

mockRepo.EXPECT().Get(ctx, id).Return(&product, nil).Once()
```

## Service Tests (Unit)

**No mocks needed - pure logic testing**

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

**Table-Driven Tests:**

```go
func TestProduct_Create(t *testing.T) {
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
            description: "Valid description",
            wantErr:     false,
        },
        {
            name:        "name too short",
            productName: "AB",
            description: "Valid description",
            wantErr:     true,
            errType:     domain.ErrInvalid,
        },
    }
    
    validate := validator.New()
    service := ProvideProduct(validate)
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

## Application Tests (Unit)

**Mock all dependencies**

```go
func TestProductImpl_Create(t *testing.T) {
    tests := []struct {
        name        string
        param       application.CreateProductParam
        setupMocks  func(*domain.MockProductRepository, *domain.MockProductService, *domain.MockCategoryRepository)
        wantErr     bool
        wantErrType error
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
                product := &domain.Product{ID: uuid.New(), Name: "Laptop"}
                
                mockCategoryRepo.EXPECT().Get(mock.Anything, mock.Anything).Return(category, nil)
                mockProductService.EXPECT().Create("Laptop", "Gaming laptop", *category).Return(product, nil)
                mockProductRepo.EXPECT().Save(mock.Anything, *product).Return(nil)
            },
            wantErr: false,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            mockProductRepo := domain.NewMockProductRepository(t)
            mockProductService := domain.NewMockProductService(t)
            mockCategoryRepo := domain.NewMockCategoryRepository(t)
            
            tt.setupMocks(mockProductRepo, mockProductService, mockCategoryRepo)
            
            app := application.ProvideProduct(
                mockProductRepo,
                mockCategoryRepo,
                mockProductService,
                nil, // S3
                nil, // Redis
            )
            
            result, err := app.Create(context.Background(), tt.param)
            
            if tt.wantErr {
                assert.Error(t, err)
                if tt.wantErrType != nil {
                    assert.ErrorIs(t, err, tt.wantErrType)
                }
            } else {
                assert.NoError(t, err)
                assert.NotNil(t, result)
            }
        })
    }
}
```

## Repository Tests (Integration)

**Use testcontainers for real database**

```go
func setupTestDB(t *testing.T) (*pgxpool.Pool, func()) {
    ctx := context.Background()
    
    pgContainer, err := postgres.RunContainer(ctx,
        testcontainers.WithImage("postgres:16-alpine"),
        postgres.WithDatabase("testdb"),
        postgres.WithUsername("test"),
        postgres.WithPassword("test"),
        postgres.WithInitScripts("../../database/schema.sql"),
    )
    require.NoError(t, err)
    
    connStr, err := pgContainer.ConnectionString(ctx)
    require.NoError(t, err)
    
    pool, err := pgxpool.New(ctx, connStr)
    require.NoError(t, err)
    
    cleanup := func() {
        pool.Close()
        pgContainer.Terminate(ctx)
    }
    
    return pool, cleanup
}

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
    
    err := repo.Save(ctx, product)
    assert.NoError(t, err)
    
    retrieved, err := repo.Get(ctx, product.ID)
    assert.NoError(t, err)
    assert.Equal(t, product.Name, retrieved.Name)
}
```

## Test Helpers

```go
// internal/test/helpers.go
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
```

## Running Tests

```bash
# All tests
go test ./...

# With coverage
go test -cover ./...

# With race detector
go test -race ./...

# Specific package
go test ./internal/service/...

# Coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Quick Rules

1. ✅ Service: unit test with validator only
2. ✅ Application: unit test with all mocks
3. ✅ Repository: integration test with testcontainers
4. ✅ Use table-driven tests for multiple scenarios
5. ✅ Mock with mockery + testify
6. ✅ Use `t.Helper()` in test utilities
7. ✅ Run with `-race` flag
8. ✅ Aim for >80% coverage
9. ✅ Each test case has descriptive name
10. ✅ Setup/teardown for integration tests
