# Testing Strategy

## Layers

- **Service/Domain:** Unit test, pure logic, validator only. Use table-driven tests covering Normal, Abnormal, and Boundary cases.
- **Application:** Integration test, use testcontainers. Test full lifecycles (Create->Get->Update->Delete) and side effects (Cache, DB).
- **Repository:** Unit test, mock repository with Mockery.

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

- No mocks, just validator and logic.
- **Strictly** use table-driven tests.
- Cover all paths: Normal, Abnormal (Errors), Boundary (Edge cases).

## Application Tests

- Use testcontainers for external dependencies (Redis, S3, DB).
- **Lifecycle style:** Test sequences of operations (Create -> Get -> Update -> List -> Delete) to verify state persistence and cache invalidation.
- Use `testify/suite`.

## Repository Tests

- Use Mockery-generated mocks for repository interface.
- Test repository logic in isolation.

## Helpers

- Use `t.Helper()` in test utilities.

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

- Mandatory for Service/Domain logic.
- Define explicit test cases for:
  - **Normal:** Standard valid inputs.
  - **Abnormal:** Invalid inputs, error conditions.
  - **Boundary:** Edge values (min/max, empty, limits).
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

- Use for Application layer integration tests.
- Verify state changes across multiple operations.
- Check side effects like Cache hits/misses/invalidation.
- Example:
  ```go
  s.Run("Create resource", func() { ... })
  s.Run("Get resource (cache miss)", func() { ... })
  s.Run("Get resource (cache hit)", func() { ... })
  s.Run("Update resource (cache invalidation)", func() { ... })
  ```

### Testify Suite

- Use `testify/suite` for integration tests requiring setup/teardown
- Especially useful for tests with testcontainers
- Example structure (see `test/integration/application/attribute_test.go`)

### Validator Registration

- Register custom validators in `internal/client/validate.go`
- Call `domain.RegisterAttributeValidators(validate)` for attribute validators
- Ensure all custom validation tags are registered before use

## Rules

1. Service/Domain: Table-driven unit tests (Normal, Abnormal, Boundary).
2. Application: Integration tests with Testcontainers & Testify Suite.
3. Application: Test lifecycles (Create->Get->Update->Delete) and cache side-effects.
4. Repository: Unit test with Mockery.
5. Coverage: > 80%, covering all logical paths.
6. Use `t.Parallel()` where possible.
7. Run with `-race`.
8. Register custom validators in `internal/client/validate.go`.
