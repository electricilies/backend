# Testing Strategy

## Layers

- **Service:** Unit test, pure logic, validator only.
- **Application:** Integration test, use testcontainers for Redis, S3, DB, etc.
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
- Use table-driven tests for scenarios.

## Application Tests

- Use testcontainers for external dependencies (Redis, S3, DB).
- Integration style: spin up containers, test real interactions.

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

### Table-Driven Tests
- Use table-driven pattern for multiple test scenarios
- Structure:
  ```go
  tests := []struct {
      name        string
      input       InputType
      expected    OutputType
      expectError bool
  }{
      {name: "valid case", input: validInput, expected: validOutput},
      {name: "error case", input: invalidInput, expectError: true},
  }
  
  for _, tt := range tests {
      t.Run(tt.name, func(t *testing.T) {
          // test logic
      })
  }
  ```

### Testify Suite
- Use `testify/suite` for integration tests requiring setup/teardown
- Especially useful for tests with testcontainers
- Example structure (see `test/integration/application/product_integration_test.go.bak`)

### Validator Registration
- Register custom validators in `internal/client/validate.go`
- Call `domain.RegisterAttributeValidators(validate)` for attribute validators
- Ensure all custom validation tags are registered before use

## Rules

1. Service: unit test, validator only
2. Application: integration test, testcontainers for external deps, use testify suite
3. Repository: unit test, mock repo with Mockery
4. Table-driven tests for multiple scenarios
5. Mock with mockery + testify
6. Use `t.Helper()` in utilities
7. Use `t.Parallel()` when tests can run concurrently
8. Run with `-race`
9. > 80% coverage
10. Descriptive test names
11. Setup/teardown for integration via testify suite
12. Register all custom validators in `internal/client/validate.go`
