# Unit Testing Strategy

## Scope

- **Service/Domain:** Pure logic, validator only.
- **Repository:** Unit test, mock repository with Mockery.

## Best Practices

### 1. Table-Driven Tests

- **Mandatory** for Service/Domain logic.
- Define explicit test cases for:
  - **Normal:** Standard valid inputs.
  - **Abnormal:** Invalid inputs, error conditions.
  - **Boundary:** Edge values (min/max, empty, limits).
- Structure:

  ```go
  func (s *AttributeTestSuite) TestNewAttributeBoundaryValues() {
      s.T().Parallel()
      testcases := []struct {
          name      string
          code      string
          attrName  string
          expectErr bool
      }{
          {
              name:      "code length 1 (min - 1)",
              code:      "a",
              attrName:  "ValidName",
              expectErr: true,
          },
          {
              name:      "code length 2 (min)",
              code:      "ab",
              attrName:  "ValidName",
              expectErr: false,
          },
          // ...
      }

      for _, tc := range testcases {
          s.Run(tc.name, func() {
              // Test logic
              // Assertions
          })
      }
  }
  ```

### 2. Parallel Testing

- Use `t.Parallel()` for tests that can run concurrently.
- Mark tests parallel if they don't share state or require sequential execution.

### 3. Mockery Usage

- **Config:** `.mockery.yml`
- **Generate:** `mockery`
- **Output:** `internal/domain/*_repository_mock.go`
- **Example:**
  ```go
  mockRepo := domain.NewMockProductRepository(t)
  mockRepo.EXPECT().Get(ctx, id).Return(&product, nil).Once()
  // Use mockRepo in repository tests
  ```

### 4. Validator Registration

- Register custom validators in `internal/client/validate.go`.
- Call `domain.RegisterAttributeValidators(validate)` for attribute validators.
- Ensure all custom validation tags are registered before use.

### 5. Running Tests

```bash
go test ./...        # All tests
go test -cover ./... # Coverage
go test -race ./...  # Race detector
```

### 6. Helpers

- Use `t.Helper()` in test utility functions to mark them as helpers, ensuring failure logs point to the caller.

## Reporting & Defects

- **Documentation:** Follow the template in `./docs/testing/whitebox-template-guidance.md`.
- **Defect Log:** Record defects in `./docs/testing/defect-log.md`.
- **Note:** Boundary test cases typically appear in unit tests, rarely in integration tests.

## Rules


1. **Service/Domain:** Table-driven unit tests (Normal, Abnormal, Boundary).
2. **Repository:** Unit test with Mockery.
3. **Coverage:** > 80%, covering all logical paths.
4. **Parallelism:** Use `t.Parallel()` where possible.
5. **Race Detection:** Run with `-race`.
