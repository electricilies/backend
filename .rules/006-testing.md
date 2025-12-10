# Testing Strategy

## Layers

- **Service/Domain:** Unit test, pure logic, validator only. Use table-driven tests covering Normal, Abnormal, and Boundary cases
- **Application:** Integration test, use testcontainers. Test full lifecycles (Create->Get->Update->Delete) and side effects (Cache, DB)
- **Repository:** Unit test, mock repository with Mockery

## Mockery Usage

- **Config:** `.mockery.yml`
- **Generate:** `mockery`
- **Output:** `internal/domain/*_repository_mock.go`
- **Example:**
  ```go
  mockRepo := domain.NewMockProductRepository(t)
  mockRepo.EXPECT().Get(ctx, id).Return(&product, nil).Once()
  // Use mockRepo in repository tests
  ```

## Service Tests

- No mocks, just validator and logic
- **Strictly** use table-driven tests
- Cover all paths: Normal, Abnormal (Errors), Boundary (Edge cases)

## Application Tests

- Use testcontainers for external dependencies (Redis, S3, DB)
- **Lifecycle style:** Test sequences of operations (Create -> Get -> Update -> List -> Delete) to verify state persistence and cache invalidation
- Use `testify/suite`

## Repository Tests

- Use Mockery-generated mocks for repository interface
- Test repository logic in isolation

## Helpers

- Use `t.Helper()` in test utilities

## Running Tests

```bash
go test ./...        # All tests
go test -cover ./... # Coverage
go test -race ./...  # Race detector
```

## Best Practices

### Parallel Testing

- Use `t.Parallel()` for tests that can run concurrently
- Mark tests parallel if they don't share state or require sequential execution
- Example:
  ```go
  func TestSomething(t *testing.T) {
      t.Parallel()
      // test code
  }
  ```

### Table-Driven Tests (Unit)

- Mandatory for Service/Domain logic
- Define explicit test cases for:
  - **Normal:** Standard valid inputs
  - **Abnormal:** Invalid inputs, error conditions
  - **Boundary:** Edge values (min/max, empty, limits)
- Structure:

  ```go
  tests := []struct {
      name        string
      input       InputType
      expected    OutputType
      expectError bool
  }{
      {name: "Normal: valid case", input: validInput, expected: validOutput},
      {name: "Boundary: max value", input: maxInput, expected: maxOutput},
      {name: "Abnormal: invalid input", input: invalidInput, expectError: true},
  }

  for _, tt := range tests {
      t.Run(tt.name, func(t *testing.T) {
          // test logic
      })
  }
  ```

### Lifecycle Testing (Integration)

- Use for Application layer integration tests
- Verify state changes across multiple operations
- Check side effects like Cache hits/misses/invalidation
- Example:
  ```go
  s.Run("Create resource", func() { ... })
  s.Run("Get resource (cache miss)", func() { ... })
  s.Run("Get resource (cache hit)", func() { ... })
  s.Run("Update resource (cache invalidation)", func() { ... })
  ```
- Use `s.Require().NoError(err)` and `Error()` only for critical assertions, not for every check

### Testify Suite

- Use `testify/suite` for integration tests requiring setup/teardown
- Especially useful for tests with testcontainers
- Example structure (see `test/integration/application/attribute_test.go`)

### Validator Registration

- Register custom validators in `internal/client/validate.go`
- Call `domain.RegisterAttributeValidators(validate)` for attribute validators
- Ensure all custom validation tags are registered before use

## Reporting & Defects

- For documenting test cases and results, follow the template in ./docs/testing/whitebox-template-guidance.md
- If any defect is found during testing, record it in ./docs/testing/defect-log.md
- The boundary test case may appear in unit test only, hardly in integration test

## Seeding

- In integration tests, use seeding setting to prepare initial data state
  ```go
  func (s *AttributeTestSuite) newContainersConfig() *component.ContainersConfig {
    containersConfig := component.NewContainersConfig(&component.NewContainersConfigParam{
      DBEnabled:    true,
    })
    containersConfig.DB.Seed = true
    return containersConfig
  }
  ```
- The seed data include some
  - Product:
    ```sql
    INSERT INTO products (id, name, price, rating, description, category_id) VALUES
      ('00000000-0000-7000-0000-000278469304', 'Điện thoại Masstel Izi 56 4G (LTE) Gọi HD Call ,Pin khủng ,loa lớn - Hàng Chính Hãng', 499000, 0, ..., '00000000-0000-7000-0000-000000001796'),
      ('00000000-0000-7000-0000-000278469345', 'Điện thoại ZTE Family 4GB/128GB, Màn OLED Full HD+, Dimensity 700, Kháng nước IP67, Sạc 22,5W - Mới nguyên seal - Hàng nhập khẩu nhật', 2599000, 0, ..., '00000000-0000-7000-0000-000000001795'),
      ...
    ```
  - Product Variants:
    ```sql
    INSERT INTO product_variants (id, sku, price, quantity, purchase_count, product_id) VALUES
      ('00000000-0000-7000-0000-000278469308', '4868714459472', 499000, 1000, 100, '00000000-0000-7000-0000-000278469304'),
      ('00000000-0000-7000-0000-000278469306', '5829048963687', 499000, 1000, 100, '00000000-0000-7000-0000-000278469304'),
      ('00000000-0000-7000-0000-000278469310', '4493764476839', 499000, 1000, 100, '00000000-0000-7000-0000-000278469304'),
      ('00000000-0000-7000-0000-000278469347', '4432167913574', 2599000, 1000, 100, '00000000-0000-7000-0000-000278469345'),
      ('00000000-0000-7000-0000-000278469350', '7717071852400', 2599000, 1000, 100, '00000000-0000-7000-0000-000278469345'),
    ...
    ```
  - Category: `00000000-0000-7000-0000-000000001796`
  - Cart:
    ```sql
    INSERT INTO carts (id, user_id) VALUES
      ('00000000-0000-7000-0000-000000000001', '00000000-0000-7000-0000-000000000003') -- This is the only user that has cart, name is 'customer'
    ```
  - Cart items:
    ```sql
    INSERT INTO cart_items (id, cart_id, product_variant_id, quantity) VALUES
      ('00000000-0000-7000-0000-000000000001', '00000000-0000-7000-0000-000000000001', '00000000-0000-7000-0000-000278469308', 10),
      ('00000000-0000-7000-0000-000000000002', '00000000-0000-7000-0000-000000000001', '00000000-0000-7000-0000-000278620836', 1)
    ```
- Do not read the seed.sql file, which is very large!

## Rules

1. Service/Domain: Table-driven unit tests (Normal, Abnormal, Boundary)
2. Application: Integration tests with Testcontainers & Testify Suite
3. Application: Test lifecycles (Create->Get->Update->Delete) and cache side-effects
4. Repository: Unit test with Mockery
5. Coverage: > 80%, covering all logical paths
6. Use `t.Parallel()` where possible
7. Run with `-race`
8. Register custom validators in `internal/client/validate.go`
