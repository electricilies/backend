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

## Rules

1. Service: unit test, validator only
2. Application: integration test, testcontainers for external deps
3. Repository: unit test, mock repo with Mockery
4. Table-driven tests for scenarios
5. Mock with mockery + testify
6. Use `t.Helper()` in utilities
7. Run with `-race`
8. > 80% coverage
9. Descriptive test names
10. Setup/teardown for integration
